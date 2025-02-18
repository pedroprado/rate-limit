package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"notification-service/src/core/domain/values"
)

func TestValidateNotification(t *testing.T) {
	t.Run("should be invalid when invalid type", func(t *testing.T) {
		notification := Notification{
			Type: values.NotificationType("INVALID"),
		}

		err := notification.Validate()

		assert.ErrorContains(t, err, "INVALID")
	})

	t.Run("should be invalid when invalid email", func(t *testing.T) {
		notification := Notification{
			Type:  values.NotificationTypeStatus,
			Email: "invalid email",
		}

		err := notification.Validate()

		assert.ErrorContains(t, err, "is not a valid email format")
	})

	t.Run("should be valid", func(t *testing.T) {
		notification := Notification{
			Type:  values.NotificationTypeStatus,
			Email: "example@mail.com",
		}

		err := notification.Validate()

		assert.NoError(t, err)
	})
}
