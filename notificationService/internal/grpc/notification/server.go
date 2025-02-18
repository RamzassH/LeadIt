package notification

import (
	"context"
	"fmt"
	notificationv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/notification"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Notification interface {
	SendEmailNotification(ctx context.Context, to string, subject string, body string) (success bool, err error)
	SendKafkaNotification(ctx context.Context, userId string, subject string, message string) (success bool, err error)
}

type ServerAPI struct {
	notificationv1.UnimplementedNotificationServer
	notification Notification
	validate     *validator.Validate
}

type sendEmailNotificationValidation struct {
	to      string `validate:"required,email"`
	subject string `validate:"required"`
	body    string `validate:"required"`
}

type sendKafkaNotificationValidation struct {
	userId  string `validate:"required"`
	message string `validate:"required"`
}

func (server *ServerAPI) ValidateStruct(data interface{}) error {
	if err := server.validate.Struct(data); err != nil {
		return fmt.Errorf("Validation Error: %w", err)
	}
	return nil
}

func RegisterGRPCServer(grpcServer *grpc.Server, validate *validator.Validate, notificationService Notification) {
	notificationv1.RegisterNotificationServer(grpcServer, &ServerAPI{
		notification: notificationService,
		validate:     validate,
	})
}

func (server *ServerAPI) sendEmailNotification(ctx context.Context, to string, subject string, body string) (success bool, err error) {
	sendReq := sendEmailNotificationValidation{
		to:      to,
		subject: subject,
		body:    body,
	}
	if err := server.validate.Struct(sendReq); err != nil {
		return false, fmt.Errorf("Validation Error: %w", err)
	}

	success, err = server.notification.SendEmailNotification(ctx, to, subject, sendReq.body)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	return success, nil
}

func (server *ServerAPI) sendKafkaNotification(ctx context.Context, userId string, subject string, message string) (success bool, err error) {
	sendReq := sendKafkaNotificationValidation{
		userId:  userId,
		message: message,
	}
	if err := server.validate.Struct(sendReq); err != nil {
		return false, fmt.Errorf("Validation Error: %w", err)
	}

	success, err = server.notification.SendKafkaNotification(ctx, userId, subject, sendReq.message)

	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	return success, nil
}
