package values

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidaNotificationTeyps(t *testing.T) {

	cases := map[string]struct {
		notificationType NotificationType
		expectedErr      error
	}{
		"should-be-valid-status-type": {
			notificationType: NotificationTypeStatus,
			expectedErr:      nil,
		},
		"should-be-valid-news-type": {
			notificationType: NotificationTypeNews,
			expectedErr:      nil,
		},
		"should-be-valid-marketing-type": {
			notificationType: NotificationTypeMarketing,
			expectedErr:      nil,
		},
		"should-be-invalid": {
			notificationType: NotificationType("INVALID_TYPE"),
			expectedErr:      errors.New("INVALID_TYPE is an invalid notification type"),
		},
	}

	for title, c := range cases {
		t.Run(title, func(t *testing.T) {

			err := c.notificationType.Validate()

			assert.Equal(t, c.expectedErr, err)
		})
	}
}
