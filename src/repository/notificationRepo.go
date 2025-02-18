package repository

import (
	"context"
	"time"

	"github.com/google/uuid"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

type notificationRepository struct{}

func NewNotificationRepository() interfaces.NotificationRepository {
	return &notificationRepository{}
}

func (n *notificationRepository) Save(ctx context.Context, notification entity.Notification) (*entity.Notification, error) {
	timeNow := time.Now()
	if notification.NotificationID == uuid.Nil {
		notification.NotificationID = uuid.New()
		notification.CreatedAt = timeNow
		notification.UpdatedAt = timeNow
	} else {
		notification.UpdatedAt = timeNow
	}

	return &notification, nil
}
