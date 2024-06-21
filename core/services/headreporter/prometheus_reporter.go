package headreporter

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	txmgrcommon "github.com/smartcontractkit/chainlink/v2/common/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"go.uber.org/multierr"
	"math/big"
	"time"
)

type (
	prometheusReporter struct {
		ds      sqlutil.DataSource
		chains  legacyevm.LegacyChainContainer
		lggr    logger.Logger
		backend PrometheusBackend
	}

	PrometheusBackend interface {
		SetUnconfirmedTransactions(*big.Int, int64)
		SetMaxUnconfirmedAge(*big.Int, float64)
		SetMaxUnconfirmedBlocks(*big.Int, int64)
		SetPipelineRunsQueued(n int)
		SetPipelineTaskRunsQueued(n int)
	}

	defaultBackend struct{}
)

var (
	promUnconfirmedTransactions = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "unconfirmed_transactions",
		Help: "Number of currently unconfirmed transactions",
	}, []string{"evmChainID"})
	promMaxUnconfirmedAge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "max_unconfirmed_tx_age",
		Help: "The length of time the oldest unconfirmed transaction has been in that state (in seconds). Will be 0 if there are no unconfirmed transactions.",
	}, []string{"evmChainID"})
	promMaxUnconfirmedBlocks = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "max_unconfirmed_blocks",
		Help: "The max number of blocks any currently unconfirmed transaction has been unconfirmed for",
	}, []string{"evmChainID"})
	promPipelineRunsQueued = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pipeline_runs_queued",
		Help: "The total number of pipeline runs that are awaiting execution",
	})
	promPipelineTaskRunsQueued = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pipeline_task_runs_queued",
		Help: "The total number of pipeline task runs that are awaiting execution",
	})
)

func NewPrometheusReporter(ds sqlutil.DataSource, chainContainer legacyevm.LegacyChainContainer, lggr logger.Logger, opts ...interface{}) *prometheusReporter {
	var backend PrometheusBackend = defaultBackend{}
	for _, opt := range opts {
		switch v := opt.(type) {
		case PrometheusBackend:
			backend = v
		}
	}
	return &prometheusReporter{
		ds:      ds,
		chains:  chainContainer,
		lggr:    lggr.Named(name),
		backend: backend,
	}
}

func (pr *prometheusReporter) getTxm(evmChainID *big.Int) (txmgr.TxManager, error) {
	chain, err := pr.chains.Get(evmChainID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain: %w", err)
	}
	return chain.TxManager(), nil
}

func (pr *prometheusReporter) reportOnHead(ctx context.Context, head *evmtypes.Head) {
	evmChainID := head.EVMChainID.ToInt()
	err := multierr.Combine(
		errors.Wrap(pr.reportPendingEthTxes(ctx, evmChainID), "reportPendingEthTxes failed"),
		errors.Wrap(pr.reportMaxUnconfirmedAge(ctx, evmChainID), "reportMaxUnconfirmedAge failed"),
		errors.Wrap(pr.reportMaxUnconfirmedBlocks(ctx, head), "reportMaxUnconfirmedBlocks failed"),
	)

	if err != nil && ctx.Err() == nil {
		pr.lggr.Errorw("Error reporting prometheus metrics", "err", err)
	}
}

func (pr *prometheusReporter) reportPendingEthTxes(ctx context.Context, evmChainID *big.Int) (err error) {
	txm, err := pr.getTxm(evmChainID)
	if err != nil {
		return fmt.Errorf("failed to get txm: %w", err)
	}

	unconfirmed, err := txm.CountTransactionsByState(ctx, txmgrcommon.TxUnconfirmed)
	if err != nil {
		return fmt.Errorf("failed to query for unconfirmed eth_tx count: %w", err)
	}
	pr.backend.SetUnconfirmedTransactions(evmChainID, int64(unconfirmed))
	return nil
}

