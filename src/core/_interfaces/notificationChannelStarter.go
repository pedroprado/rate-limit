package interfaces

import (
	"context"

	"notification-service/src/core/domain/entity"
)

type NotificationChannelStarter interface {
	StartNotifyingRecipient(
		ctx context.Context,
		emailRecipient string,
		notificationForRecipientChan chan entity.Notification,
	)
}
