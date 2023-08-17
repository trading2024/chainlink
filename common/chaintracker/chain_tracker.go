package chaintracker

import (
	"context"
	"database/sql"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/core/utils/mathutil"
)

//go:generate mockery --quiet --name LogPoller --output ./mocks/ --case=underscore --structname LogPoller --filename log_poller.go
type ChainTracker interface {
	services.ServiceCtx
	Replay(ctx context.Context, fromBlock int64) error
	ReplayAsync(fromBlock int64)
	RegisterFilter(filter Filter, qopts ...pg.QOpt) error
	UnregisterFilter(name string, qopts ...pg.QOpt) error
	LatestBlock(qopts ...pg.QOpt) (int64, error)
	GetBlocksRange(ctx context.Context, numbers []uint64, qopts ...pg.QOpt) ([]LogPollerBlock, error)
}

type ChainTrackerClient interface {
	HeadByNumber(ctx context.Context, n *big.Int) (*evmtypes.Head, error)
	HeadByHash(ctx context.Context, n common.Hash) (*evmtypes.Head, error)
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	ConfiguredChainID() *big.Int
}

var (
	ErrReplayRequestAborted = errors.New("aborted, replay request cancelled")
	ErrReplayInProgress     = errors.New("replay request cancelled, but replay is already in progress")
	ErrLogPollerShutdown    = errors.New("replay aborted due to log poller shutdown")
)

type chainTracker struct {
	utils.StartStopOnce
	ec                    ChainTrackerClient
	lr                    LogReader
	orm                   *ORM
	lggr                  logger.Logger
	pollPeriod            time.Duration // poll period set by block production rate
	finalityDepth         int64         // finality depth is taken to mean that block (head - finality) is finalized
	keepBlocksDepth       int64         // the number of blocks behind the head for which we keep the blocks. Must be greater than finality depth + 1.
	backfillBatchSize     int64         // batch size to use when backfilling finalized logs
	rpcBatchSize          int64         // batch size to use for fallback RPC calls made in GetBlocks
	backupPollerNextBlock int64

	filterMu        sync.RWMutex
	filters         map[string]Filter
	filterDirty     bool

	replayStart    chan int64
	replayComplete chan error
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
}

// NewLogPoller creates a log poller. Note there is an assumption
// that blocks can be processed faster than they are produced for the given chain, or the poller will fall behind.
// Block processing involves the following calls in steady state (without reorgs):
//   - eth_getBlockByNumber - headers only (transaction hashes, not full transaction objects),
//   - eth_getLogs - get the logs for the block
//   - 1 db read latest block - for checking reorgs
//   - 1 db tx including block write and logs write to logs.
//
// How fast that can be done depends largely on network speed and DB, but even for the fastest
// support chain, polygon, which has 2s block times, we need RPCs roughly with <= 500ms latency
func NewChainTracker(orm *ORM, ec ChainTrackerClient, lr LogReader, lggr logger.Logger, pollPeriod time.Duration,
	finalityDepth int64, backfillBatchSize int64, rpcBatchSize int64, keepBlocksDepth int64) *chainTracker {

	return &chainTracker{
		ec:                ec,
		lr:                lr,
		orm:               orm,
		lggr:              lggr.Named("LogPoller"),
		replayStart:       make(chan int64),
		replayComplete:    make(chan error),
		pollPeriod:        pollPeriod,
		finalityDepth:     finalityDepth,
		backfillBatchSize: backfillBatchSize,
		rpcBatchSize:      rpcBatchSize,
		keepBlocksDepth:   keepBlocksDepth,
		filters:           make(map[string]Filter),
		filterDirty:       true, // Always build Filter on first call to cache an empty filter if nothing registered yet.
	}
}

