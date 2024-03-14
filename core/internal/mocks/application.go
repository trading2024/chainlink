// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	audit "github.com/smartcontractkit/chainlink/v2/core/logger/audit"

	bridges "github.com/smartcontractkit/chainlink/v2/core/bridges"

	chainlink "github.com/smartcontractkit/chainlink/v2/core/services/chainlink"

	context "context"

	feeds "github.com/smartcontractkit/chainlink/v2/core/services/feeds"

	job "github.com/smartcontractkit/chainlink/v2/core/services/job"

	keystore "github.com/smartcontractkit/chainlink/v2/core/services/keystore"

	logger "github.com/smartcontractkit/chainlink/v2/core/logger"

	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/smartcontractkit/chainlink/v2/core/services/pipeline"

	plugins "github.com/smartcontractkit/chainlink/v2/plugins"

	services "github.com/smartcontractkit/chainlink/v2/core/services"

	sessions "github.com/smartcontractkit/chainlink/v2/core/sessions"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	sqlx "github.com/jmoiron/sqlx"

	txmgr "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"

	types "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"

	uuid "github.com/google/uuid"

	webhook "github.com/smartcontractkit/chainlink/v2/core/services/webhook"

	zapcore "go.uber.org/zap/zapcore"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// AddJobV2 provides a mock function with given fields: ctx, _a1
func (_m *Application) AddJobV2(ctx context.Context, _a1 *job.Job) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddJobV2")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.Job) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider provides a mock function with given fields:
func (_m *Application) AuthenticationProvider() sessions.AuthenticationProvider {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AuthenticationProvider")
	}

	var r0 sessions.AuthenticationProvider
	if rf, ok := ret.Get(0).(func() sessions.AuthenticationProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sessions.AuthenticationProvider)
		}
	}

	return r0
}

// BasicAdminUsersORM provides a mock function with given fields:
func (_m *Application) BasicAdminUsersORM() sessions.BasicAdminUsersORM {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BasicAdminUsersORM")
	}

	var r0 sessions.BasicAdminUsersORM
	if rf, ok := ret.Get(0).(func() sessions.BasicAdminUsersORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sessions.BasicAdminUsersORM)
		}
	}

	return r0
}

// BridgeORM provides a mock function with given fields:
func (_m *Application) BridgeORM() bridges.ORM {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BridgeORM")
	}

	var r0 bridges.ORM
	if rf, ok := ret.Get(0).(func() bridges.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bridges.ORM)
		}
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, jobID
func (_m *Application) DeleteJob(ctx context.Context, jobID int32) error {
	ret := _m.Called(ctx, jobID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EVMORM provides a mock function with given fields:
func (_m *Application) EVMORM() types.Configs {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EVMORM")
	}

	var r0 types.Configs
	if rf, ok := ret.Get(0).(func() types.Configs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Configs)
		}
	}

	return r0
}

// GetAuditLogger provides a mock function with given fields:
func (_m *Application) GetAuditLogger() audit.AuditLogger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuditLogger")
	}

	var r0 audit.AuditLogger
	if rf, ok := ret.Get(0).(func() audit.AuditLogger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(audit.AuditLogger)
		}
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *Application) GetConfig() chainlink.GeneralConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConfig")
	}

	var r0 chainlink.GeneralConfig
	if rf, ok := ret.Get(0).(func() chainlink.GeneralConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chainlink.GeneralConfig)
		}
	}

	return r0
}

// GetDB provides a mock function with given fields:
func (_m *Application) GetDB() sqlutil.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDB")
	}

	var r0 sqlutil.DB
	if rf, ok := ret.Get(0).(func() sqlutil.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sqlutil.DB)
		}
	}

	return r0
}

// GetExternalInitiatorManager provides a mock function with given fields:
func (_m *Application) GetExternalInitiatorManager() webhook.ExternalInitiatorManager {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExternalInitiatorManager")
	}

	var r0 webhook.ExternalInitiatorManager
	if rf, ok := ret.Get(0).(func() webhook.ExternalInitiatorManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(webhook.ExternalInitiatorManager)
		}
	}

	return r0
}

// GetFeedsService provides a mock function with given fields:
func (_m *Application) GetFeedsService() feeds.Service {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFeedsService")
	}

	var r0 feeds.Service
	if rf, ok := ret.Get(0).(func() feeds.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(feeds.Service)
		}
	}

	return r0
}

// GetHealthChecker provides a mock function with given fields:
func (_m *Application) GetHealthChecker() services.Checker {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHealthChecker")
	}

	var r0 services.Checker
	if rf, ok := ret.Get(0).(func() services.Checker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.Checker)
		}
	}

	return r0
}

// GetKeyStore provides a mock function with given fields:
func (_m *Application) GetKeyStore() keystore.Master {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetKeyStore")
	}

	var r0 keystore.Master
	if rf, ok := ret.Get(0).(func() keystore.Master); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Master)
		}
	}

	return r0
}

// GetLogger provides a mock function with given fields:
func (_m *Application) GetLogger() logger.SugaredLogger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLogger")
	}

	var r0 logger.SugaredLogger
	if rf, ok := ret.Get(0).(func() logger.SugaredLogger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.SugaredLogger)
		}
	}

	return r0
}

// GetLoopRegistry provides a mock function with given fields:
func (_m *Application) GetLoopRegistry() *plugins.LoopRegistry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLoopRegistry")
	}

	var r0 *plugins.LoopRegistry
	if rf, ok := ret.Get(0).(func() *plugins.LoopRegistry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*plugins.LoopRegistry)
		}
	}

	return r0
}

