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
	emailSender           interfaces.EmailSender
}

func NewNotificationChannelStarter(
	notificationFrequency int,
	emailSender interfaces.EmailSender,
) interfaces.NotificationChannelStarter {
	return &notificationSender{
		notificationFrequency: notificationFrequency,
		emailSender:           emailSender,
	}
}

func (n *notificationSender) StartForRecipient(
	ctx context.Context,
	emailRecipient string,
	notificationForRecipientChan chan entity.Notification,
) {
	logrus.WithContext(ctx).Infof("[NotificationProcessor] starting notification channel for recipient: %s", emailRecipient)
	ticker := time.Tick(time.Second * time.Duration(n.notificationFrequency))

	for notificationForRecipient := range notificationForRecipientChan {
		<-ticker
		n.emailSender.Send(ctx, notificationForRecipient.Content, notificationForRecipient.Email)
	}
}
