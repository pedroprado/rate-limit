package processor

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type notificationSender struct {
	notificationFrequency int
	notificationType      string
	emailSender           interfaces.EmailSender
}

func NewNotificationChannelStarter(
	notificationFrequency int,
	notificationType string,
	emailSender interfaces.EmailSender,
) interfaces.NotificationChannelStarter {
	return &notificationSender{
		notificationFrequency: notificationFrequency,
		notificationType:      notificationType,
		emailSender:           emailSender,
	}
}

func (n *notificationSender) StartForRecipient(
	ctx context.Context,
	emailRecipient string,
	notificationChannelForRecipient chan entity.Notification,
) {
	logrus.WithContext(ctx).Infof("[NotificationProcessor] starting notification %s channel for recipient: %s", n.notificationType, emailRecipient)
	ticker := time.Tick(time.Second * time.Duration(n.notificationFrequency))

	for notificationForRecipient := range notificationChannelForRecipient {
		<-ticker
		n.emailSender.Send(ctx, notificationForRecipient)
	}
}
