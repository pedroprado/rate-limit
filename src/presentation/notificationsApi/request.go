package notificationsapi

import (
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

type CreateNotificationRequest struct {
	Type    string `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

func (request CreateNotificationRequest) ToDomain() entity.Notification {
	return entity.Notification{
		Type:    values.NotificationType(request.Type),
		Content: request.Content,
		Email:   request.Email,
	}
}
