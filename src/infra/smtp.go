package infra

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
)

type smtpService struct{}

func NewSmtpService() interfaces.SmtpService {
	return &smtpService{}
}

func (s *smtpService) SendEmail(ctx context.Context, content string, to string) error {
	// TODO: replace with actual smtp service implementation for sending email
	logrus.WithFields(
		map[string]interface{}{
			"to":      to,
			"content": content,
		},
	).Info("Sending email")

	return nil
}
