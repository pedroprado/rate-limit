package processor

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocks "notification-service/src/core/_mocks"
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

var (
	newChannel    = make(chan entity.Notification, 2)
	createChannel = func() chan entity.Notification {
		return newChannel
	}
)

func TestProcess(t *testing.T) {
	t.Run("should send to recipient channel when it already exists", func(t *testing.T) {
		ctx := context.TODO()

		emailRecipient := uuid.NewString()
		notificationChan := make(chan entity.Notification, 10)
		for i := 0; i < 10; i++ {
			notificationChan <- entity.Notification{
				Type:    values.NotificationTypeStatus,
				Content: uuid.NewString(),
				Email:   emailRecipient,
			}
		}
		close(notificationChan)

		notificationChannelStarter := &mocks.NotificationChannelStarter{}
		existingChannel := make(chan entity.Notification, 2)
		recipientsChannel[emailRecipient] = existingChannel

		notificationChannelStarter.On("StartForRecipient", ctx, emailRecipient, existingChannel).Return().Times(2)

		processor := NewNotificationProcessor(notificationChan, notificationChannelStarter, createChannel)
		time.Sleep(time.Second)

		processor.Process(ctx)

		notificationChannelStarter.AssertNumberOfCalls(t, "StartForRecipient", 0)
	})

	t.Run("should start new recipient channel when it does not exist", func(t *testing.T) {
		ctx := context.TODO()

		emailRecipient := uuid.NewString()
		notificationChan := make(chan entity.Notification, 10)
		for i := 0; i < 10; i++ {
			notificationChan <- entity.Notification{
				Type:    values.NotificationTypeStatus,
				Content: uuid.NewString(),
				Email:   emailRecipient,
			}
		}
		close(notificationChan)

		notificationChannelStarter := &mocks.NotificationChannelStarter{}

		notificationChannelStarter.On("StartForRecipient", ctx, emailRecipient, newChannel).Return().Times(1)

		processor := NewNotificationProcessor(notificationChan, notificationChannelStarter, createChannel)
		time.Sleep(time.Second)

		processor.Process(ctx)

		mock.AssertExpectationsForObjects(t, notificationChannelStarter)
	})
}

func TestGetNotificationChannelForRecipient(t *testing.T) {
	t.Run("should get existing channel for recipient", func(t *testing.T) {
		ctx := context.TODO()
		notificationChannelStarter := &mocks.NotificationChannelStarter{}
		emailRecipient := uuid.NewString()
		existingChannel := make(chan entity.Notification)
		recipientsChannel[emailRecipient] = existingChannel

		expected := existingChannel
		received := getNotificationChannelForRecipient(ctx, notificationChannelStarter, createChannel, emailRecipient)

		assert.Equal(t, expected, received)
	})

	t.Run("should start channel for recipient", func(t *testing.T) {
		ctx := context.TODO()
		notificationChannelStarter := &mocks.NotificationChannelStarter{}
		emailRecipient := uuid.NewString()

		notificationChannelStarter.On("StartForRecipient", ctx, emailRecipient, newChannel).Return()

		expected := newChannel
		received := getNotificationChannelForRecipient(ctx, notificationChannelStarter, createChannel, emailRecipient)
		time.Sleep(time.Second)

		assert.Equal(t, expected, received)
		mock.AssertExpectationsForObjects(t, notificationChannelStarter)
	})
}
