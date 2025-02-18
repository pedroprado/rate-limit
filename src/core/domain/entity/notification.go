package entity

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"

	"notification-service/src/core/domain/values"
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Notification struct {
	NotificationID uuid.UUID
	Type           values.NotificationType
	Content        string
	Email          string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (notification Notification) Validate() error {
	if err := notification.Type.Validate(); err != nil {
		return fmt.Errorf("notification_type: %s", err)
	}

	if match := emailRegex.MatchString(notification.Email); !match {
		return fmt.Errorf("notifaction_email: is not a valid email format")
	}

	return nil
}
