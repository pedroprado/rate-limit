package notificationsapi

import (
	"time"

	"github.com/google/uuid"

	"notification-service/src/core/domain/entity"
)

type NotificationResponse struct {
	NotificationID uuid.UUID `json:"notification_id"`
	CreateNotificationRequest
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_At"`
}

func NotificationResponseFromDomain(notification entity.Notification) NotificationResponse {
	return NotificationResponse{
		NotificationID: notification.NotificationID,
		CreateNotificationRequest: CreateNotificationRequest{
			Type:    string(notification.Type),
			Content: notification.Content,
			Email:   notification.Email,
		},
		Status:    notification.Status,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}
