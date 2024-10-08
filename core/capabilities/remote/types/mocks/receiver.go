// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	mock "github.com/stretchr/testify/mock"
)

// Receiver is an autogenerated mock type for the Receiver type
type Receiver struct {
	mock.Mock
}

type Receiver_Expecter struct {
	mock *mock.Mock
}

func (_m *Receiver) EXPECT() *Receiver_Expecter {
	return &Receiver_Expecter{mock: &_m.Mock}
}

// Receive provides a mock function with given fields: ctx, msg
func (_m *Receiver) Receive(ctx context.Context, msg *types.MessageBody) {
	_m.Called(ctx, msg)
}

// Receiver_Receive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Receive'
type Receiver_Receive_Call struct {
	*mock.Call
}

// Receive is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *types.MessageBody
func (_e *Receiver_Expecter) Receive(ctx interface{}, msg interface{}) *Receiver_Receive_Call {
	return &Receiver_Receive_Call{Call: _e.mock.On("Receive", ctx, msg)}
}

func (_c *Receiver_Receive_Call) Run(run func(ctx context.Context, msg *types.MessageBody)) *Receiver_Receive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.MessageBody))
	})
	return _c
}

func (_c *Receiver_Receive_Call) Return() *Receiver_Receive_Call {
	_c.Call.Return()
	return _c
}

func (_c *Receiver_Receive_Call) RunAndReturn(run func(context.Context, *types.MessageBody)) *Receiver_Receive_Call {
	_c.Call.Return(run)
	return _c
}

// NewReceiver creates a new instance of Receiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReceiver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Receiver {
	mock := &Receiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
