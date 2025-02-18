package emailsender

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	mocks "notification-service/src/core/_mocks"
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
)

func TestSend(t *testing.T) {
	ctx := context.TODO()

	t.Run("should send email successfully", func(t *testing.T) {
		stmpService := &mocks.SmtpService{}
		notificationRepo := &mocks.NotificationRepository{}
		sender := NewEmailSender(stmpService, notificationRepo)

		notification := entity.Notification{
			NotificationID: uuid.New(),
			Type:           values.NotificationTypeMarketing,
			Content:        uuid.NewString(),
			Email:          uuid.NewString() + "@mail.com",
			Status:         "PENDING",
		}

		notificationToSave := notification
		notificationToSave.Status = "SENT"

		stmpService.On("SendEmail", ctx, notification.Content, notification.Email).Return(nil)
		notificationRepo.On("Save", ctx, notificationToSave).Return(&notificationToSave, nil)

		sender.Send(ctx, notification)

		mock.AssertExpectationsForObjects(t, stmpService, notificationRepo)
	})

	t.Run("should not email successfully", func(t *testing.T) {
		stmpService := &mocks.SmtpService{}
		notificationRepo := &mocks.NotificationRepository{}
		sender := NewEmailSender(stmpService, notificationRepo)

		notification := entity.Notification{
			NotificationID: uuid.New(),
			Type:           values.NotificationTypeMarketing,
			Content:        uuid.NewString(),
			Email:          uuid.NewString() + "@mail.com",
			Status:         "PENDING",
		}

		notificationToSave := notification
		notificationToSave.Status = "ERROR"

		stmpService.On("SendEmail", ctx, notification.Content, notification.Email).Return(errors.New("error sending email"))
		notificationRepo.On("Save", ctx, notificationToSave).Return(&notificationToSave, nil)

		sender.Send(ctx, notification)

		mock.AssertExpectationsForObjects(t, stmpService, notificationRepo)
	})
}
