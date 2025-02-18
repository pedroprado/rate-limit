package interfaces

import (
	"context"

	"notification-service/src/core/domain/entity"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification entity.Notification) (*entity.Notification, error)
}
