package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
)

const (
	collection = "notifications"
)

type notificationRepository struct {
	db *firestore.Client
}

func NewNotificationRepository(db *firestore.Client) interfaces.NotificationRepository {
	return &notificationRepository{db: db}
}

func (n *notificationRepository) Save(ctx context.Context, notification entity.Notification) (*entity.Notification, error) {
	timeNow := time.Now()
	if notification.NotificationID == uuid.Nil {
		notification.NotificationID = uuid.New()
		notification.CreatedAt = timeNow
		notification.UpdatedAt = timeNow
	} else {
		notification.UpdatedAt = timeNow
	}

	docRef := n.db.Collection(collection).Doc(notification.NotificationID.String())
	record := NewNotificationRecordFromDomain(notification)
	_, err := docRef.Set(ctx, record)
	if err != nil {
		return nil, err
	}

	return record.ToDomain(), nil
}
