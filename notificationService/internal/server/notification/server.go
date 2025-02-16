package notification

import (
	"context"
)

type Notification interface {
	sendEmailNotification(ctx context.Context, to string, subject string, body string) (success bool, err error)

	sendKafkaNotification(ctx context.Context, userId string, message string) (success bool, err error)
}

type ServerAPI struct {
}
