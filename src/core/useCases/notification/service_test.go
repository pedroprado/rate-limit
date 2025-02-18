package notificationservice

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

func TestProcessNotifications(t *testing.T) {
	ctx := context.TODO()

	statusNotifications, newsNotifications, marketingNotifications := generateNotificationsSample()

	notificationChan := make(chan entity.Notification, 9)
	for _, notification := range statusNotifications {
		notificationChan <- notification
	}
	for _, notification := range newsNotifications {
		notificationChan <- notification
	}
	for _, notification := range marketingNotifications {
		notificationChan <- notification
	}
	close(notificationChan)

	statusChan := make(chan entity.Notification)
	newsChan := make(chan entity.Notification)
	marketingChan := make(chan entity.Notification)
	notificationChannels, _ := entity.NewNotificationsChannels(statusChan, newsChan, marketingChan)

	service := NewNotificationService(notificationChan, notificationChannels)
	go service.ProcessNotifications(ctx)

	statusNotificationIndex := 0
	for notification := range statusChan {
		assert.Equal(t, notification, statusNotifications[statusNotificationIndex])
		statusNotificationIndex++
		if statusNotificationIndex == 3 {
			break
		}
	}

	newsNotificationIndex := 0
	for notification := range newsChan {
		assert.Equal(t, notification, newsNotifications[newsNotificationIndex])
		newsNotificationIndex++
		if newsNotificationIndex == 3 {
			break
		}
	}

	makertingNotificationIndex := 0
	for notification := range marketingChan {
		assert.Equal(t, notification, marketingNotifications[makertingNotificationIndex])
		makertingNotificationIndex++
		if makertingNotificationIndex == 3 {
			break
		}
	}
}

func generateNotificationsSample() ([]entity.Notification, []entity.Notification, []entity.Notification) {
	statusNotifications := []entity.Notification{
		{
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
	}
	newsNotifications := []entity.Notification{
		{
			Type:    values.NotificationTypeNews,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeNews,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeNews,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
	}
	marketingNotifications := []entity.Notification{
		{
			Type:    values.NotificationTypeMarketing,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeMarketing,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeMarketing,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
	}

	return statusNotifications, newsNotifications, marketingNotifications
}
