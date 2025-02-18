// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "notification-service/src/core/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// EmailSender is an autogenerated mock type for the EmailSender type
type EmailSender struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, notification
func (_m *EmailSender) Send(ctx context.Context, notification entity.Notification) {
	_m.Called(ctx, notification)
}

// NewEmailSender creates a new instance of EmailSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailSender(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailSender {
	mock := &EmailSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
