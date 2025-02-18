package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
	emailsender "notification-service/src/core/useCases/emailSender"
	notificationservice "notification-service/src/core/useCases/notification"
	"notification-service/src/core/useCases/notification/processor"
)

var (
	notificationsChan = make(chan entity.Notification)
	statusChan        = make(chan entity.Notification)
	newsChan          = make(chan entity.Notification)
	marketingChan     = make(chan entity.Notification)

	// statusNotificationFrequency    = 30
	// newsNotificationFrequency      = 86400
	// marketingNotificationFrequency = 1200

	statusNotificationFrequency    = 10
	newsNotificationFrequency      = 11
	marketingNotificationFrequency = 12

	createChannel = func() chan entity.Notification {
		return make(chan entity.Notification, 2)
	}

	ctx = context.Background()
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	notifications := generateSampleNotifications()

	go func() {
		for _, notification := range notifications {
			time.Sleep(time.Second)
			notificationsChan <- notification
		}
		// close(notificationsChan)
	}()

	startNotificationProcessors()

	notificationChannels, err := entity.NewNotificationsChannels(statusChan, newsChan, marketingChan)
	if err != nil {
		log.Fatal(err.Error())
	}
	notificationsService := notificationservice.NewNotificationService(notificationsChan, notificationChannels)
	go func() {
		notificationsService.ProcessNotifications(ctx)
		// wg.Done()
	}()

	logrus.Infof("Server up. Receiving notifications")
	wg.Wait()
}

func startNotificationProcessors() {
	emailSender := emailsender.NewEmailSender()
	statusProcessor := processor.NewNotificationProcessor(statusChan, processor.NewNotificationChannelStarter(statusNotificationFrequency, emailSender), createChannel)
	newsProcessor := processor.NewNotificationProcessor(newsChan, processor.NewNotificationChannelStarter(newsNotificationFrequency, emailSender), createChannel)
	marketingProcessor := processor.NewNotificationProcessor(marketingChan, processor.NewNotificationChannelStarter(marketingNotificationFrequency, emailSender), createChannel)

	go func() {
		statusProcessor.Process(ctx)
	}()
	go func() {
		newsProcessor.Process(ctx)
	}()
	go func() {
		marketingProcessor.Process(ctx)
	}()
}

func generateSampleNotifications() []entity.Notification {
	statusNotifications := make([]entity.Notification, 100)
	for i := 0; i < 100; i++ {
		statusNotifications[i] = entity.Notification{
			Type:    values.NotificationTypeStatus,
			Content: "Some status content " + uuid.NewString(),
			Email:   "status@mail.com",
		}
	}
	newsNotifications := make([]entity.Notification, 10)
	for i := 0; i < 10; i++ {
		newsNotifications[i] = entity.Notification{
			Type:    values.NotificationTypeNews,
			Content: "Some news content " + uuid.NewString(),
			Email:   "news@mail.com",
		}
	}
	marketingNotifications := make([]entity.Notification, 50)
	for i := 0; i < 50; i++ {
		marketingNotifications[i] = entity.Notification{
			Type:    values.NotificationTypeMarketing,
			Content: "Some marketing content " + uuid.NewString(),
			Email:   "marketing@mail.com",
		}
	}

	notifications := []entity.Notification{}
	notifications = append(notifications, statusNotifications...)
	notifications = append(notifications, newsNotifications...)
	notifications = append(notifications, marketingNotifications...)

	return notifications

}
