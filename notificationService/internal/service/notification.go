package notification

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"net/smtp"
)

type SMTP struct {
	Host     string
	Port     string
	User     string
	Password string
}

type NotificationService struct {
	logger             *zerolog.Logger
	notificationSender NotificationSender
	smtp               SMTP
}

func New(logger *zerolog.Logger, smtpConfig SMTP) *NotificationService {
	return &NotificationService{
		logger: logger,
		smtp:   smtpConfig,
	}
}

type NotificationSender interface {
	SendEmail(ctx context.Context, to, subject, body string) error
	SendNotification(ctx context.Context, userId, subject, body string) error
}

func (ns *NotificationService) SendEmailNotification(ctx context.Context, to, subject, body string) (bool, error) {
	auth := smtp.PlainAuth("", ns.smtp.User, ns.smtp.Password, ns.smtp.Host)
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	addr := fmt.Sprintf("%s:%s", ns.smtp.Host, ns.smtp.Port)
	if err := smtp.SendMail(addr, auth, ns.smtp.User, []string{to}, msg); err != nil {
		ns.logger.Error().Err(err).Msg("Failed to send email")
		return false, err
	}

	ns.logger.Info().Msg("Email sent successfully")
	return true, nil
}

func (ns *NotificationService) SendKafkaNotification(ctx context.Context, userId string, subject string, message string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
