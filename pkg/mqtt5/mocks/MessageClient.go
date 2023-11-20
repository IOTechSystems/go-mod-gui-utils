// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/IOTechSystems/go-mod-edge-utils/pkg/bootstrap/interfaces"
	errors "github.com/IOTechSystems/go-mod-edge-utils/pkg/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/IOTechSystems/go-mod-edge-utils/pkg/mqtt5/models"
)

// MessageClient is an autogenerated mock type for the MessageClient type
type MessageClient struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *MessageClient) Connect() errors.Error {
	ret := _m.Called()

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func() errors.Error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *MessageClient) Disconnect() errors.Error {
	ret := _m.Called()

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func() errors.Error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// Publish provides a mock function with given fields: topic, message
func (_m *MessageClient) Publish(topic string, message models.MessageEnvelope) errors.Error {
	ret := _m.Called(topic, message)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(string, models.MessageEnvelope) errors.Error); ok {
		r0 = rf(topic, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// SetAuthData provides a mock function with given fields: secretProvider
func (_m *MessageClient) SetAuthData(secretProvider interfaces.SecretProvider) errors.Error {
	ret := _m.Called(secretProvider)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(interfaces.SecretProvider) errors.Error); ok {
		r0 = rf(secretProvider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// Subscribe provides a mock function with given fields: topics, handlerType
func (_m *MessageClient) Subscribe(topics []string, handlerType interface{}) errors.Error {
	ret := _m.Called(topics, handlerType)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func([]string, interface{}) errors.Error); ok {
		r0 = rf(topics, handlerType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: topics
func (_m *MessageClient) Unsubscribe(topics []string) errors.Error {
	ret := _m.Called(topics)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func([]string) errors.Error); ok {
		r0 = rf(topics)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// NewMessageClient creates a new instance of MessageClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageClient {
	mock := &MessageClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
