package entity

import (
	"errors"

	"notification-service/src/core/domain/values"
)

type NotificationsChannelsMap map[values.NotificationType]chan Notification

func NewNotificationsChannelsMap(
	statusChan chan Notification,
	newsChan chan Notification,
	marketingChan chan Notification,
) (NotificationsChannelsMap, error) {
	if statusChan == nil || newsChan == nil || marketingChan == nil {
		return nil, errors.New("should have all types of notification channels")
	}

	return NotificationsChannelsMap{
		values.NotificationTypeStatus:    statusChan,
		values.NotificationTypeNews:      newsChan,
		values.NotificationTypeMarketing: marketingChan,
	}, nil
}
