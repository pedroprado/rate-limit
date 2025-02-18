package notificationservice

import (
	"context"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type notificationService struct {
	notificationsChan    chan entity.Notification
	notificationChannels entity.NotificationsChannels
}

func NewNotificationService(
	notificationsChan chan entity.Notification,
	notificationChannels entity.NotificationsChannels,
) interfaces.NotificationService {
	return &notificationService{
		notificationsChan:    notificationsChan,
		notificationChannels: notificationChannels,
	}
}

func (n *notificationService) ProcessNotifications(ctx context.Context) {
	for {
		processNotifications(n.notificationsChan, n.notificationChannels)
	}
}

func processNotifications(
	notificationsChan chan entity.Notification,
	notificationChannels entity.NotificationsChannels,
) {
	notification := <-notificationsChan

	destinationChannel := notificationChannels[notification.Type]

	destinationChannel <- notification
}
