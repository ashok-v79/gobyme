// main.go

package main

import "fmt"

// PaymentProcessor interface defines behavior for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// PayPalProcessor processes payments using PayPal
type PayPalProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed $%.2f via PayPal", amount)
}

// StripeProcessor processes payments using Stripe
type StripeProcessor struct{}

func (s StripeProcessor) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed $%.2f via Stripe", amount)
}

// ProcessPayment processes a payment using any PaymentProcessor
func ProcessPayment(p PaymentProcessor, amount float64) string {
	return p.ProcessPayment(amount)
}

func main() {
	paypal := PayPalProcessor{}
	stripe := StripeProcessor{}

	fmt.Println(ProcessPayment(paypal, 50.0))
	fmt.Println(ProcessPayment(stripe, 25.0))
}
