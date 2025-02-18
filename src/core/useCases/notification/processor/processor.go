package processor

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

var (
	recipientsChannel = map[string]chan entity.Notification{}
)

type notificationProcessor struct {
	notificationChan           chan entity.Notification
	notificationChannelStarter interfaces.NotificationChannelStarter
	createChannel              func() chan entity.Notification
}

func NewNotificationProcessor(
	notificationChan chan entity.Notification,
	notificationChannelStarter interfaces.NotificationChannelStarter,
	createChannel func() chan entity.Notification,
) interfaces.NotificationProcessor {
	return &notificationProcessor{
		notificationChan:           notificationChan,
		notificationChannelStarter: notificationChannelStarter,
		createChannel:              createChannel,
	}
}

func (s *notificationProcessor) Process(ctx context.Context) {
	for notification := range s.notificationChan {
		notificationForRecipientChan := getNotificationChannelForRecipient(ctx, s.notificationChannelStarter, s.createChannel, notification.Email)

		select {
		case notificationForRecipientChan <- notification:
		default:
			logrus.WithContext(ctx).Infof("[NotificationProcessor] discarding notification for recipient %s. Exceeded rate limit", notification.Email)
		}
	}
}

func getNotificationChannelForRecipient(
	ctx context.Context,
	notificationChannelStarter interfaces.NotificationChannelStarter,
	createChannel func() chan entity.Notification,
	emailRecipient string) chan entity.Notification {
	notificationChannelForRecipient, exists := recipientsChannel[emailRecipient]
	if exists {
		return notificationChannelForRecipient
	}

	newNotificationChannelForRecipient := createChannel()
	recipientsChannel[emailRecipient] = newNotificationChannelForRecipient

	go notificationChannelStarter.StartForRecipient(ctx, emailRecipient, newNotificationChannelForRecipient)

	return newNotificationChannelForRecipient
}
