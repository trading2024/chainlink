// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"
	mock "github.com/stretchr/testify/mock"

	time "time"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"

	types "github.com/smartcontractkit/chainlink/v2/common/types"

	uuid "github.com/google/uuid"
)

// TxStore is an autogenerated mock type for the TxStore type
type TxStore[ADDR types.Hashable, CHAIN_ID types.ID, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee] struct {
	mock.Mock
}

// Abandon provides a mock function with given fields: ctx, id, addr
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Abandon(ctx context.Context, id CHAIN_ID, addr ADDR) error {
	ret := _m.Called(ctx, id, addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID, ADDR) error); ok {
		r0 = rf(ctx, id, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckTxQueueCapacity provides a mock function with given fields: ctx, fromAddress, maxQueuedTransactions, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CheckTxQueueCapacity(ctx context.Context, fromAddress ADDR, maxQueuedTransactions uint64, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, fromAddress, maxQueuedTransactions, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, uint64, CHAIN_ID) error); ok {
		r0 = rf(ctx, fromAddress, maxQueuedTransactions, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Close() {
	_m.Called()
}

// CountUnconfirmedTransactions provides a mock function with given fields: ctx, fromAddress, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnconfirmedTransactions(ctx context.Context, fromAddress ADDR, chainID CHAIN_ID) (uint32, error) {
	ret := _m.Called(ctx, fromAddress, chainID)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) (uint32, error)); ok {
		return rf(ctx, fromAddress, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) uint32); ok {
		r0 = rf(ctx, fromAddress, chainID)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, fromAddress, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountUnstartedTransactions provides a mock function with given fields: ctx, fromAddress, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnstartedTransactions(ctx context.Context, fromAddress ADDR, chainID CHAIN_ID) (uint32, error) {
	ret := _m.Called(ctx, fromAddress, chainID)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) (uint32, error)); ok {
		return rf(ctx, fromAddress, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) uint32); ok {
		r0 = rf(ctx, fromAddress, chainID)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, fromAddress, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransaction provides a mock function with given fields: ctx, txRequest, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CreateTransaction(ctx context.Context, txRequest txmgrtypes.TxRequest[ADDR, TX_HASH], chainID CHAIN_ID) (txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, txRequest, chainID)

	var r0 txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID) (txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, txRequest, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID) txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, txRequest, chainID)
	} else {
		r0 = ret.Get(0).(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(context.Context, txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID) error); ok {
		r1 = rf(ctx, txRequest, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInProgressAttempt provides a mock function with given fields: ctx, attempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) DeleteInProgressAttempt(ctx context.Context, attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindLatestSequence provides a mock function with given fields: ctx, fromAddress, chainId
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindLatestSequence(ctx context.Context, fromAddress ADDR, chainId CHAIN_ID) (SEQ, error) {
	ret := _m.Called(ctx, fromAddress, chainId)

	var r0 SEQ
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) (SEQ, error)); ok {
		return rf(ctx, fromAddress, chainId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) SEQ); ok {
		r0 = rf(ctx, fromAddress, chainId)
	} else {
		r0 = ret.Get(0).(SEQ)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, fromAddress, chainId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNextUnstartedTransactionFromAddress provides a mock function with given fields: ctx, etx, fromAddress, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindNextUnstartedTransactionFromAddress(ctx context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], fromAddress ADDR, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, etx, fromAddress, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ADDR, CHAIN_ID) error); ok {
		r0 = rf(ctx, etx, fromAddress, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTransactionsConfirmedInBlockRange provides a mock function with given fields: ctx, highBlockNumber, lowBlockNumber, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTransactionsConfirmedInBlockRange(ctx context.Context, highBlockNumber int64, lowBlockNumber int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, highBlockNumber, lowBlockNumber, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, highBlockNumber, lowBlockNumber, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, highBlockNumber, lowBlockNumber, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, highBlockNumber, lowBlockNumber, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsConfirmedMissingReceipt provides a mock function with given fields: ctx, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsConfirmedMissingReceipt(ctx context.Context, chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, CHAIN_ID) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsRequiringReceiptFetch provides a mock function with given fields: ctx, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringReceiptFetch(ctx context.Context, chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, CHAIN_ID) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsRequiringResend provides a mock function with given fields: ctx, olderThan, maxInFlightTransactions, chainID, address
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringResend(ctx context.Context, olderThan time.Time, maxInFlightTransactions uint32, chainID CHAIN_ID, address ADDR) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, olderThan, maxInFlightTransactions, chainID, address)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32, CHAIN_ID, ADDR) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, olderThan, maxInFlightTransactions, chainID, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32, CHAIN_ID, ADDR) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, olderThan, maxInFlightTransactions, chainID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, uint32, CHAIN_ID, ADDR) error); ok {
		r1 = rf(ctx, olderThan, maxInFlightTransactions, chainID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxWithIdempotencyKey provides a mock function with given fields: ctx, idempotencyKey, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxWithIdempotencyKey(ctx context.Context, idempotencyKey string, chainID CHAIN_ID) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, idempotencyKey, chainID)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, CHAIN_ID) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, idempotencyKey, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, CHAIN_ID) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, idempotencyKey, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, CHAIN_ID) error); ok {
		r1 = rf(ctx, idempotencyKey, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxWithSequence provides a mock function with given fields: ctx, fromAddress, seq
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxWithSequence(ctx context.Context, fromAddress ADDR, seq SEQ) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, fromAddress, seq)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, SEQ) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, fromAddress, seq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, SEQ) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, fromAddress, seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, SEQ) error); ok {
		r1 = rf(ctx, fromAddress, seq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxesByMetaFieldAndStates provides a mock function with given fields: ctx, metaField, metaValue, states, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesByMetaFieldAndStates(ctx context.Context, metaField string, metaValue string, states []txmgrtypes.TxState, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, metaField, metaValue, states, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []txmgrtypes.TxState, *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, metaField, metaValue, states, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []txmgrtypes.TxState, *big.Int) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, metaField, metaValue, states, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []txmgrtypes.TxState, *big.Int) error); ok {
		r1 = rf(ctx, metaField, metaValue, states, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxesPendingCallback provides a mock function with given fields: ctx, blockNum, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesPendingCallback(ctx context.Context, blockNum int64, chainID CHAIN_ID) ([]txmgrtypes.ReceiptPlus[R], error) {
	ret := _m.Called(ctx, blockNum, chainID)

	var r0 []txmgrtypes.ReceiptPlus[R]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAIN_ID) ([]txmgrtypes.ReceiptPlus[R], error)); ok {
		return rf(ctx, blockNum, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAIN_ID) []txmgrtypes.ReceiptPlus[R]); ok {
		r0 = rf(ctx, blockNum, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.ReceiptPlus[R])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, blockNum, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxesWithAttemptsAndReceiptsByIdsAndState provides a mock function with given fields: ctx, ids, states, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithAttemptsAndReceiptsByIdsAndState(ctx context.Context, ids []big.Int, states []txmgrtypes.TxState, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, ids, states, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []big.Int, []txmgrtypes.TxState, *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, ids, states, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []big.Int, []txmgrtypes.TxState, *big.Int) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, ids, states, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []big.Int, []txmgrtypes.TxState, *big.Int) error); ok {
		r1 = rf(ctx, ids, states, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxesWithMetaFieldByReceiptBlockNum provides a mock function with given fields: ctx, metaField, blockNum, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithMetaFieldByReceiptBlockNum(ctx context.Context, metaField string, blockNum int64, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, metaField, blockNum, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, metaField, blockNum, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, *big.Int) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, metaField, blockNum, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, *big.Int) error); ok {
		r1 = rf(ctx, metaField, blockNum, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxesWithMetaFieldByStates provides a mock function with given fields: ctx, metaField, states, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxesWithMetaFieldByStates(ctx context.Context, metaField string, states []txmgrtypes.TxState, chainID *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, metaField, states, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []txmgrtypes.TxState, *big.Int) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, metaField, states, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []txmgrtypes.TxState, *big.Int) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, metaField, states, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []txmgrtypes.TxState, *big.Int) error); ok {
		r1 = rf(ctx, metaField, states, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxsRequiringGasBump provides a mock function with given fields: ctx, address, blockNum, gasBumpThreshold, depth, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringGasBump(ctx context.Context, address ADDR, blockNum int64, gasBumpThreshold int64, depth int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, address, blockNum, gasBumpThreshold, depth, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxsRequiringResubmissionDueToInsufficientFunds provides a mock function with given fields: ctx, address, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringResubmissionDueToInsufficientFunds(ctx context.Context, address ADDR, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, address, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, address, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, address, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInProgressTxAttempts provides a mock function with given fields: ctx, address, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetInProgressTxAttempts(ctx context.Context, address ADDR, chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, address, chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, address, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, address, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonFatalTransactions provides a mock function with given fields: ctx, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetNonFatalTransactions(ctx context.Context, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, CHAIN_ID) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxByID provides a mock function with given fields: ctx, id
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetTxByID(ctx context.Context, id int64) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, id)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxInProgress provides a mock function with given fields: ctx, fromAddress
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetTxInProgress(ctx context.Context, fromAddress ADDR) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, fromAddress)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, fromAddress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, fromAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR) error); ok {
		r1 = rf(ctx, fromAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasInProgressTransaction provides a mock function with given fields: ctx, account, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) HasInProgressTransaction(ctx context.Context, account ADDR, chainID CHAIN_ID) (bool, error) {
	ret := _m.Called(ctx, account, chainID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) (bool, error)); ok {
		return rf(ctx, account, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) bool); ok {
		r0 = rf(ctx, account, chainID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, account, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTxFinalized provides a mock function with given fields: ctx, blockHeight, txID, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) IsTxFinalized(ctx context.Context, blockHeight int64, txID int64, chainID CHAIN_ID) (bool, error) {
	ret := _m.Called(ctx, blockHeight, txID, chainID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, CHAIN_ID) (bool, error)); ok {
		return rf(ctx, blockHeight, txID, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, CHAIN_ID) bool); ok {
		r0 = rf(ctx, blockHeight, txID, chainID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, blockHeight, txID, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadTxAttempts provides a mock function with given fields: ctx, etx
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) LoadTxAttempts(ctx context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, etx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, etx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllConfirmedMissingReceipt provides a mock function with given fields: ctx, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkAllConfirmedMissingReceipt(ctx context.Context, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CHAIN_ID) error); ok {
		r0 = rf(ctx, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkOldTxesMissingReceiptAsErrored provides a mock function with given fields: ctx, blockNum, finalityDepth, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkOldTxesMissingReceiptAsErrored(ctx context.Context, blockNum int64, finalityDepth uint32, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, blockNum, finalityDepth, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint32, CHAIN_ID) error); ok {
		r0 = rf(ctx, blockNum, finalityDepth, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreloadTxes provides a mock function with given fields: ctx, attempts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PreloadTxes(ctx context.Context, attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, attempts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, attempts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PruneUnstartedTxQueue provides a mock function with given fields: ctx, queueSize, subject
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PruneUnstartedTxQueue(ctx context.Context, queueSize uint32, subject uuid.UUID) (int64, error) {
	ret := _m.Called(ctx, queueSize, subject)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uuid.UUID) (int64, error)); ok {
		return rf(ctx, queueSize, subject)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uuid.UUID) int64); ok {
		r0 = rf(ctx, queueSize, subject)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uuid.UUID) error); ok {
		r1 = rf(ctx, queueSize, subject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReapTxHistory provides a mock function with given fields: ctx, minBlockNumberToKeep, timeThreshold, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) ReapTxHistory(ctx context.Context, minBlockNumberToKeep int64, timeThreshold time.Time, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, minBlockNumberToKeep, timeThreshold, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time, CHAIN_ID) error); ok {
		r0 = rf(ctx, minBlockNumberToKeep, timeThreshold, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveConfirmedMissingReceiptAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveConfirmedMissingReceiptAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveFetchedReceipts provides a mock function with given fields: ctx, receipts, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveFetchedReceipts(ctx context.Context, receipts []R, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, receipts, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []R, CHAIN_ID) error); ok {
		r0 = rf(ctx, receipts, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInProgressAttempt provides a mock function with given fields: ctx, attempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInProgressAttempt(ctx context.Context, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInsufficientFundsAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInsufficientFundsAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveReplacementInProgressAttempt provides a mock function with given fields: ctx, oldAttempt, replacementAttempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveReplacementInProgressAttempt(ctx context.Context, oldAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], replacementAttempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, oldAttempt, replacementAttempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, oldAttempt, replacementAttempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSentAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveSentAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBroadcastBeforeBlockNum provides a mock function with given fields: ctx, blockNum, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SetBroadcastBeforeBlockNum(ctx context.Context, blockNum int64, chainID CHAIN_ID) error {
	ret := _m.Called(ctx, blockNum, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAIN_ID) error); ok {
		r0 = rf(ctx, blockNum, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBroadcastAts provides a mock function with given fields: ctx, now, etxIDs
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateBroadcastAts(ctx context.Context, now time.Time, etxIDs []int64) error {
	ret := _m.Called(ctx, now, etxIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, []int64) error); ok {
		r0 = rf(ctx, now, etxIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxAttemptInProgressToBroadcast provides a mock function with given fields: ctx, etx, attempt, NewAttemptState
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxAttemptInProgressToBroadcast(ctx context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], NewAttemptState txmgrtypes.TxAttemptState) error {
	ret := _m.Called(ctx, etx, attempt, NewAttemptState)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttemptState) error); ok {
		r0 = rf(ctx, etx, attempt, NewAttemptState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxCallbackCompleted provides a mock function with given fields: ctx, pipelineTaskRunRid, chainId
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxCallbackCompleted(ctx context.Context, pipelineTaskRunRid uuid.UUID, chainId CHAIN_ID) error {
	ret := _m.Called(ctx, pipelineTaskRunRid, chainId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, CHAIN_ID) error); ok {
		r0 = rf(ctx, pipelineTaskRunRid, chainId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxFatalError provides a mock function with given fields: ctx, etx
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxFatalError(ctx context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, etx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, etx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxForRebroadcast provides a mock function with given fields: ctx, etx, etxAttempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxForRebroadcast(ctx context.Context, etx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], etxAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, etx, etxAttempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, etx, etxAttempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxUnstartedToInProgress provides a mock function with given fields: ctx, etx, attempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxUnstartedToInProgress(ctx context.Context, etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, etx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, etx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxsUnconfirmed provides a mock function with given fields: ctx, ids
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxsUnconfirmed(ctx context.Context, ids []int64) error {
	ret := _m.Called(ctx, ids)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) error); ok {
		r0 = rf(ctx, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTxStore creates a new instance of TxStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxStore[ADDR types.Hashable, CHAIN_ID types.ID, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee](t interface {
	mock.TestingT
	Cleanup(func())
}) *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE] {
	mock := &TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