// Replay signals that the poller should resume from a new block.
// Blocks until the replay is complete.
// Replay can be used to ensure that filter modification has been applied for all blocks from "fromBlock" up to latest.
// If ctx is cancelled before the replay request has been initiated, ErrReplayRequestAborted is returned.  If the replay
// is already in progress, the replay will continue and ErrReplayInProgress will be returned.  If the client needs a
// guarantee that the replay is complete before proceeding, it should either avoid cancelling or retry until nil is returned
func (ct *chainTracker) Replay(ctx context.Context, fromBlock int64) error {
	ct.lggr.Debugf("Replaying from block %d", fromBlock)
	latest, err := ct.ec.HeadByNumber(ctx, nil)
	if err != nil {
		return err
	}
	if fromBlock < 1 || fromBlock > latest.Number {
		return errors.Errorf("Invalid replay block number %v, acceptable range [1, %v]", fromBlock, latest.Number)
	}
	// Block until replay notification accepted or cancelled.
	select {
	case ct.replayStart <- fromBlock:
	case <-ctx.Done():
		return errors.Wrap(ErrReplayRequestAborted, ctx.Err().Error())
	}
	// Block until replay complete or cancelled.
	select {
	case err = <-ct.replayComplete:
		return err
	case <-ctx.Done():
		// Note: this will not abort the actual replay, it just means the client gave up on waiting for it to complete
		ct.wg.Add(1)
		go ct.recvReplayComplete()
		return ErrReplayInProgress
	}
}

func (ct *chainTracker) recvReplayComplete() {
	err := <-ct.replayComplete
	if err != nil {
		ct.lggr.Error(err)
	}
	ct.wg.Done()
}

// Asynchronous wrapper for Replay()
func (ct *chainTracker) ReplayAsync(fromBlock int64) {
	ct.wg.Add(1)
	go func() {
		if err := ct.Replay(context.Background(), fromBlock); err != nil {
			ct.lggr.Error(err)
		}
		ct.wg.Done()
	}()
}

func (ct *chainTracker) Start(parentCtx context.Context) error {
	if ct.keepBlocksDepth < (ct.finalityDepth + 1) {
		// We add 1 since for reorg detection on the first unfinalized block
		// we need to keep 1 finalized block.
		return errors.Errorf("keepBlocksDepth %d must be greater than finality %d + 1", ct.keepBlocksDepth, ct.finalityDepth)
	}
	return ct.StartOnce("ChainTracker", func() error {
		ctx, cancel := context.WithCancel(parentCtx)
		ct.ctx = ctx
		ct.cancel = cancel
		ct.wg.Add(1)
		go ct.run()
		return nil
	})
}

func (ct *chainTracker) Close() error {
	return ct.StopOnce("ChainTracker", func() error {
		select {
		case ct.replayComplete <- ErrLogPollerShutdown:
		default:
		}
		ct.cancel()
		ct.wg.Wait()
		return nil
	})
}

func (ct *chainTracker) Name() string {
	return ct.lggr.Name()
}

func (ct *chainTracker) HealthReport() map[string]error {
	return map[string]error{ct.Name(): ct.StartStopOnce.Healthy()}
}

func (ct *chainTracker) GetReplayFromBlock(ctx context.Context, requested int64) (int64, error) {
	lastProcessed, err := ct.orm.SelectLatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			// Real DB error
			return 0, err
		}
		// Nothing in db, use requested
		return requested, nil
	}
	// We have lastProcessed, take min(requested, lastProcessed).
	// This is to avoid replaying from a block later than what we have in the DB
	// and skipping blocks.
	return mathutil.Min(requested, lastProcessed.BlockNumber), nil
}

