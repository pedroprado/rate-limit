// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "notification-service/src/core/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// NotificationChannelStarter is an autogenerated mock type for the NotificationChannelStarter type
type NotificationChannelStarter struct {
	mock.Mock
}

// StartNotifyingRecipient provides a mock function with given fields: ctx, emailRecipient, notificationForRecipientChan
func (_m *NotificationChannelStarter) StartNotifyingRecipient(ctx context.Context, emailRecipient string, notificationForRecipientChan chan entity.Notification) {
	_m.Called(ctx, emailRecipient, notificationForRecipientChan)
}

// NewNotificationChannelStarter creates a new instance of NotificationChannelStarter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationChannelStarter(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationChannelStarter {
	mock := &NotificationChannelStarter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
