package postman

import (
	"fmt"
	"net/smtp"
)

// PostmanConfig holds SMTP server details
type PostmanConfig struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

// Mailer handles email sending
type Postman struct {
	Config PostmanConfig
}

// Email represents an email to be sent
type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
	IsHTML  bool
}

func NewPostman(config PostmanConfig) Postman {
	return Postman{Config: config}
}

func (pm Postman) Send(email Email) error {
	address := fmt.Sprintf("%s:%d", pm.Config.SMTPHost, pm.Config.SMTPPort)

	contentType := "text/plain; charset=\"UTF-8\""
	if email.IsHTML {
		contentType = "text/html; charset=\"UTF-8\""
	}

	message := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
		email.From,
		email.To[0], // For simplicity, handling one recipient
		email.Subject,
		contentType,
		email.Body,
	)

	auth := smtp.PlainAuth("", pm.Config.Username, pm.Config.Password, pm.Config.SMTPHost)

	err := smtp.SendMail(address, auth, email.From, email.To, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Send mail successfully")
	return nil
}
