package interfaces

import "context"

type EmailSender interface {
	Send(ctx context.Context, content, email string)
}
