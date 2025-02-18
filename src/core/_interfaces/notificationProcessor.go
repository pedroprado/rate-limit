package interfaces

import "context"

type NotificationProcessor interface {
	Process(ctx context.Context)
}
