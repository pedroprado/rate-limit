// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// NotificationService is an autogenerated mock type for the NotificationService type
type NotificationService struct {
	mock.Mock
}

// ProcessNotifications provides a mock function with given fields: ctx
func (_m *NotificationService) ProcessNotifications(ctx context.Context) {
	_m.Called(ctx)
}

// NewNotificationService creates a new instance of NotificationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationService {
	mock := &NotificationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
