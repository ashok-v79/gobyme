package main

import (
	"log"
)

// EmailSender is an interface for sending emails
type EmailSender interface {
	SendEmail(to, subject, body string) error
}

// EmailService is a service that uses an EmailSender to send emails
type EmailService struct {
	sender EmailSender
}

// NewEmailService creates a new EmailService with the given EmailSender
func NewEmailService(sender EmailSender) *EmailService {
	return &EmailService{
		sender: sender,
	}
}

// SendEmail sends an email using the EmailService's sender
func (es *EmailService) SendEmail(to, subject, body string) error {
	return es.sender.SendEmail(to, subject, body)
}

func main() {
	// Create a real implementation of the EmailSender interface
	realEmailSender := &RealEmailSender{}

	// Create an instance of the EmailService using the real implementation
	emailService := NewEmailService(realEmailSender)

	// Send an email using the EmailService
	err := emailService.SendEmail("test@example.com", "Test", "This is a test email")
	if err != nil {
		log.Fatal("Failed to send email:", err)
	}

	log.Println("Email sent successfully")
}

// RealEmailSender is a real implementation of the EmailSender interface
type RealEmailSender struct{}

// SendEmail sends an email using a real email sending implementation
func (es *RealEmailSender) SendEmail(to, subject, body string) error {
	// Implement your actual email sending logic here
	log.Printf("Sending email to: %s, Subject: %s, Body: %s", to, subject, body)
	return nil
}
