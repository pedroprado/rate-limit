package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"notification-service/src/core/domain/values"
)

func TestNewNotificationChannels(t *testing.T) {

	t.Run("should create notification channels", func(t *testing.T) {
		statusChan := make(chan Notification)
		newsChan := make(chan Notification)
		marketingChan := make(chan Notification)

		expected := NotificationsChannels{
			values.NotificationTypeStatus:    statusChan,
			values.NotificationTypeNews:      newsChan,
			values.NotificationTypeMarketing: marketingChan,
		}

		received, err := NewNotificationsChannels(statusChan, newsChan, marketingChan)

		assert.Nil(t, err)
		assert.Equal(t, expected, received)
	})

	t.Run("should not create when missing status channel", func(t *testing.T) {
		statusChan := make(chan Notification)
		marketingChan := make(chan Notification)

		received, err := NewNotificationsChannels(statusChan, nil, marketingChan)

		assert.ErrorContains(t, err, "should have all types of notification channels")
		assert.Nil(t, received)
	})

	t.Run("should not create when missing news channel", func(t *testing.T) {
		statusChan := make(chan Notification)
		newsChan := make(chan Notification)

		received, err := NewNotificationsChannels(statusChan, newsChan, nil)

		assert.ErrorContains(t, err, "should have all types of notification channels")
		assert.Nil(t, received)
	})

	t.Run("should not create when missing marketing channel", func(t *testing.T) {
		newsChan := make(chan Notification)
		marketingChan := make(chan Notification)

		received, err := NewNotificationsChannels(nil, newsChan, marketingChan)

		assert.ErrorContains(t, err, "should have all types of notification channels")
		assert.Nil(t, received)
	})
}
