package processor

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type notificationProcessor struct {
	notificationRepo           interfaces.NotificationRepository
	notificationChan           chan entity.Notification
	notificationChannelStarter interfaces.NotificationChannelStarter
	createChannelForRecipient  func() chan entity.Notification
	recipientsChannels         *entity.SafeRecipientsChannel
}

func NewNotificationProcessor(
	notificationRepo interfaces.NotificationRepository,
	notificationChan chan entity.Notification,
	notificationChannelStarter interfaces.NotificationChannelStarter,
	createChannelForRecipient func() chan entity.Notification,
	recipientsChannels *entity.SafeRecipientsChannel,
) interfaces.NotificationProcessor {
	return &notificationProcessor{
		notificationRepo:           notificationRepo,
		notificationChan:           notificationChan,
		notificationChannelStarter: notificationChannelStarter,
		createChannelForRecipient:  createChannelForRecipient,
		recipientsChannels:         recipientsChannels,
	}
}

func (s *notificationProcessor) Process(ctx context.Context) {
	for notification := range s.notificationChan {
		notificationForRecipientChan := getNotificationChannelForRecipient(ctx, s.notificationChannelStarter, s.createChannelForRecipient, s.recipientsChannels, notification.Email)

		select {
		case notificationForRecipientChan <- notification:
		default:
			logrus.WithContext(ctx).Infof("[NotificationProcessor] discarding notification for recipient %s. Exceeded rate limit", notification.Email)
			notification.Status = "REJECTED"
			_, err := s.notificationRepo.Save(ctx, notification)
			if err != nil {
				logrus.Errorf("could not mark notification as REJECTED: %s", notification.NotificationID)
			}
		}
	}
}

func getNotificationChannelForRecipient(
	ctx context.Context,
	notificationChannelStarter interfaces.NotificationChannelStarter,
	createChannelForRecipient func() chan entity.Notification,
	recipientsChannels *entity.SafeRecipientsChannel,
	emailRecipient string,
) chan entity.Notification {

	recipientsChannels.Lock()
	defer recipientsChannels.Unlock()
	notificationChannelForRecipient, exists := recipientsChannels.Channels[emailRecipient]
	if exists {
		return notificationChannelForRecipient
	}

	newNotificationChannelForRecipient := createChannelForRecipient()
	recipientsChannels.Channels[emailRecipient] = newNotificationChannelForRecipient

	go notificationChannelStarter.StartNotifyingRecipient(ctx, emailRecipient, newNotificationChannelForRecipient)

	return newNotificationChannelForRecipient
}
