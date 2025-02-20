package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	interfaces "notification-service/src/core/_interfaces"
	"notification-service/src/core/domain/entity"
	"notification-service/src/core/domain/values"
	emailsender "notification-service/src/core/useCases/emailSender"
	notificationservice "notification-service/src/core/useCases/notification"
	"notification-service/src/core/useCases/notification/processor"
	"notification-service/src/infra"
	"notification-service/src/presentation"
	notificationsapi "notification-service/src/presentation/notificationsApi"
	"notification-service/src/repository"
)

var (
	basePath   = "/notification-service"
	serverPort = "8182"

	notificationsChan = make(chan entity.Notification)
	statusChan        = make(chan entity.Notification)
	newsChan          = make(chan entity.Notification)
	marketingChan     = make(chan entity.Notification)

	statusNotificationFrequencySeconds    string
	newsNotificationFrequencySeconds      string
	marketingNotificationFrequencySeconds string
	firestoreProject                      string

	createChannelForRecipient = func() chan entity.Notification {
		return make(chan entity.Notification, 1)
	}

	ctx context.Context
	db  *firestore.Client
)

// @BasePath /notification-service
func main() {
	loadEnvs()

	ctx = context.Background()
	db = getDB(ctx)
	notificationRepo := repository.NewNotificationRepository(db)

	startNotificationProcessors(notificationRepo)
	notificationChannelsMap, err := entity.NewNotificationsChannelsMap(statusChan, newsChan, marketingChan)
	if err != nil {
		log.Fatal(err.Error())
	}
	notificationsService := notificationservice.NewNotificationsService(notificationRepo, notificationsChan, notificationChannelsMap)

	logrus.Infof("Server up. Receiving notifications")

	server := presentation.NewServerHttpGin(true)
	routerGroup := server.GetGinRouterGroup(basePath)
	presentation.RegisterInfraApi(routerGroup, false)
	notificationsapi.RegisterNotificationsApi(routerGroup, notificationsService)

	server.StartServer(ctx, serverPort)
}

func startNotificationProcessors(notificationRepo interfaces.NotificationRepository) {
	statusNotificationFrequency, err := strconv.Atoi(statusNotificationFrequencySeconds)
	if err != nil {
		log.Fatal(err.Error())
	}
	newsNotificationFrequency, err := strconv.Atoi(newsNotificationFrequencySeconds)
	if err != nil {
		log.Fatal(err.Error())
	}
	marketingNotificationFrequency, err := strconv.Atoi(marketingNotificationFrequencySeconds)
	if err != nil {
		log.Fatal(err.Error())
	}

	emailSender := emailsender.NewEmailSender(infra.NewGoogleSmtpService(), notificationRepo)
	statusProcessor := processor.NewNotificationProcessor(
		notificationRepo,
		statusChan,
		processor.NewNotificationChannelStarter(statusNotificationFrequency, string(values.NotificationTypeStatus), emailSender),
		createChannelForRecipient,
		entity.NewRecipientsChannel(),
	)
	newsProcessor := processor.NewNotificationProcessor(
		notificationRepo,
		newsChan,
		processor.NewNotificationChannelStarter(newsNotificationFrequency, string(values.NotificationTypeNews), emailSender),
		createChannelForRecipient,
		entity.NewRecipientsChannel(),
	)
	marketingProcessor := processor.NewNotificationProcessor(
		notificationRepo,
		marketingChan,
		processor.NewNotificationChannelStarter(marketingNotificationFrequency, string(values.NotificationTypeMarketing), emailSender),
		createChannelForRecipient,
		entity.NewRecipientsChannel(),
	)

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

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	statusNotificationFrequencySeconds = os.Getenv("STATUS_NOTIFICATION_FREQUENCY_SECONDS")
	newsNotificationFrequencySeconds = os.Getenv("NEWS_NOTIFICATION_FREQUENCY_SECONDS")
	marketingNotificationFrequencySeconds = os.Getenv("MARKETING_NOTIFICATION_FREQUENCY_SECONDS")
	firestoreProject = os.Getenv("FIRESTORE_PROJECT")
}

func getDB(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, firestoreProject)
	if err != nil {
		log.Fatal(ctx, "could not create firestore client", err)
	}

	return client
}
