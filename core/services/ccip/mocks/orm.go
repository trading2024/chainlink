// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ccip "github.com/smartcontractkit/chainlink/v2/core/services/ccip"

	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// ClearGasPricesByDestChain provides a mock function with given fields: ctx, destChainSelector, expireSec
func (_m *ORM) ClearGasPricesByDestChain(ctx context.Context, destChainSelector uint64, expireSec int) error {
	ret := _m.Called(ctx, destChainSelector, expireSec)

	if len(ret) == 0 {
		panic("no return value specified for ClearGasPricesByDestChain")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int) error); ok {
		r0 = rf(ctx, destChainSelector, expireSec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearTokenPricesByDestChain provides a mock function with given fields: ctx, destChainSelector, expireSec
func (_m *ORM) ClearTokenPricesByDestChain(ctx context.Context, destChainSelector uint64, expireSec int) error {
	ret := _m.Called(ctx, destChainSelector, expireSec)

	if len(ret) == 0 {
		panic("no return value specified for ClearTokenPricesByDestChain")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int) error); ok {
		r0 = rf(ctx, destChainSelector, expireSec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGasPricesByDestChain provides a mock function with given fields: ctx, destChainSelector
func (_m *ORM) GetGasPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]ccip.GasPrice, error) {
	ret := _m.Called(ctx, destChainSelector)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPricesByDestChain")
	}

	var r0 []ccip.GasPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]ccip.GasPrice, error)); ok {
		return rf(ctx, destChainSelector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []ccip.GasPrice); ok {
		r0 = rf(ctx, destChainSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.GasPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, destChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenPricesByDestChain provides a mock function with given fields: ctx, destChainSelector
func (_m *ORM) GetTokenPricesByDestChain(ctx context.Context, destChainSelector uint64) ([]ccip.TokenPrice, error) {
	ret := _m.Called(ctx, destChainSelector)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenPricesByDestChain")
	}

	var r0 []ccip.TokenPrice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) ([]ccip.TokenPrice, error)); ok {
		return rf(ctx, destChainSelector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []ccip.TokenPrice); ok {
		r0 = rf(ctx, destChainSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.TokenPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, destChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertGasPricesForDestChain provides a mock function with given fields: ctx, destChainSelector, jobId, gasPrices
func (_m *ORM) InsertGasPricesForDestChain(ctx context.Context, destChainSelector uint64, jobId int32, gasPrices []ccip.GasPriceUpdate) error {
	ret := _m.Called(ctx, destChainSelector, jobId, gasPrices)

	if len(ret) == 0 {
		panic("no return value specified for InsertGasPricesForDestChain")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int32, []ccip.GasPriceUpdate) error); ok {
		r0 = rf(ctx, destChainSelector, jobId, gasPrices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertTokenPricesForDestChain provides a mock function with given fields: ctx, destChainSelector, jobId, tokenPrices
func (_m *ORM) InsertTokenPricesForDestChain(ctx context.Context, destChainSelector uint64, jobId int32, tokenPrices []ccip.TokenPriceUpdate) error {
	ret := _m.Called(ctx, destChainSelector, jobId, tokenPrices)

	if len(ret) == 0 {
		panic("no return value specified for InsertTokenPricesForDestChain")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, int32, []ccip.TokenPriceUpdate) error); ok {
		r0 = rf(ctx, destChainSelector, jobId, tokenPrices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
