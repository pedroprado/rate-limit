package interfaces

import "context"

type SmtpService interface {
	SendEmail(ctx context.Context, content, to string) error
}
