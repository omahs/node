// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Iterator is an autogenerated mock type for the Iterator type
type Iterator struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *Iterator) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Key provides a mock function with given fields:
func (_m *Iterator) Key() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Iterator) Next() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Release provides a mock function with given fields:
func (_m *Iterator) Release() {
	_m.Called()
}

// Value provides a mock function with given fields:
func (_m *Iterator) Value() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

type mockConstructorTestingTNewIterator interface {
	mock.TestingT
	Cleanup(func())
}

// NewIterator creates a new instance of Iterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIterator(t mockConstructorTestingTNewIterator) *Iterator {
	mock := &Iterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
