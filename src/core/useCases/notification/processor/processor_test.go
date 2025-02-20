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
	newChannelForRecipient    = make(chan entity.Notification, 1)
	createChannelForRecipient = func() chan entity.Notification {
		return newChannelForRecipient
	}
)

func TestProcess(t *testing.T) {
	ctx := context.TODO()

	t.Run("should send to existing recipient channel, rejecting when exceed channel capacity", func(t *testing.T) {
		notificationRepo := &mocks.NotificationRepository{}
		notificationChannelStarter := &mocks.NotificationChannelStarter{}

		emailRecipient := uuid.NewString() + "@mail.com"
		notifications := []entity.Notification{
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
		}
		notificationChan := make(chan entity.Notification, 4)

		for _, notification := range notifications {
			notificationChan <- notification
		}
		close(notificationChan)
		recipientsChannels := entity.NewRecipientsChannel()
		recipientsChannels.Channels[emailRecipient] = make(chan entity.Notification, 1)

		processor := NewNotificationProcessor(notificationRepo, notificationChan, notificationChannelStarter, createChannelForRecipient, recipientsChannels)

		for i := 1; i < 4; i++ {
			notificationRejected := notifications[i]
			notificationRejected.Status = "REJECTED"
			notificationRepo.On("Save", ctx, notificationRejected).Return(nil, nil)
		}

		go processor.Process(ctx)
		time.Sleep(time.Second * 1)

		notificationChannelStarter.AssertNumberOfCalls(t, "StartNotifyingRecipient", 0)
		mock.AssertExpectationsForObjects(t, notificationRepo)
	})

	t.Run("should send to new recipient channel, rejecting when execeed channel capacity", func(t *testing.T) {
		notificationRepo := &mocks.NotificationRepository{}
		notificationChannelStarter := &mocks.NotificationChannelStarter{}

		emailRecipient := uuid.NewString() + "@mail.com"
		notifications := []entity.Notification{
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
			{
				NotificationID: uuid.New(),
				Type:           values.NotificationTypeStatus,
				Content:        uuid.NewString(),
				Email:          emailRecipient,
			},
		}
		notificationChan := make(chan entity.Notification, 4)

		for _, notification := range notifications {
			notificationChan <- notification
		}
		close(notificationChan)
		recipientsChannels := entity.NewRecipientsChannel()
		processor := NewNotificationProcessor(notificationRepo, notificationChan, notificationChannelStarter, createChannelForRecipient, recipientsChannels)

		for i := 1; i < 4; i++ {
			notificationRejected := notifications[i]
			notificationRejected.Status = "REJECTED"
			notificationRepo.On("Save", ctx, notificationRejected).Return(nil, nil).Times(1)
		}
		notificationChannelStarter.On("StartNotifyingRecipient", ctx, emailRecipient, newChannelForRecipient).Return().Times(1)

		go processor.Process(ctx)
		time.Sleep(time.Second * 1)

		mock.AssertExpectationsForObjects(t, notificationChannelStarter, notificationRepo)
	})
}

func TestGetNotificationChannelForRecipient(t *testing.T) {
	ctx := context.TODO()
	recipientsChannels := entity.NewRecipientsChannel()

	t.Run("should get existing channel for recipient", func(t *testing.T) {
		notificationChannelStarter := &mocks.NotificationChannelStarter{}
		emailRecipient := uuid.NewString()
		existingChannel := make(chan entity.Notification)
		recipientsChannels.Channels[emailRecipient] = existingChannel

		expected := existingChannel
		received := getNotificationChannelForRecipient(ctx, notificationChannelStarter, createChannelForRecipient, recipientsChannels, emailRecipient)

		assert.Equal(t, expected, received)
	})

	t.Run("should start channel for recipient", func(t *testing.T) {
		ctx := context.TODO()
		notificationChannelStarter := &mocks.NotificationChannelStarter{}
		emailRecipient := uuid.NewString()

		notificationChannelStarter.On("StartNotifyingRecipient", ctx, emailRecipient, newChannelForRecipient).Return()

		expected := newChannelForRecipient
		received := getNotificationChannelForRecipient(ctx, notificationChannelStarter, createChannelForRecipient, recipientsChannels, emailRecipient)
		time.Sleep(time.Second)

		assert.Equal(t, expected, received)
		mock.AssertExpectationsForObjects(t, notificationChannelStarter)
	})
}
