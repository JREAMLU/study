// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MessageService is an autogenerated mock type for the MessageService type
type MessageService struct {
	mock.Mock
}

// SendChargeNotification provides a mock function with given fields: _a0
func (_m *MessageService) SendChargeNotification(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
