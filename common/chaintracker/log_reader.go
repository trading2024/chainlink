package chaintracker

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/core/utils/mathutil"
)

type LogReader interface {
	Backfill(ctx context.Context, start, end int64) error
	Filter(ctx context.Context, from, to *big.Int, bh *common.Hash) ([]types.Log, error)
	PruneLogs(ctx context.Context) error
}

type LogReaderClient interface {
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	ConfiguredChainID() *big.Int
}

type logReader struct {
	ec                LogReaderClient
	lf                LogFilterer
	orm               *ORM
	lggr              logger.Logger
	backfillBatchSize int64 // batch size to use when backfilling finalized logs
	rpcBatchSize      int64 // batch size to use for fallback RPC calls made in GetBlocks
}

func NewLogReader(orm *ORM, ec LogReaderClient, lf LogFilterer, lggr logger.Logger, pollPeriod time.Duration,
	finalityDepth int64, backfillBatchSize int64, rpcBatchSize int64, keepBlocksDepth int64) *logReader {

	return &logReader{
		ec:                ec,
		lf:                lf,
		orm:               orm,
		lggr:              lggr.Named("LogReader"),
		backfillBatchSize: backfillBatchSize,
		rpcBatchSize:      rpcBatchSize,
	}
}

// convertLogs converts an array of geth logs ([]type.Log) to an array of logpoller logs ([]Log)
//
//	Block timestamps are extracted from blocks param.  If len(blocks) == 1, the same timestamp from this block
//	will be used for all logs.  If len(blocks) == len(logs) then the block number of each block is used for the
//	corresponding log.  Any other length for blocks is invalid.
func convertLogs(logs []types.Log, blocks []LogPollerBlock, lggr logger.Logger, chainID *big.Int) []Log {
	var lgs []Log
	blockTimestamp := time.Now()
	if len(logs) == 0 {
		return lgs
	}
	if len(blocks) != 1 && len(blocks) != len(logs) {
		lggr.Errorf("AssumptionViolation:  invalid params passed to convertLogs, length of blocks must either be 1 or match length of logs")
		return lgs
	}

	for i, l := range logs {
		if i == 0 || len(blocks) == len(logs) {
			blockTimestamp = blocks[i].BlockTimestamp
		}
		lgs = append(lgs, Log{
			EvmChainId: utils.NewBig(chainID),
			LogIndex:   int64(l.Index),
			BlockHash:  l.BlockHash,
			// We assume block numbers fit in int64
			// in many places.
			BlockNumber:    int64(l.BlockNumber),
			BlockTimestamp: blockTimestamp,
			EventSig:       l.Topics[0], // First topic is always event signature.
			Topics:         convertTopics(l.Topics),
			Address:        l.Address,
			TxHash:         l.TxHash,
			Data:           l.Data,
		})
	}
	return lgs
}

func convertTopics(topics []common.Hash) [][]byte {
	var topicsForDB [][]byte
	for _, t := range topics {
		topicsForDB = append(topicsForDB, t.Bytes())
	}
	return topicsForDB
}

func (lr *logReader) blocksFromLogs(ctx context.Context, logs []types.Log) (blocks []LogPollerBlock, err error) {
	var numbers []uint64
	for _, log := range logs {
		numbers = append(numbers, log.BlockNumber)
	}

	return lr.getBlocksRange(ctx, numbers)
}

func (lr *logReader) PruneLogs(ctx context.Context) error {
	return lr.orm.DeleteExpiredLogs(pg.WithParentCtx(ctx))
}

func (lr *logReader) Filter(ctx context.Context, from, to *big.Int, bh *common.Hash) ([]types.Log, error) {
	err := lr.lf.LoadFilters(ctx)
	if err != nil {
		return nil, err
	}
	return lr.lf.Filter(ctx, from, to, bh)
}

const jsonRpcLimitExceeded = -32005 // See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1474.md

// backfill will query FilterLogs in batches for logs in the
// block range [start, end] and save them to the db.
// Retries until ctx cancelled. Will return an error if cancelled
// or if there is an error backfilling.
func (lr *logReader) Backfill(ctx context.Context, start, end int64) error {
	batchSize := lr.backfillBatchSize
	for from := start; from <= end; from += batchSize {
		to := mathutil.Min(from+batchSize-1, end)
		gethLogs, err := lr.Filter(ctx, big.NewInt(from), big.NewInt(to), nil)
		if err != nil {
			var rpcErr client.JsonError
			if errors.As(err, &rpcErr) {
				if rpcErr.Code != jsonRpcLimitExceeded {
					lr.lggr.Errorw("Unable to query for logs", "err", err, "from", from, "to", to)
					return err
				}
			}
			if batchSize == 1 {
				lr.lggr.Criticalw("Too many log results in a single block, failed to retrieve logs! Node may run in a degraded state unless LogBackfillBatchSize is increased", "err", err, "from", from, "to", to, "LogBackfillBatchSize", lr.backfillBatchSize)
				return err
			}
			batchSize /= 2
			lr.lggr.Warnw("Too many log results, halving block range batch size.  Consider increasing LogBackfillBatchSize if this happens frequently", "err", err, "from", from, "to", to, "newBatchSize", batchSize, "LogBackfillBatchSize", lr.backfillBatchSize)
			from -= batchSize // counteract +=batchSize on next loop iteration, so starting block does not change
			continue
		}
		if len(gethLogs) == 0 {
			continue
		}
		blocks, err := lr.blocksFromLogs(ctx, gethLogs)
		if err != nil {
			return err
		}

		lr.lggr.Debugw("Backfill found logs", "from", from, "to", to, "logs", len(gethLogs), "blocks", blocks)
		err = lr.orm.q.WithOpts(pg.WithParentCtx(ctx)).Transaction(func(tx pg.Queryer) error {
			return lr.orm.InsertLogs(convertLogs(gethLogs, blocks, lr.lggr, lr.ec.ConfiguredChainID()), pg.WithQueryer(tx))
		})
		if err != nil {
			lr.lggr.Warnw("Unable to insert logs, retrying", "err", err, "from", from, "to", to)
			return err
		}
	}
	return nil
}

