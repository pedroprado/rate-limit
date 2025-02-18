package values

import "fmt"

type NotificationType string

const (
	NotificationTypeStatus    NotificationType = "STATUS"
	NotificationTypeNews      NotificationType = "NEWS"
	NotificationTypeMarketing NotificationType = "MARKETING"
)

var validNotificationTypes = map[NotificationType]NotificationType{
	NotificationTypeStatus:    NotificationTypeStatus,
	NotificationTypeNews:      NotificationTypeNews,
	NotificationTypeMarketing: NotificationTypeMarketing,
}

func (notificationType NotificationType) Validate() error {
	_, exists := validNotificationTypes[notificationType]
	if !exists {
		return fmt.Errorf("%s is an invalid notification type", notificationType)
	}

	return nil
}