func (ct *chainTracker) run() {
	defer ct.wg.Done()
	logPollTick := time.After(0)
	// stagger these somewhat, so they don't all run back-to-back
	backupLogPollTick := time.After(100 * time.Millisecond)
	blockPruneTick := time.After(3 * time.Second)
	logPruneTick := time.After(5 * time.Second)
	filtersLoaded := false

	loadFilters := func() error {
		ct.filterMu.Lock()
		defer ct.filterMu.Unlock()
		filters, err := ct.orm.LoadFilters(pg.WithParentCtx(ct.ctx))

		if err != nil {
			return errors.Wrapf(err, "Failed to load initial filters from db, retrying")
		}

		ct.filters = filters
		ct.filterDirty = true
		filtersLoaded = true
		return nil
	}

	for {
		select {
		case <-ct.ctx.Done():
			return
		case fromBlockReq := <-ct.replayStart:
			fromBlock, err := ct.GetReplayFromBlock(ct.ctx, fromBlockReq)
			if err == nil {
				if !filtersLoaded {
					ct.lggr.Warnw("Received replayReq before filters loaded", "fromBlock", fromBlock, "requested", fromBlockReq)
					if err = loadFilters(); err != nil {
						ct.lggr.Errorw("Failed loading filters during Replay", "err", err, "fromBlock", fromBlock)
					}
				}
				if err == nil {
					// Serially process replay requests.
					ct.lggr.Infow("Executing replay", "fromBlock", fromBlock, "requested", fromBlockReq)
					ct.PollAndSaveLogs(ct.ctx, fromBlock)
				}
			} else {
				ct.lggr.Errorw("Error executing replay, could not get fromBlock", "err", err)
			}
			select {
			case <-ct.ctx.Done():
				// We're shutting down, notify client and exit
				select {
				case ct.replayComplete <- ErrReplayRequestAborted:
				default:
				}
				return
			case ct.replayComplete <- err:
			}
		case <-logPollTick:
			logPollTick = time.After(utils.WithJitter(ct.pollPeriod))
			if !filtersLoaded {
				if err := loadFilters(); err != nil {
					ct.lggr.Errorw("Failed loading filters in main logpoller loop, retrying later", "err", err)
					continue
				}
			}

			// Always start from the latest block in the db.
			var start int64
			lastProcessed, err := ct.orm.SelectLatestBlock(pg.WithParentCtx(ct.ctx))
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					// Assume transient db reading issue, retry forever.
					ct.lggr.Errorw("unable to get starting block", "err", err)
					continue
				}
				// Otherwise this is the first poll _ever_ on a new chain.
				// Only safe thing to do is to start at the first finalized block.
				latest, err := ct.ec.HeadByNumber(ct.ctx, nil)
				if err != nil {
					ct.lggr.Warnw("Unable to get latest for first poll", "err", err)
					continue
				}
				latestNum := latest.Number
				// Do not support polling chains which don't even have finality depth worth of blocks.
				// Could conceivably support this but not worth the effort.
				// Need finality depth + 1, no block 0.
				if latestNum <= ct.finalityDepth {
					ct.lggr.Warnw("Insufficient number of blocks on chain, waiting for finality depth", "err", err, "latest", latestNum, "finality", ct.finalityDepth)
					continue
				}
				// Starting at the first finalized block. We do not backfill the first finalized block.
				start = latestNum - ct.finalityDepth
			} else {
				start = lastProcessed.BlockNumber + 1
			}
			ct.PollAndSaveLogs(ct.ctx, start)
		case <-backupLogPollTick:
			// Backup log poller:  this serves as an emergency backup to protect against eventual-consistency behavior
			// of an rpc node (seen occasionally on optimism, but possibly could happen on other chains?).  If the first
			// time we request a block, no logs or incomplete logs come back, this ensures that every log is eventually
			// re-requested after it is finalized.  This doesn't add much overhead, because we can request all of them
			// in one shot, since we don't need to worry about re-orgs after finality depth, and it runs 100x less
			// frequently than the primary log poller.

			// If pollPeriod is set to 1 block time, backup log poller will run once every 100 blocks
			const backupPollerBlockDelay = 100

			backupLogPollTick = time.After(utils.WithJitter(backupPollerBlockDelay * ct.pollPeriod))
			if !filtersLoaded {
				ct.lggr.Warnw("Backup log poller ran before filters loaded, skipping")
				continue
			}
			ct.BackupPollAndSaveLogs(ct.ctx, backupPollerBlockDelay)
		case <-blockPruneTick:
			blockPruneTick = time.After(utils.WithJitter(ct.pollPeriod * 1000))
			if err := ct.pruneOldBlocks(ct.ctx); err != nil {
				ct.lggr.Errorw("Unable to prune old blocks", "err", err)
			}
		case <-logPruneTick:
			logPruneTick = time.After(utils.WithJitter(ct.pollPeriod * 2401)) // = 7^5 avoids common factors with 1000
			if err := ct.lr.PruneLogs(ct.ctx); err != nil {
				ct.lggr.Error(err)
			}
		}
	}
}

