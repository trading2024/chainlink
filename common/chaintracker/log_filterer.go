package chaintracker

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type LogFilterer interface {
	RegisterFilter(filter Filter, qopts ...pg.QOpt) error
	UnregisterFilter(name string, qopts ...pg.QOpt) error
	Filter(ctx context.Context, from, to *big.Int, bh *common.Hash) ([]types.Log, error)
	LoadFilters(ctx context.Context) error

	// General querying
	Logs(start, end int64, eventSig common.Hash, address common.Address, qopts ...pg.QOpt) ([]Log, error)
	LogsWithSigs(start, end int64, eventSigs []common.Hash, address common.Address, qopts ...pg.QOpt) ([]Log, error)
	LogsCreatedAfter(eventSig common.Hash, address common.Address, time time.Time, confs int, qopts ...pg.QOpt) ([]Log, error)
	LatestLogByEventSigWithConfs(eventSig common.Hash, address common.Address, confs int, qopts ...pg.QOpt) (*Log, error)
	LatestLogEventSigsAddrsWithConfs(fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs int, qopts ...pg.QOpt) ([]Log, error)
	LatestBlockByEventSigsAddrsWithConfs(fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs int, qopts ...pg.QOpt) (int64, error)

	// Content based querying
	IndexedLogs(eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error)
	IndexedLogsByBlockRange(start, end int64, eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, qopts ...pg.QOpt) ([]Log, error)
	IndexedLogsCreatedAfter(eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, after time.Time, confs int, qopts ...pg.QOpt) ([]Log, error)
	IndexedLogsTopicGreaterThan(eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error)
	IndexedLogsTopicRange(eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, topicValueMax common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error)
	IndexedLogsWithSigsExcluding(address common.Address, eventSigA, eventSigB common.Hash, topicIndex int, fromBlock, toBlock int64, confs int, qopts ...pg.QOpt) ([]Log, error)
	LogsDataWordRange(eventSig common.Hash, address common.Address, wordIndex int, wordValueMin, wordValueMax common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error)
	LogsDataWordGreaterThan(eventSig common.Hash, address common.Address, wordIndex int, wordValueMin common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error)
}

type LogFiltererClient interface {
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
}

