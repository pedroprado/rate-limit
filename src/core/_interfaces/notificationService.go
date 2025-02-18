package interfaces

import (
	"context"
)

type NotificationService interface {
	ProcessNotifications(ctx context.Context)
}
