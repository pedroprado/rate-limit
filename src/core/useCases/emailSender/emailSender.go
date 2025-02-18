package emailsender

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type emailSender struct {
	sfmtService      interfaces.SmtpService
	notificationRepo interfaces.NotificationRepository
}

func NewEmailSender(
	sfmtService interfaces.SmtpService,
	notificationRepo interfaces.NotificationRepository) interfaces.EmailSender {
	return &emailSender{
		sfmtService:      sfmtService,
		notificationRepo: notificationRepo,
	}
}

func (e *emailSender) Send(ctx context.Context, notification entity.Notification) {
	if err := e.sfmtService.SendEmail(ctx, notification.Content, notification.Email); err != nil {
		notification.Status = "ERROR"
		logrus.WithField("notification_id", notification.NotificationID.String()).Errorf("[EmailSender] error sending email: %s", err.Error())
	} else {
		notification.Status = "SENT"
	}

	_, err := e.notificationRepo.Save(ctx, notification)
	if err != nil {
		logrus.Errorf("[EmailSender] could not mark notification as %s: %s", notification.Status, notification.NotificationID)
	}
}