type logFilterer struct {
	utils.StartStopOnce
	ec                LogFiltererClient
	orm               *ORM
	lggr              logger.Logger
	backfillBatchSize int64 // batch size to use when backfilling finalized logs
	rpcBatchSize      int64 // batch size to use for fallback RPC calls made in GetBlocks

	filterMu        sync.RWMutex
	filters         map[string]Filter
	filterDirty     bool
	filtersLoaded   bool
	cachedAddresses []common.Address
	cachedEventSigs []common.Hash
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
func NewLogPoller(orm *ORM, ec LogFiltererClient, lggr logger.Logger,
	finalityDepth int64, backfillBatchSize int64, rpcBatchSize int64) *logFilterer {

	return &logFilterer{
		ec:                ec,
		orm:               orm,
		lggr:              lggr.Named("LogPoller"),
		backfillBatchSize: backfillBatchSize,
		rpcBatchSize:      rpcBatchSize,
		filters:           make(map[string]Filter),
		filterDirty:       true, // Always build Filter on first call to cache an empty filter if nothing registered yet.
		filtersLoaded:     false,
	}
}

type Filter struct {
	Name      string // see FilterName(id, args) below
	EventSigs evmtypes.HashArray
	Addresses evmtypes.AddressArray
	Retention time.Duration
}

// FilterName is a suggested convenience function for clients to construct unique filter names
// to populate Name field of struct Filter
func FilterName(id string, args ...any) string {
	if len(args) == 0 {
		return id
	}
	s := &strings.Builder{}
	s.WriteString(id)
	s.WriteString(" - ")
	fmt.Fprintf(s, "%s", args[0])
	for _, a := range args[1:] {
		fmt.Fprintf(s, ":%s", a)
	}
	return s.String()
}

// Contains returns true if this filter already fully Contains a
// filter passed to it.
func (filter *Filter) Contains(other *Filter) bool {
	if other == nil {
		return true
	}
	addresses := make(map[common.Address]interface{})
	for _, addr := range filter.Addresses {
		addresses[addr] = struct{}{}
	}
	events := make(map[common.Hash]interface{})
	for _, ev := range filter.EventSigs {
		events[ev] = struct{}{}
	}

	for _, addr := range other.Addresses {
		if _, ok := addresses[addr]; !ok {
			return false
		}
	}
	for _, ev := range other.EventSigs {
		if _, ok := events[ev]; !ok {
			return false
		}
	}
	return true
}

// RegisterFilter adds the provided EventSigs and Addresses to the log poller's log filter query.
// If any eventSig is emitted from any address, it will be captured by the log poller.
// If an event matching any of the given event signatures is emitted from any of the provided Addresses,
// the log poller will pick those up and save them. For topic specific queries see content based querying.
// Clients may choose to MergeFilter and then Replay in order to ensure desired logs are present.
// NOTE: due to constraints of the eth filter, there is "leakage" between successive MergeFilter calls, for example
//
//	RegisterFilter(event1, addr1)
//	RegisterFilter(event2, addr2)
//
// will result in the poller saving (event1, addr2) or (event2, addr1) as well, should it exist.
// Generally speaking this is harmless. We enforce that EventSigs and Addresses are non-empty,
// which means that anonymous events are not supported and log.Topics >= 1 always (log.Topics[0] is the event signature).
// The filter may be unregistered later by Filter.Name
func (lf *logFilterer) RegisterFilter(filter Filter, qopts ...pg.QOpt) error {
	if len(filter.Addresses) == 0 {
		return errors.Errorf("at least one address must be specified")
	}
	if len(filter.EventSigs) == 0 {
		return errors.Errorf("at least one event must be specified")
	}

	for _, eventSig := range filter.EventSigs {
		if eventSig == [common.HashLength]byte{} {
			return errors.Errorf("empty event sig")
		}
	}
	for _, addr := range filter.Addresses {
		if addr == [common.AddressLength]byte{} {
			return errors.Errorf("empty address")
		}
	}

	lf.filterMu.Lock()
	defer lf.filterMu.Unlock()

	if existingFilter, ok := lf.filters[filter.Name]; ok {
		if existingFilter.Contains(&filter) {
			// Nothing new in this Filter
			return nil
		}
		lf.lggr.Warnw("Updating existing filter with more events or addresses", "filter", filter)
	} else {
		lf.lggr.Debugw("Creating new filter", "filter", filter)
	}

	if err := lf.orm.InsertFilter(filter, qopts...); err != nil {
		return errors.Wrap(err, "RegisterFilter failed to save filter to db")
	}
	lf.filters[filter.Name] = filter
	lf.filterDirty = true
	return nil
}

func (lf *logFilterer) UnregisterFilter(name string, qopts ...pg.QOpt) error {
	lf.filterMu.Lock()
	defer lf.filterMu.Unlock()

	_, ok := lf.filters[name]
	if !ok {
		lf.lggr.Errorf("Filter %s not found", name)
		return nil
	}

	if err := lf.orm.DeleteFilter(name, qopts...); err != nil {
		return errors.Wrapf(err, "Failed to delete filter %s", name)
	}
	delete(lf.filters, name)
	lf.filterDirty = true
	return nil
}

func (lf *logFilterer) Filter(ctx context.Context, from, to *big.Int, bh *common.Hash) ([]types.Log, error) {
	return lf.ec.FilterLogs(ctx, lf.filter(from, to, bh))
}

func (lf *logFilterer) filter(from, to *big.Int, bh *common.Hash) ethereum.FilterQuery {
	lf.filterMu.Lock()
	defer lf.filterMu.Unlock()
	if !lf.filterDirty {
		return ethereum.FilterQuery{FromBlock: from, ToBlock: to, BlockHash: bh, Topics: [][]common.Hash{lf.cachedEventSigs}, Addresses: lf.cachedAddresses}
	}
	var (
		addresses  []common.Address
		eventSigs  []common.Hash
		addressMp  = make(map[common.Address]struct{})
		eventSigMp = make(map[common.Hash]struct{})
	)
	// Merge filters.
	for _, filter := range lf.filters {
		for _, addr := range filter.Addresses {
			addressMp[addr] = struct{}{}
		}
		for _, eventSig := range filter.EventSigs {
			eventSigMp[eventSig] = struct{}{}
		}
	}
	for addr := range addressMp {
		addresses = append(addresses, addr)
	}
	sort.Slice(addresses, func(i, j int) bool {
		return bytes.Compare(addresses[i][:], addresses[j][:]) < 0
	})
	for eventSig := range eventSigMp {
		eventSigs = append(eventSigs, eventSig)
	}
	sort.Slice(eventSigs, func(i, j int) bool {
		return bytes.Compare(eventSigs[i][:], eventSigs[j][:]) < 0
	})
	if len(eventSigs) == 0 && len(addresses) == 0 {
		// If no filter specified, ignore everything.
		// This allows us to keep the log poller up and running with no filters present (e.g. no jobs on the node),
		// then as jobs are added dynamically start using their filters.
		addresses = []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000000")}
		eventSigs = []common.Hash{}
	}
	lf.cachedAddresses = addresses
	lf.cachedEventSigs = eventSigs
	lf.filterDirty = false
	return ethereum.FilterQuery{FromBlock: from, ToBlock: to, BlockHash: bh, Topics: [][]common.Hash{eventSigs}, Addresses: addresses}
}
func (lf *logFilterer) LoadFilters(ctx context.Context) error {
	if lf.filtersLoaded {
		return nil
	}
	lf.filterMu.Lock()
	defer lf.filterMu.Unlock()
	filters, err := lf.orm.LoadFilters(pg.WithParentCtx(ctx))

	if err != nil {
		return errors.Wrapf(err, "Failed to load initial filters from db, retrying")
	}

	lf.filters = filters
	lf.filterDirty = true
	lf.filtersLoaded = true
	return nil
}

// Logs returns logs matching topics and address (exactly) in the given block range,
// which are canonical at time of query.
func (lf *logFilterer) Logs(start, end int64, eventSig common.Hash, address common.Address, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectLogsByBlockRangeFilter(start, end, address, eventSig, qopts...)
}

func (lf *logFilterer) LogsWithSigs(start, end int64, eventSigs []common.Hash, address common.Address, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectLogsWithSigsByBlockRangeFilter(start, end, address, eventSigs, qopts...)
}

func (lf *logFilterer) LogsCreatedAfter(eventSig common.Hash, address common.Address, after time.Time, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectLogsCreatedAfter(eventSig[:], address, after, confs, qopts...)
}

// IndexedLogs finds all the logs that have a topic value in topicValues at index topicIndex.
func (lf *logFilterer) IndexedLogs(eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexedLogs(address, eventSig, topicIndex, topicValues, confs, qopts...)
}

// IndexedLogsByBlockRange finds all the logs that have a topic value in topicValues at index topicIndex within the block range
func (lf *logFilterer) IndexedLogsByBlockRange(start, end int64, eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexedLogsByBlockRangeFilter(start, end, address, eventSig, topicIndex, topicValues, qopts...)
}

func (lf *logFilterer) IndexedLogsCreatedAfter(eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, after time.Time, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexedLogsCreatedAfter(address, eventSig, topicIndex, topicValues, after, confs, qopts...)
}

// LogsDataWordGreaterThan note index is 0 based.
func (lf *logFilterer) LogsDataWordGreaterThan(eventSig common.Hash, address common.Address, wordIndex int, wordValueMin common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectDataWordGreaterThan(address, eventSig, wordIndex, wordValueMin, confs, qopts...)
}

// LogsDataWordRange note index is 0 based.
func (lf *logFilterer) LogsDataWordRange(eventSig common.Hash, address common.Address, wordIndex int, wordValueMin, wordValueMax common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectDataWordRange(address, eventSig, wordIndex, wordValueMin, wordValueMax, confs, qopts...)
}

// IndexedLogsTopicGreaterThan finds all the logs that have a topic value greater than topicValueMin at index topicIndex.
// Only works for integer topics.
func (lf *logFilterer) IndexedLogsTopicGreaterThan(eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexLogsTopicGreaterThan(address, eventSig, topicIndex, topicValueMin, confs, qopts...)
}

func (lf *logFilterer) IndexedLogsTopicRange(eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, topicValueMax common.Hash, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexLogsTopicRange(address, eventSig, topicIndex, topicValueMin, topicValueMax, confs, qopts...)
}

func (lf *logFilterer) BlockByNumber(n int64, qopts ...pg.QOpt) (*LogPollerBlock, error) {
	return lf.orm.SelectBlockByNumber(n, qopts...)
}

// LatestLogByEventSigWithConfs finds the latest log that has confs number of blocks on top of the log.
func (lf *logFilterer) LatestLogByEventSigWithConfs(eventSig common.Hash, address common.Address, confs int, qopts ...pg.QOpt) (*Log, error) {
	return lf.orm.SelectLatestLogEventSigWithConfs(eventSig, address, confs, qopts...)
}

func (lf *logFilterer) LatestLogEventSigsAddrsWithConfs(fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectLatestLogEventSigsAddrsWithConfs(fromBlock, addresses, eventSigs, confs, qopts...)
}

func (lf *logFilterer) LatestBlockByEventSigsAddrsWithConfs(fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs int, qopts ...pg.QOpt) (int64, error) {
	return lf.orm.SelectLatestBlockNumberEventSigsAddrsWithConfs(fromBlock, eventSigs, addresses, confs, qopts...)
}

// IndexedLogsWithSigsExcluding returns the set difference(A-B) of logs with signature sigA and sigB, matching is done on the topics index
//
// For example, query to retrieve unfulfilled requests by querying request log events without matching fulfillment log events.
// The order of events is not significant. Both logs must be inside the block range and have the minimum number of confirmations
func (lf *logFilterer) IndexedLogsWithSigsExcluding(address common.Address, eventSigA, eventSigB common.Hash, topicIndex int, fromBlock, toBlock int64, confs int, qopts ...pg.QOpt) ([]Log, error) {
	return lf.orm.SelectIndexedLogsWithSigsExcluding(eventSigA, eventSigB, topicIndex, address, fromBlock, toBlock, confs, qopts...)
}

func EvmWord(i uint64) common.Hash {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return common.BytesToHash(b)
}
