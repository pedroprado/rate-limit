package notificationsapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	interfaces "notification-service/src/core/_interfaces"
)

type notificationsApi struct {
	notificationsService interfaces.NotificationsService
}

func RegisterNotificationsApi(ginRouterGroup *gin.RouterGroup, notificationsService interfaces.NotificationsService) {
	notificationsApi := &notificationsApi{
		notificationsService: notificationsService,
	}
	ginRouterGroup.POST("/notifications", notificationsApi.createNotification)

}

func (api *notificationsApi) createNotification(c *gin.Context) {
	request := CreateNotificationRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	notification := request.ToDomain()
	if err := notification.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	created, err := api.notificationsService.CreateNotification(c, notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, NotificationResponseFromDomain(*created))
}
