package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const htmlTemplate = `<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f5f5f5;">
<div style="max-width: 600px; margin: 20px auto; padding: 40px 20px; background-color: #012C3D; border-radius: 10px; text-align: center;">
<h1 style="color: #ffffff; margin-bottom: 30px;">Подтверждение регистрации</h1>

<p style="color: #F7F8F3; line-height: 1.6; margin-bottom: 30px;">
Спасибо за регистрацию! Для завершения процесса подтвердите ваш email, нажав на кнопку ниже.
</p>

<a href="http://localhost:8080/v1/auth/verify/%s"
style="display: inline-block; padding: 15px 30px;
background-color: #F8444F; color: #F7F8F3;
text-decoration: none; border-radius: 5px;
font-weight: bold; margin-bottom: 30px;
transition: background-color 0.3s;">
Подтвердить Email
</a>

<p style="color: #636e72; font-size: 14px; line-height: 1.6;">
Если кнопка не работает, скопируйте и вставьте эту ссылку в браузер:<br>
<span style="color: #78BDC4; word-break: break-all;">
http://localhost:8080/v1/auth/verify/%s
</span>
</p>
</div>
</body>`

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
	To      string `validate:"required,email"`
	Subject string `validate:"required"`
	Body    string `validate:"required"`
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
			body := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
    		<meta charset="UTF-8">
    		<meta name="viewport" content="width=device-width, initial-scale=1.0">
			</head>
			%s
			</html>
			`, fmt.Sprintf(htmlTemplate, confirmationMessage.Code, confirmationMessage.Code))

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
		To:      to,
		Subject: subject,
		Body:    body,
	}
	if err := server.validate.Struct(sendReq); err != nil {
		return false, fmt.Errorf("Validation Error: %w", err)
	}

	success, err = server.notification.SendEmailNotification(ctx, to, subject, sendReq.Body)
	if err != nil {
		return false, status.Error(codes.Internal, err.Error())
	}

	return success, nil
}
