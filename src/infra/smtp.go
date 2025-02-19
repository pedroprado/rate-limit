package infra

import (
	"context"

	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
)

type googleSmtpService struct{}

func NewGoogleSmtpService() interfaces.SmtpService {
	return &googleSmtpService{}
}

func (s *googleSmtpService) SendEmail(ctx context.Context, content string, to string) error {
	// TODO: replace with actual smtp service implementation for sending email
	logrus.WithFields(
		map[string]interface{}{
			"to":      to,
			"content": content,
		},
	).Info("Sending email")

	return nil
}
