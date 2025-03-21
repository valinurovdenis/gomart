// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	orderstorage "github.com/valinurovdenis/gomart/internal/app/orderstorage"

	userstorage "github.com/valinurovdenis/gomart/internal/app/userstorage"

	withdrawstorage "github.com/valinurovdenis/gomart/internal/app/withdrawstorage"
)

// ServiceStorage is an autogenerated mock type for the ServiceStorage type
type ServiceStorage struct {
	mock.Mock
}

// AddUserOrder provides a mock function with given fields: _a0, order
func (_m *ServiceStorage) AddUserOrder(_a0 context.Context, order orderstorage.UserOrder) error {
	ret := _m.Called(_a0, order)

	if len(ret) == 0 {
		panic("no return value specified for AddUserOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, orderstorage.UserOrder) error); ok {
		r0 = rf(_a0, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUserWithdraw provides a mock function with given fields: ctx, order
func (_m *ServiceStorage) AddUserWithdraw(ctx context.Context, order withdrawstorage.UserWithdraw) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for AddUserWithdraw")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, withdrawstorage.UserWithdraw) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalance provides a mock function with given fields: _a0, login
func (_m *ServiceStorage) GetBalance(_a0 context.Context, login string) (userstorage.UserBalance, error) {
	ret := _m.Called(_a0, login)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 userstorage.UserBalance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (userstorage.UserBalance, error)); ok {
		return rf(_a0, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) userstorage.UserBalance); ok {
		r0 = rf(_a0, login)
	} else {
		r0 = ret.Get(0).(userstorage.UserBalance)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserOrders provides a mock function with given fields: _a0, login
func (_m *ServiceStorage) GetUserOrders(_a0 context.Context, login string) ([]orderstorage.UserOrder, error) {
	ret := _m.Called(_a0, login)

	if len(ret) == 0 {
		panic("no return value specified for GetUserOrders")
	}

	var r0 []orderstorage.UserOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]orderstorage.UserOrder, error)); ok {
		return rf(_a0, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []orderstorage.UserOrder); ok {
		r0 = rf(_a0, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]orderstorage.UserOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserWithdrawals provides a mock function with given fields: ctx, login
func (_m *ServiceStorage) GetUserWithdrawals(ctx context.Context, login string) ([]withdrawstorage.UserWithdraw, error) {
	ret := _m.Called(ctx, login)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWithdrawals")
	}

	var r0 []withdrawstorage.UserWithdraw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]withdrawstorage.UserWithdraw, error)); ok {
		return rf(ctx, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []withdrawstorage.UserWithdraw); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]withdrawstorage.UserWithdraw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceStorage creates a new instance of ServiceStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceStorage {
	mock := &ServiceStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