func (pr *prometheusReporter) reportMaxUnconfirmedAge(ctx context.Context, evmChainID *big.Int) (err error) {
	txm, err := pr.getTxm(evmChainID)
	if err != nil {
		return fmt.Errorf("failed to get txm: %w", err)
	}

	broadcastAt, err := txm.FindEarliestUnconfirmedBroadcastTime(ctx)
	if err != nil {
		return fmt.Errorf("failed to query for min broadcast time: %w", err)
	}

	var seconds float64
	if broadcastAt.Valid {
		seconds = time.Since(broadcastAt.ValueOrZero()).Seconds()
	}
	pr.backend.SetMaxUnconfirmedAge(evmChainID, seconds)
	return nil
}

func (pr *prometheusReporter) reportMaxUnconfirmedBlocks(ctx context.Context, head *evmtypes.Head) (err error) {
	txm, err := pr.getTxm(head.EVMChainID.ToInt())
	if err != nil {
		return fmt.Errorf("failed to get txm: %w", err)
	}

	earliestUnconfirmedTxBlock, err := txm.FindEarliestUnconfirmedTxAttemptBlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to query for earliest unconfirmed tx block: %w", err)
	}

	var blocksUnconfirmed int64
	if !earliestUnconfirmedTxBlock.IsZero() {
		blocksUnconfirmed = head.Number - earliestUnconfirmedTxBlock.ValueOrZero()
	}
	pr.backend.SetMaxUnconfirmedBlocks(head.EVMChainID.ToInt(), blocksUnconfirmed)
	return nil
}

func (pr *prometheusReporter) reportPeriodic(ctx context.Context) {
	if err := errors.Wrap(pr.reportPipelineRunStats(ctx), "reportPipelineRunStats failed"); err != nil {
		pr.lggr.Errorw("Error reporting prometheus metrics", "err", err)
	}
}

func (pr *prometheusReporter) reportPipelineRunStats(ctx context.Context) (err error) {
	rows, err := pr.ds.QueryContext(ctx, `
SELECT pipeline_run_id FROM pipeline_task_runs WHERE finished_at IS NULL
`)
	if err != nil {
		return errors.Wrap(err, "failed to query for pipeline_run_id")
	}
	defer func() {
		err = multierr.Combine(err, rows.Close())
	}()

	pipelineTaskRunsQueued := 0
	pipelineRunsQueuedSet := make(map[int32]struct{})
	for rows.Next() {
		var pipelineRunID int32
		if err = rows.Scan(&pipelineRunID); err != nil {
			return errors.Wrap(err, "unexpected error scanning row")
		}
		pipelineTaskRunsQueued++
		pipelineRunsQueuedSet[pipelineRunID] = struct{}{}
	}
	if err = rows.Err(); err != nil {
		return err
	}
	pipelineRunsQueued := len(pipelineRunsQueuedSet)

	pr.backend.SetPipelineTaskRunsQueued(pipelineTaskRunsQueued)
	pr.backend.SetPipelineRunsQueued(pipelineRunsQueued)

	return nil
}

func (defaultBackend) SetUnconfirmedTransactions(evmChainID *big.Int, n int64) {
	promUnconfirmedTransactions.WithLabelValues(evmChainID.String()).Set(float64(n))
}

func (defaultBackend) SetMaxUnconfirmedAge(evmChainID *big.Int, s float64) {
	promMaxUnconfirmedAge.WithLabelValues(evmChainID.String()).Set(s)
}

func (defaultBackend) SetMaxUnconfirmedBlocks(evmChainID *big.Int, n int64) {
	promMaxUnconfirmedBlocks.WithLabelValues(evmChainID.String()).Set(float64(n))
}

func (defaultBackend) SetPipelineRunsQueued(n int) {
	promPipelineTaskRunsQueued.Set(float64(n))
}

func (defaultBackend) SetPipelineTaskRunsQueued(n int) {
	promPipelineRunsQueued.Set(float64(n))
}