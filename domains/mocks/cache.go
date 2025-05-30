// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
// Copyright (c) Abstract Machines

// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"

	"github.com/absmach/supermq/domains"
	mock "github.com/stretchr/testify/mock"
)

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

type Cache_Expecter struct {
	mock *mock.Mock
}

func (_m *Cache) EXPECT() *Cache_Expecter {
	return &Cache_Expecter{mock: &_m.Mock}
}

// Remove provides a mock function for the type Cache
func (_mock *Cache) Remove(ctx context.Context, domainID string) error {
	ret := _mock.Called(ctx, domainID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = returnFunc(ctx, domainID)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Cache_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type Cache_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx
//   - domainID
func (_e *Cache_Expecter) Remove(ctx interface{}, domainID interface{}) *Cache_Remove_Call {
	return &Cache_Remove_Call{Call: _e.mock.On("Remove", ctx, domainID)}
}

func (_c *Cache_Remove_Call) Run(run func(ctx context.Context, domainID string)) *Cache_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Cache_Remove_Call) Return(err error) *Cache_Remove_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Cache_Remove_Call) RunAndReturn(run func(ctx context.Context, domainID string) error) *Cache_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function for the type Cache
func (_mock *Cache) Save(ctx context.Context, domainID string, status domains.Status) error {
	ret := _mock.Called(ctx, domainID, status)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, domains.Status) error); ok {
		r0 = returnFunc(ctx, domainID, status)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Cache_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Cache_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx
//   - domainID
//   - status
func (_e *Cache_Expecter) Save(ctx interface{}, domainID interface{}, status interface{}) *Cache_Save_Call {
	return &Cache_Save_Call{Call: _e.mock.On("Save", ctx, domainID, status)}
}

func (_c *Cache_Save_Call) Run(run func(ctx context.Context, domainID string, status domains.Status)) *Cache_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domains.Status))
	})
	return _c
}

func (_c *Cache_Save_Call) Return(err error) *Cache_Save_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Cache_Save_Call) RunAndReturn(run func(ctx context.Context, domainID string, status domains.Status) error) *Cache_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function for the type Cache
func (_mock *Cache) Status(ctx context.Context, domainID string) (domains.Status, error) {
	ret := _mock.Called(ctx, domainID)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 domains.Status
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (domains.Status, error)); ok {
		return returnFunc(ctx, domainID)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) domains.Status); ok {
		r0 = returnFunc(ctx, domainID)
	} else {
		r0 = ret.Get(0).(domains.Status)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, domainID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Cache_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type Cache_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - ctx
//   - domainID
func (_e *Cache_Expecter) Status(ctx interface{}, domainID interface{}) *Cache_Status_Call {
	return &Cache_Status_Call{Call: _e.mock.On("Status", ctx, domainID)}
}

func (_c *Cache_Status_Call) Run(run func(ctx context.Context, domainID string)) *Cache_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Cache_Status_Call) Return(status domains.Status, err error) *Cache_Status_Call {
	_c.Call.Return(status, err)
	return _c
}

func (_c *Cache_Status_Call) RunAndReturn(run func(ctx context.Context, domainID string) (domains.Status, error)) *Cache_Status_Call {
	_c.Call.Return(run)
	return _c
}
