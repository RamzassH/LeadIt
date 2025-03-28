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
	logger             zerolog.Logger
	notificationSender NotificationSender
	smtp               SMTP
}

func New(logger zerolog.Logger, smtpConfig SMTP) *NotificationService {
	return &NotificationService{
		logger: logger,
		smtp:   smtpConfig,
	}
}

type NotificationSender interface {
	SendEmailNotification(ctx context.Context, to, subject, body string) (bool, error)
}

func (ns *NotificationService) SendEmailNotification(ctx context.Context, to, subject, body string) (bool, error) {
	if ns.smtp.Host == "" || ns.smtp.Port == "" || ns.smtp.User == "" || ns.smtp.Password == "" {
		ns.logger.Error().Msg("SMTP credentials are empty! Check configuration.")
		return false, fmt.Errorf("SMTP credentials are missing")
	}

	auth := smtp.PlainAuth("", ns.smtp.User, ns.smtp.Password, ns.smtp.Host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		ns.smtp.User, to, subject, body,
	))

	addr := fmt.Sprintf("%s:%s", ns.smtp.Host, ns.smtp.Port)
	if err := smtp.SendMail(addr, auth, ns.smtp.User, []string{to}, msg); err != nil {
		ns.logger.Error().Err(err).Msg("Failed to send email")
		return false, err
	}

	ns.logger.Info().Str("email", to).Msg("Email sent successfully")
	return true, nil
}
