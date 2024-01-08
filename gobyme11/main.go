package main

import (
	"fmt"
)

// EmailSender interface that defines the method to send an email
type EmailSender interface {
	SendEmail(to, subject, body string) error
}

// RealEmailSender is a real implementation of EmailSender
type RealEmailSender struct{}

// SendEmail is a method to send an email. Here it's just a mock implementation.
func (s *RealEmailSender) SendEmail(to, subject, body string) error {
	fmt.Printf("Email sent to %s: %s - %s\n", to, subject, body)
	// Real email sending logic would be here
	return nil
}

// SendCustomerEmail sends an email to a customer using the EmailSender
func SendCustomerEmail(sender EmailSender, email string) error {
	subject := "Welcome to My App!"
	body := "Thank you for signing up for our service."
	return sender.SendEmail(email, subject, body)
}

// func main() {
// 	sender := &RealEmailSender{}

// 	err := SendCustomerEmail(sender, "customer@example.com")
// 	if err != nil {
// 		log.Fatalf("Failed to send email: %v", err)
// 	}
// }