// GetBlocksRange tries to get the specified block numbers from the log pollers
// blocks table. It falls back to the RPC for any unfulfilled requested blocks.
func (lr *logReader) getBlocksRange(ctx context.Context, numbers []uint64, qopts ...pg.QOpt) ([]LogPollerBlock, error) {
	var blocks []LogPollerBlock

	// Do nothing if no blocks are requested.
	if len(numbers) == 0 {
		return blocks, nil
	}

	// Assign the requested blocks to a mapping.
	blocksRequested := make(map[uint64]struct{})
	for _, b := range numbers {
		blocksRequested[b] = struct{}{}
	}

	// Retrieve all blocks within this range from the log poller.
	blocksFound := make(map[uint64]LogPollerBlock)
	qopts = append(qopts, pg.WithParentCtx(ctx))
	minRequestedBlock := mathutil.Min(numbers[0], numbers[1:]...)
	maxRequestedBlock := mathutil.Max(numbers[0], numbers[1:]...)
	lpBlocks, err := lr.orm.GetBlocksRange(minRequestedBlock, maxRequestedBlock, qopts...)
	if err != nil {
		lr.lggr.Warnw("Error while retrieving blocks from log pollers blocks table. Falling back to RPC...", "requestedBlocks", numbers, "err", err)
	} else {
		for _, b := range lpBlocks {
			if _, ok := blocksRequested[uint64(b.BlockNumber)]; ok {
				// Only fill requested blocks.
				blocksFound[uint64(b.BlockNumber)] = b
			}
		}
		lr.lggr.Debugw("Got blocks from log poller", "blockNumbers", maps.Keys(blocksFound))
	}

	// Fill any remaining blocks from the client.
	blocksFoundFromRPC, err := lr.fillRemainingBlocksFromRPC(ctx, numbers, blocksFound)
	if err != nil {
		return nil, err
	}
	for num, b := range blocksFoundFromRPC {
		blocksFound[num] = b
	}

	var blocksNotFound []uint64
	for _, num := range numbers {
		b, ok := blocksFound[num]
		if !ok {
			blocksNotFound = append(blocksNotFound, num)
		}
		blocks = append(blocks, b)
	}

	if len(blocksNotFound) > 0 {
		return nil, errors.Errorf("blocks were not found in db or RPC call: %v", blocksNotFound)
	}

	return blocks, nil
}

func (lr *logReader) fillRemainingBlocksFromRPC(
	ctx context.Context,
	blocksRequested []uint64,
	blocksFound map[uint64]LogPollerBlock,
) (map[uint64]LogPollerBlock, error) {
	var reqs []rpc.BatchElem
	var remainingBlocks []uint64
	for _, num := range blocksRequested {
		if _, ok := blocksFound[num]; !ok {
			req := rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{hexutil.EncodeBig(big.NewInt(0).SetUint64(num)), false},
				Result: &evmtypes.Head{},
			}
			reqs = append(reqs, req)
			remainingBlocks = append(remainingBlocks, num)
		}
	}

	if len(remainingBlocks) > 0 {
		lr.lggr.Debugw("Falling back to RPC for blocks not found in log poller blocks table",
			"remainingBlocks", remainingBlocks)
	}

	for i := 0; i < len(reqs); i += int(lr.rpcBatchSize) {
		j := i + int(lr.rpcBatchSize)
		if j > len(reqs) {
			j = len(reqs)
		}

		err := lr.ec.BatchCallContext(ctx, reqs[i:j])
		if err != nil {
			return nil, err
		}
	}

	var blocksFoundFromRPC = make(map[uint64]LogPollerBlock)
	for _, r := range reqs {
		if r.Error != nil {
			return nil, r.Error
		}
		block, is := r.Result.(*evmtypes.Head)

		if !is {
			return nil, errors.Errorf("expected result to be a %T, got %T", &evmtypes.Head{}, r.Result)
		}
		if block == nil {
			return nil, errors.New("invariant violation: got nil block")
		}
		if block.Hash == (common.Hash{}) {
			return nil, errors.Errorf("missing block hash for block number: %d", block.Number)
		}
		if block.Number < 0 {
			return nil, errors.Errorf("expected block number to be >= to 0, got %d", block.Number)
		}
		blocksFoundFromRPC[uint64(block.Number)] = LogPollerBlock{
			EvmChainId:     block.EVMChainID,
			BlockHash:      block.Hash,
			BlockNumber:    block.Number,
			BlockTimestamp: block.Timestamp,
			CreatedAt:      block.Timestamp,
		}
	}

	return blocksFoundFromRPC, nil
}
