package repository

import (
	"notification-service/src/core/domain/values"
	"time"

	"github.com/google/uuid"
)

type NotificationRecord struct {
	NotificationID uuid.UUID
	Type           values.NotificationType
	Content        string
	Email          string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