func (ct *chainTracker) BackupPollAndSaveLogs(ctx context.Context, backupPollerBlockDelay int64) {
	if ct.backupPollerNextBlock == 0 {
		lastProcessed, err := ct.orm.SelectLatestBlock(pg.WithParentCtx(ctx))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ct.lggr.Warnw("Backup log poller ran before first successful log poller run, skipping")
			} else {
				ct.lggr.Errorw("Backup log poller unable to get starting block", "err", err)
			}
			return
		}

		// If this is our first run, start max(finalityDepth+1, backupPollerBlockDelay) blocks behind the last processed
		// (or at block 0 if whole blockchain is too short)
		ct.backupPollerNextBlock = lastProcessed.BlockNumber - mathutil.Max(ct.finalityDepth+1, backupPollerBlockDelay)
		if ct.backupPollerNextBlock < 0 {
			ct.backupPollerNextBlock = 0
		}
	}

	latestBlock, err := ct.ec.HeadByNumber(ctx, nil)
	if err != nil {
		ct.lggr.Warnw("Backup logpoller failed to get latest block", "err", err)
		return
	}

	lastSafeBackfillBlock := latestBlock.Number - ct.finalityDepth - 1
	if lastSafeBackfillBlock >= ct.backupPollerNextBlock {
		ct.lggr.Infow("Backup poller backfilling logs", "start", ct.backupPollerNextBlock, "end", lastSafeBackfillBlock)
		if err = ct.lr.Backfill(ctx, ct.backupPollerNextBlock, lastSafeBackfillBlock); err != nil {
			// If there's an error backfilling, we can just return and retry from the last block saved
			// since we don't save any blocks on backfilling. We may re-insert the same logs but thats ok.
			ct.lggr.Warnw("Backup poller failed", "err", err)
			return
		}
		ct.backupPollerNextBlock = lastSafeBackfillBlock + 1
	}
}

