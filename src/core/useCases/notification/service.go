package notificationservice

import (
	"context"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type notificationsService struct {
	notificationRepo        interfaces.NotificationRepository
	notificationsChan       chan entity.Notification
	notificationChannelsMap entity.NotificationsChannelsMap
}

func NewNotificationsService(
	notificationRepo interfaces.NotificationRepository,
	notificationsChan chan entity.Notification,
	notificationChannelsMap entity.NotificationsChannelsMap,
) interfaces.NotificationsService {
	return &notificationsService{
		notificationRepo:        notificationRepo,
		notificationsChan:       notificationsChan,
		notificationChannelsMap: notificationChannelsMap,
	}
}

func (n *notificationsService) CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error) {
	notification.Status = "PENDING"
	saved, err := n.notificationRepo.Save(ctx, notification)
	if err != nil {
		return nil, err
	}

	destinationChannel := n.notificationChannelsMap[notification.Type]

	destinationChannel <- *saved

	return saved, nil
}

func processNotifications(
	notificationsChan chan entity.Notification,
	notificationChannels entity.NotificationsChannelsMap,
) {
	for {
		processNotification(notificationsChan, notificationChannels)
	}
}

func processNotification(
	notificationsChan chan entity.Notification,
	notificationChannels entity.NotificationsChannelsMap,
) {
	notification := <-notificationsChan

	destinationChannel := notificationChannels[notification.Type]

	destinationChannel <- notification
}
