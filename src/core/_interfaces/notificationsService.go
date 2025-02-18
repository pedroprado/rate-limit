package interfaces

import (
	"context"
	"notification-service/src/core/domain/entity"
)

type NotificationsService interface {
	CreateNotification(ctx context.Context, notification entity.Notification) (*entity.Notification, error)
}
