package interfaces

import (
	"context"

	"notification-service/src/core/domain/entity"
)

type NotificationChannelStarter interface {
	StartForRecipient(
		ctx context.Context,
		emailRecipient string,
		notificationForRecipientChan chan entity.Notification,
	)
}
