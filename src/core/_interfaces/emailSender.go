package interfaces

import (
	"context"
	"notification-service/src/core/domain/entity"
)

type EmailSender interface {
	Send(ctx context.Context, notification entity.Notification)
}
