package repository

import (
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
	"time"

	"github.com/google/uuid"
)

type NotificationRecord struct {
	NotificationID uuid.UUID `firestore:"notification_id"`
	Type           string    `firestore:"type"`
	Content        string    `firestore:"content"`
	Email          string    `firestore:"email"`
	Status         string    `firestore:"status"`
	CreatedAt      time.Time `firestore:"created_at"`
	UpdatedAt      time.Time `firestore:"updated_at"`
}

func NewNotificationRecordFromDomain(notification entity.Notification) NotificationRecord {
	return NotificationRecord{
		NotificationID: notification.NotificationID,
		Type:           string(notification.Type),
		Content:        notification.Content,
		Email:          notification.Email,
		Status:         notification.Status,
		CreatedAt:      notification.CreatedAt,
		UpdatedAt:      notification.UpdatedAt,
	}
}

func (record NotificationRecord) ToDomain() *entity.Notification {
	return &entity.Notification{
		NotificationID: record.NotificationID,
		Type:           values.NotificationType(record.Type),
		Content:        record.Content,
		Email:          record.Email,
		Status:         record.Status,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
}