// GetRelayers provides a mock function with given fields:
func (_m *Application) GetRelayers() chainlink.RelayerChainInteroperators {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRelayers")
	}

	var r0 chainlink.RelayerChainInteroperators
	if rf, ok := ret.Get(0).(func() chainlink.RelayerChainInteroperators); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chainlink.RelayerChainInteroperators)
		}
	}

	return r0
}

// GetSqlxDB provides a mock function with given fields:
func (_m *Application) GetSqlxDB() *sqlx.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSqlxDB")
	}

	var r0 *sqlx.DB
	if rf, ok := ret.Get(0).(func() *sqlx.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.DB)
		}
	}

	return r0
}

// GetWebAuthnConfiguration provides a mock function with given fields:
func (_m *Application) GetWebAuthnConfiguration() sessions.WebAuthnConfiguration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWebAuthnConfiguration")
	}

	var r0 sessions.WebAuthnConfiguration
	if rf, ok := ret.Get(0).(func() sessions.WebAuthnConfiguration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sessions.WebAuthnConfiguration)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Application) ID() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// JobORM provides a mock function with given fields:
func (_m *Application) JobORM() job.ORM {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JobORM")
	}

	var r0 job.ORM
	if rf, ok := ret.Get(0).(func() job.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.ORM)
		}
	}

	return r0
}

// JobSpawner provides a mock function with given fields:
func (_m *Application) JobSpawner() job.Spawner {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JobSpawner")
	}

	var r0 job.Spawner
	if rf, ok := ret.Get(0).(func() job.Spawner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.Spawner)
		}
	}

	return r0
}

// PipelineORM provides a mock function with given fields:
func (_m *Application) PipelineORM() pipeline.ORM {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PipelineORM")
	}

	var r0 pipeline.ORM
	if rf, ok := ret.Get(0).(func() pipeline.ORM); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pipeline.ORM)
		}
	}

	return r0
}

// ReplayFromBlock provides a mock function with given fields: chainID, number, forceBroadcast
func (_m *Application) ReplayFromBlock(chainID *big.Int, number uint64, forceBroadcast bool) error {
	ret := _m.Called(chainID, number, forceBroadcast)

	if len(ret) == 0 {
		panic("no return value specified for ReplayFromBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, uint64, bool) error); ok {
		r0 = rf(chainID, number, forceBroadcast)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResumeJobV2 provides a mock function with given fields: ctx, taskID, result
func (_m *Application) ResumeJobV2(ctx context.Context, taskID uuid.UUID, result pipeline.Result) error {
	ret := _m.Called(ctx, taskID, result)

	if len(ret) == 0 {
		panic("no return value specified for ResumeJobV2")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, pipeline.Result) error); ok {
		r0 = rf(ctx, taskID, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunJobV2 provides a mock function with given fields: ctx, jobID, meta
func (_m *Application) RunJobV2(ctx context.Context, jobID int32, meta map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, jobID, meta)

	if len(ret) == 0 {
		panic("no return value specified for RunJobV2")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, map[string]interface{}) (int64, error)); ok {
		return rf(ctx, jobID, meta)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, map[string]interface{}) int64); ok {
		r0 = rf(ctx, jobID, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, map[string]interface{}) error); ok {
		r1 = rf(ctx, jobID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunWebhookJobV2 provides a mock function with given fields: ctx, jobUUID, requestBody, meta
func (_m *Application) RunWebhookJobV2(ctx context.Context, jobUUID uuid.UUID, requestBody string, meta pipeline.JSONSerializable) (int64, error) {
	ret := _m.Called(ctx, jobUUID, requestBody, meta)

	if len(ret) == 0 {
		panic("no return value specified for RunWebhookJobV2")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, pipeline.JSONSerializable) (int64, error)); ok {
		return rf(ctx, jobUUID, requestBody, meta)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, pipeline.JSONSerializable) int64); ok {
		r0 = rf(ctx, jobUUID, requestBody, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string, pipeline.JSONSerializable) error); ok {
		r1 = rf(ctx, jobUUID, requestBody, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecretGenerator provides a mock function with given fields:
func (_m *Application) SecretGenerator() chainlink.SecretGenerator {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SecretGenerator")
	}

	var r0 chainlink.SecretGenerator
	if rf, ok := ret.Get(0).(func() chainlink.SecretGenerator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chainlink.SecretGenerator)
		}
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: lvl
func (_m *Application) SetLogLevel(lvl zapcore.Level) error {
	ret := _m.Called(lvl)

	if len(ret) == 0 {
		panic("no return value specified for SetLogLevel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(zapcore.Level) error); ok {
		r0 = rf(lvl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Application) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Application) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxmStorageService provides a mock function with given fields:
func (_m *Application) TxmStorageService() txmgr.EvmTxStore {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxmStorageService")
	}

	var r0 txmgr.EvmTxStore
	if rf, ok := ret.Get(0).(func() txmgr.EvmTxStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(txmgr.EvmTxStore)
		}
	}

	return r0
}

// WakeSessionReaper provides a mock function with given fields:
func (_m *Application) WakeSessionReaper() {
	_m.Called()
}

// NewApplication creates a new instance of Application. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApplication(t interface {
	mock.TestingT
	Cleanup(func())
}) *Application {
	mock := &Application{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
