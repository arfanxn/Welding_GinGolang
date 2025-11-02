package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/config"
)

type MailService interface {
	Send(to []string, subject, body string) error
}

type smtpMailService struct {
	host        string
	port        int
	identity    string
	username    string
	password    string
	fromAddress string
	fromName    string
}

func NewSmtpMailServiceFromConfig(cfg *config.Config) MailService {
	return &smtpMailService{
		host:        cfg.MailHost,
		port:        cfg.MailPort,
		identity:    cfg.MailIdentity,
		username:    cfg.MailUsername,
		password:    cfg.MailPassword,
		fromAddress: cfg.MailFromAddress,
		fromName:    cfg.MailFromName,
	}
}

func (smtpMailService) buildMessage(from, fromName string, recipients []string, subject, body string) []byte {
	// Join all recipients with commas
	toHeader := strings.Join(recipients, ", ")
	return fmt.Appendf(nil,
		"From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		fromName, from, toHeader, subject, body,
	)
}

func (s *smtpMailService) Send(recipients []string, subject, body string) error {
	auth := smtp.PlainAuth(s.identity, s.username, s.password, s.host)
	msg := s.buildMessage(s.fromAddress, s.fromName, recipients, subject, body)
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	return smtp.SendMail(addr, auth, s.fromAddress, recipients, msg)
}
