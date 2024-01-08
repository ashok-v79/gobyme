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

// func main() {
// 	paypal := PayPalProcessor{}
// 	stripe := StripeProcessor{}

// 	fmt.Println(ProcessPayment(paypal, 50.0))
// 	fmt.Println(ProcessPayment(stripe, 25.0))
// }

/*
Notice that we didn't have to change the ProcessPayment function or any other part of our system.
The StripeProcessor just plugs into the existing system seamlessly because it adheres to the PaymentProcessor interface.
This demonstrates the power of coding to an interface -
 it allows the system to be extended with new functionalities without modifying the existing codebase.
*/
