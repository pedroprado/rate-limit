package entity

import (
	"errors"

	"notification-service/src/core/domain/values"
)

type NotificationsChannels map[values.NotificationType]chan Notification

func NewNotificationsChannels(
	statusChan chan Notification,
	newsChan chan Notification,
	marketingChan chan Notification,
) (NotificationsChannels, error) {
	if statusChan == nil || newsChan == nil || marketingChan == nil {
		return nil, errors.New("should have all types of notification channels")
	}

	return NotificationsChannels{
		values.NotificationTypeStatus:    statusChan,
		values.NotificationTypeNews:      newsChan,
		values.NotificationTypeMarketing: marketingChan,
	}, nil
}