// getCurrentBlockMaybeHandleReorg accepts a block number
// and will return that block if its parent points to our last saved block.
// One can optionally pass the block header if it has already been queried to avoid an extra RPC call.
// If its parent does not point to our last saved block we know a reorg has occurred,
// so we:
// 1. Find the LCA by following parent hashes.
// 2. Delete all logs and blocks after the LCA
// 3. Return the LCA+1, i.e. our new current (unprocessed) block.
func (ct *chainTracker) getCurrentBlockMaybeHandleReorg(ctx context.Context, currentBlockNumber int64, currentBlock *evmtypes.Head) (*evmtypes.Head, error) {
	var err1 error
	if currentBlock == nil {
		// If we don't have the current block already, lets get it.
		currentBlock, err1 = ct.ec.HeadByNumber(ctx, big.NewInt(currentBlockNumber))
		if err1 != nil {
			ct.lggr.Warnw("Unable to get currentBlock", "err", err1, "currentBlockNumber", currentBlockNumber)
			return nil, err1
		}
		// Additional sanity checks, don't necessarily trust the RPC.
		if currentBlock == nil {
			ct.lggr.Errorf("Unexpected nil block from RPC", "currentBlockNumber", currentBlockNumber)
			return nil, errors.Errorf("Got nil block for %d", currentBlockNumber)
		}
		if currentBlock.Number != currentBlockNumber {
			ct.lggr.Warnw("Unable to get currentBlock, rpc returned incorrect block", "currentBlockNumber", currentBlockNumber, "got", currentBlock.Number)
			return nil, errors.Errorf("Block mismatch have %d want %d", currentBlock.Number, currentBlockNumber)
		}
	}
	// Does this currentBlock point to the same parent that we have saved?
	// If not, there was a reorg, so we need to rewind.
	expectedParent, err1 := ct.orm.SelectBlockByNumber(currentBlockNumber-1, pg.WithParentCtx(ctx))
	if err1 != nil && !errors.Is(err1, sql.ErrNoRows) {
		// If err is not a 'no rows' error, assume transient db issue and retry
		ct.lggr.Warnw("Unable to read latestBlockNumber currentBlock saved", "err", err1, "currentBlockNumber", currentBlockNumber)
		return nil, errors.New("Unable to read latestBlockNumber currentBlock saved")
	}
	// We will not have the previous currentBlock on initial poll.
	havePreviousBlock := err1 == nil
	if !havePreviousBlock {
		ct.lggr.Infow("Do not have previous block, first poll ever on new chain or after backfill", "currentBlockNumber", currentBlockNumber)
		return currentBlock, nil
	}
	// Check for reorg.
	if currentBlock.ParentHash != expectedParent.BlockHash {
		// There can be another reorg while we're finding the LCA.
		// That is ok, since we'll detect it on the next iteration.
		// Since we go currentBlock by currentBlock for unfinalized logs, the mismatch starts at currentBlockNumber - 1.
		blockAfterLCA, err2 := ct.findBlockAfterLCA(ctx, currentBlock)
		if err2 != nil {
			ct.lggr.Warnw("Unable to find LCA after reorg, retrying", "err", err2)
			return nil, errors.New("Unable to find LCA after reorg, retrying")
		}

		ct.lggr.Infow("Reorg detected", "blockAfterLCA", blockAfterLCA.Number, "currentBlockNumber", currentBlockNumber)
		// We truncate all the blocks and logs after the LCA.
		// We could preserve the logs for forensics, since its possible
		// that applications see them and take action upon it, however that
		// results in significantly slower reads since we must then compute
		// the canonical set per read. Typically, if an application took action on a log
		// it would be saved elsewhere e.g. eth_txes, so it seems better to just support the fast reads.
		// Its also nicely analogous to reading from the chain itself.
		err2 = ct.orm.q.WithOpts(pg.WithParentCtx(ctx)).Transaction(func(tx pg.Queryer) error {
			// These deletes are bounded by reorg depth, so they are
			// fast and should not slow down the log readers.
			err3 := ct.orm.DeleteBlocksAfter(blockAfterLCA.Number, pg.WithQueryer(tx))
			if err3 != nil {
				ct.lggr.Warnw("Unable to clear reorged blocks, retrying", "err", err3)
				return err3
			}
			err3 = ct.orm.DeleteLogsAfter(blockAfterLCA.Number, pg.WithQueryer(tx))
			if err3 != nil {
				ct.lggr.Warnw("Unable to clear reorged logs, retrying", "err", err3)
				return err3
			}
			return nil
		})
		if err2 != nil {
			// If we error on db commit, we can't know if the tx went through or not.
			// We return an error here which will cause us to restart polling from lastBlockSaved + 1
			return nil, err2
		}
		return blockAfterLCA, nil
	}
	// No reorg, return current block.
	return currentBlock, nil
}

