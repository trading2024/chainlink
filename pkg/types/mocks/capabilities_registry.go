// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	capabilities "github.com/smartcontractkit/chainlink-common/pkg/capabilities"

	mock "github.com/stretchr/testify/mock"
)

// CapabilitiesRegistry is an autogenerated mock type for the CapabilitiesRegistry type
type CapabilitiesRegistry struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, c
func (_m *CapabilitiesRegistry) Add(ctx context.Context, c capabilities.BaseCapability) error {
	ret := _m.Called(ctx, c)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, capabilities.BaseCapability) error); ok {
		r0 = rf(ctx, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, ID
func (_m *CapabilitiesRegistry) Get(ctx context.Context, ID string) (capabilities.BaseCapability, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 capabilities.BaseCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (capabilities.BaseCapability, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) capabilities.BaseCapability); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(capabilities.BaseCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAction provides a mock function with given fields: ctx, ID
func (_m *CapabilitiesRegistry) GetAction(ctx context.Context, ID string) (capabilities.ActionCapability, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetAction")
	}

	var r0 capabilities.ActionCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (capabilities.ActionCapability, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) capabilities.ActionCapability); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(capabilities.ActionCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsensus provides a mock function with given fields: ctx, ID
func (_m *CapabilitiesRegistry) GetConsensus(ctx context.Context, ID string) (capabilities.ConsensusCapability, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetConsensus")
	}

	var r0 capabilities.ConsensusCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (capabilities.ConsensusCapability, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) capabilities.ConsensusCapability); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(capabilities.ConsensusCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTarget provides a mock function with given fields: ctx, ID
func (_m *CapabilitiesRegistry) GetTarget(ctx context.Context, ID string) (capabilities.TargetCapability, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetTarget")
	}

	var r0 capabilities.TargetCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (capabilities.TargetCapability, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) capabilities.TargetCapability); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(capabilities.TargetCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrigger provides a mock function with given fields: ctx, ID
func (_m *CapabilitiesRegistry) GetTrigger(ctx context.Context, ID string) (capabilities.TriggerCapability, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetTrigger")
	}

	var r0 capabilities.TriggerCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (capabilities.TriggerCapability, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) capabilities.TriggerCapability); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(capabilities.TriggerCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *CapabilitiesRegistry) List(ctx context.Context) ([]capabilities.BaseCapability, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []capabilities.BaseCapability
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]capabilities.BaseCapability, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []capabilities.BaseCapability); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]capabilities.BaseCapability)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCapabilitiesRegistry creates a new instance of CapabilitiesRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCapabilitiesRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *CapabilitiesRegistry {
	mock := &CapabilitiesRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}