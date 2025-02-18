package processor

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mocks "notification-service/src/core/_mocks"
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

func TestSendForRecipient(t *testing.T) {
	ctx := context.TODO()
	notificationFrequency := 1
	emailSender := &mocks.EmailSender{}

	starter := NewNotificationChannelStarter(notificationFrequency, emailSender)

	notificationForRecipientChan := make(chan entity.Notification, 3)
	notifications := []entity.Notification{
		{
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
		{
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		}, {
			Type:    values.NotificationTypeStatus,
			Content: uuid.NewString(),
			Email:   uuid.NewString(),
		},
	}
	for i := 0; i < 3; i++ {
		notificationForRecipientChan <- notifications[i]
	}
	close(notificationForRecipientChan)

	emailSender.On("Send", ctx, notifications[0].Content, notifications[0].Email).Return()
	emailSender.On("Send", ctx, notifications[1].Content, notifications[1].Email).Return()
	emailSender.On("Send", ctx, notifications[2].Content, notifications[2].Email).Return()

	starter.StartForRecipient(ctx, "email", notificationForRecipientChan)

	mock.AssertExpectationsForObjects(t, emailSender)
}