// PollAndSaveLogs On startup/crash current is the first block after the last processed block.
// currentBlockNumber is the block from where new logs are to be polled & saved. Under normal
// conditions this would be equal to lastProcessed.BlockNumber + 1.
func (ct *chainTracker) PollAndSaveLogs(ctx context.Context, currentBlockNumber int64) {
	ct.lggr.Debugw("Polling for logs", "currentBlockNumber", currentBlockNumber)
	latestBlock, err := ct.ec.HeadByNumber(ctx, nil)
	if err != nil {
		ct.lggr.Warnw("Unable to get latestBlockNumber block", "err", err, "currentBlockNumber", currentBlockNumber)
		return
	}
	latestBlockNumber := latestBlock.Number
	if currentBlockNumber > latestBlockNumber {
		// Note there can also be a reorg "shortening" i.e. chain height decreases but TDD increases. In that case
		// we also just wait until the new tip is longer and then detect the reorg.
		ct.lggr.Debugw("No new blocks since last poll", "currentBlockNumber", currentBlockNumber, "latestBlockNumber", latestBlockNumber)
		return
	}
	var currentBlock *evmtypes.Head
	if currentBlockNumber == latestBlockNumber {
		// Can re-use our currentBlock and avoid an extra RPC call.
		currentBlock = latestBlock
	}
	// Possibly handle a reorg. For example if we crash, we'll be in the middle of processing unfinalized blocks.
	// Returns (currentBlock || LCA+1 if reorg detected, error)
	currentBlock, err = ct.getCurrentBlockMaybeHandleReorg(ctx, currentBlockNumber, currentBlock)
	if err != nil {
		// If there's an error handling the reorg, we can't be sure what state the db was left in.
		// Resume from the latest block saved and retry.
		ct.lggr.Errorw("Unable to get current block, retrying", "err", err)
		return
	}
	currentBlockNumber = currentBlock.Number

	// backfill finalized blocks if we can for performance. If we crash during backfill, we
	// may reprocess logs.  Log insertion is idempotent so this is ok.
	// E.g. 1<-2<-3(currentBlockNumber)<-4<-5<-6<-7(latestBlockNumber), finality is 2. So 3,4 can be batched.
	// Although 5 is finalized, we still need to save it to the db for reorg detection if 6 is a reorg.
	// start = currentBlockNumber = 3, end = latestBlockNumber - finality - 1 = 7-2-1 = 4 (inclusive range).
	lastSafeBackfillBlock := latestBlockNumber - ct.finalityDepth - 1
	if lastSafeBackfillBlock >= currentBlockNumber {
		ct.lggr.Infow("Backfilling logs", "start", currentBlockNumber, "end", lastSafeBackfillBlock)
		if err = ct.lr.Backfill(ctx, currentBlockNumber, lastSafeBackfillBlock); err != nil {
			// If there's an error backfilling, we can just return and retry from the last block saved
			// since we don't save any blocks on backfilling. We may re-insert the same logs but thats ok.
			ct.lggr.Warnw("Unable to backfill finalized logs, retrying later", "err", err)
			return
		}
		currentBlockNumber = lastSafeBackfillBlock + 1
	}

	if currentBlockNumber > currentBlock.Number {
		// If we successfully backfilled we have logs up to and including lastSafeBackfillBlock,
		// now load the first unfinalized block.
		currentBlock, err = ct.getCurrentBlockMaybeHandleReorg(ctx, currentBlockNumber, nil)
		if err != nil {
			// If there's an error handling the reorg, we can't be sure what state the db was left in.
			// Resume from the latest block saved.
			ct.lggr.Errorw("Unable to get current block", "err", err)
			return
		}
	}

	for {
		h := currentBlock.Hash
		var logs []types.Log
		logs, err = ct.lr.Filter(ctx, nil, nil, &h)
		if err != nil {
			ct.lggr.Warnw("Unable to query for logs, retrying", "err", err, "block", currentBlockNumber)
			return
		}
		ct.lggr.Debugw("Unfinalized log query", "logs", len(logs), "currentBlockNumber", currentBlockNumber, "blockHash", currentBlock.Hash, "timestamp", currentBlock.Timestamp.Unix())
		err = ct.orm.q.WithOpts(pg.WithParentCtx(ctx)).Transaction(func(tx pg.Queryer) error {
			if err2 := ct.orm.InsertBlock(h, currentBlockNumber, currentBlock.Timestamp, pg.WithQueryer(tx)); err2 != nil {
				return err2
			}
			if len(logs) == 0 {
				return nil
			}
			return ct.orm.InsertLogs(convertLogs(logs,
				[]LogPollerBlock{{BlockNumber: currentBlockNumber,
					BlockTimestamp: currentBlock.Timestamp}},
				ct.lggr,
				ct.ec.ConfiguredChainID(),
			), pg.WithQueryer(tx))
		})
		if err != nil {
			ct.lggr.Warnw("Unable to save logs resuming from last saved block + 1", "err", err, "block", currentBlockNumber)
			return
		}
		// Update current block.
		// Same reorg detection on unfinalized blocks.
		currentBlockNumber++
		if currentBlockNumber > latestBlockNumber {
			break
		}
		currentBlock, err = ct.getCurrentBlockMaybeHandleReorg(ctx, currentBlockNumber, nil)
		if err != nil {
			// If there's an error handling the reorg, we can't be sure what state the db was left in.
			// Resume from the latest block saved.
			ct.lggr.Errorw("Unable to get current block", "err", err)
			return
		}
		currentBlockNumber = currentBlock.Number
	}
}

