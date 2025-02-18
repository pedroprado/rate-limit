package notificationservice

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

func TestProcessNotifications(t *testing.T) {

	t.Run("should delivery notifications to correct channels", func(t *testing.T) {
		statusNotifications, newsNotifications, marketingNotifications, notificationsChan := generateNotificationsSample()

		statusChan := make(chan entity.Notification)
		newsChan := make(chan entity.Notification)
		marketingChan := make(chan entity.Notification)
		notificationChannelsMap, _ := entity.NewNotificationsChannelsMap(statusChan, newsChan, marketingChan)

		go processNotifications(notificationsChan, notificationChannelsMap)

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
	})
}

func generateNotificationsSample() ([]entity.Notification, []entity.Notification, []entity.Notification, chan entity.Notification) {
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

	notificationsChan := make(chan entity.Notification, 9)
	for _, notification := range statusNotifications {
		notificationsChan <- notification
	}
	for _, notification := range newsNotifications {
		notificationsChan <- notification
	}
	for _, notification := range marketingNotifications {
		notificationsChan <- notification
	}
	close(notificationsChan)

	return statusNotifications, newsNotifications, marketingNotifications, notificationsChan
}
