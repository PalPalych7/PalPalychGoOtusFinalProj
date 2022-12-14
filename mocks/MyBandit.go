// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mainstructs "github.com/PalPalych7/OtusProjectWork/internal/mainstructs"
	mock "github.com/stretchr/testify/mock"
)

// MyBandit is an autogenerated mock type for the MyBandit type
type MyBandit struct {
	mock.Mock
}

// GetBannerNum provides a mock function with given fields: arrStruct
func (_m *MyBandit) GetBannerNum(arrStruct []mainstructs.BannerStruct) int {
	ret := _m.Called(arrStruct)

	var r0 int
	if rf, ok := ret.Get(0).(func([]mainstructs.BannerStruct) int); ok {
		r0 = rf(arrStruct)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewMyBandit interface {
	mock.TestingT
	Cleanup(func())
}

// NewMyBandit creates a new instance of MyBandit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMyBandit(t mockConstructorTestingTNewMyBandit) *MyBandit {
	mock := &MyBandit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