// Find the first place where our chain and their chain have the same block,
// that block number is the LCA. Return the block after that, where we want to resume polling.
func (ct *chainTracker) findBlockAfterLCA(ctx context.Context, current *evmtypes.Head) (*evmtypes.Head, error) {
	// Current is where the mismatch starts.
	// Check its parent to see if its the same as ours saved.
	parent, err := ct.ec.HeadByHash(ctx, current.ParentHash)
	if err != nil {
		return nil, err
	}
	blockAfterLCA := *current
	reorgStart := parent.Number
	// We expect reorgs up to the block after (current - finalityDepth),
	// since the block at (current - finalityDepth) is finalized.
	// We loop via parent instead of current so current always holds the LCA+1.
	// If the parent block number becomes < the first finalized block our reorg is too deep.
	for parent.Number >= (reorgStart - ct.finalityDepth) {
		ourParentBlockHash, err := ct.orm.SelectBlockByNumber(parent.Number, pg.WithParentCtx(ctx))
		if err != nil {
			return nil, err
		}
		if parent.Hash == ourParentBlockHash.BlockHash {
			// If we do have the blockhash, return blockAfterLCA
			return &blockAfterLCA, nil
		}
		// Otherwise get a new parent and update blockAfterLCA.
		blockAfterLCA = *parent
		parent, err = ct.ec.HeadByHash(ctx, parent.ParentHash)
		if err != nil {
			return nil, err
		}
	}
	ct.lggr.Criticalw("Reorg greater than finality depth detected", "max reorg depth", ct.finalityDepth-1)
	rerr := errors.New("Reorg greater than finality depth")
	ct.SvcErrBuffer.Append(rerr)
	return nil, rerr
}

// pruneOldBlocks removes blocks that are > ct.ancientBlockDepth behind the head.
func (ct *chainTracker) pruneOldBlocks(ctx context.Context) error {
	latest, err := ct.ec.HeadByNumber(ctx, nil)
	if err != nil {
		return err
	}
	if latest == nil {
		return errors.Errorf("received nil block from RPC")
	}
	if latest.Number <= ct.keepBlocksDepth {
		// No-op, keep all blocks
		return nil
	}
	// 1-2-3-4-5(latest), keepBlocksDepth=3
	// Remove <= 2
	return ct.orm.DeleteBlocksBefore(latest.Number-ct.keepBlocksDepth, pg.WithParentCtx(ctx))
}
