package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kafka"
)

type ConfirmationMessage struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
	Time   string `json:"timestamp"`
}

type Notification interface {
	SendEmailNotification(ctx context.Context, to string, subject string, body string) (success bool, err error)
}

type ServerAPI struct {
	notification  Notification
	logger        zerolog.Logger
	kafkaConsumer *kafka.Consumer
	validate      *validator.Validate
}

func New(notification Notification, logger zerolog.Logger, kafkaConsumer *kafka.Consumer, validate *validator.Validate) *ServerAPI {
	return &ServerAPI{
		notification:  notification,
		logger:        logger,
		kafkaConsumer: kafkaConsumer,
		validate:      validate,
	}
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

func (server *ServerAPI) StartKafkaConsumer() {
	go func() {
		ctx := context.Background()
		err := server.kafkaConsumer.ReadMessage(ctx, func(msg kafka.Message) error {
			var confirmationMessage ConfirmationMessage
			if err := json.Unmarshal(msg.Value, &confirmationMessage); err != nil {
				server.logger.Error().Err(err).Msg("Error unmarshalling confirmation message")
				return err
			}

			subject := "Verification code"
			body := fmt.Sprintf("Hello, your verification code is: %s", confirmationMessage.Code)

			success, err := server.notification.SendEmailNotification(ctx, confirmationMessage.Email, subject, body)

			if err != nil {
				server.logger.Error().Err(err).Msg("Error sending email notification")
				return err
			}
			if success {
				server.logger.Info().Msg("Email notification sent")
			}
			return nil
		})

		if err != nil {
			server.logger.Error().Err(err).Msg("Error while reading messages")
		}
	}()
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
