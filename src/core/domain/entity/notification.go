package entity

import "notification-service/src/core/domain/values"

type Notification struct {
	Type    values.NotificationType
	Content string
	Email   string
}
