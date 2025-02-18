package emailsender

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
)

type emailSender struct {
}

func NewEmailSender() interfaces.EmailSender {
	return &emailSender{}
}

func (e *emailSender) Send(ctx context.Context, content string, email string) {
	logrus.WithFields(
		map[string]interface{}{
			"to":      email,
			"content": content,
		},
	).Info("Sending email")
}
