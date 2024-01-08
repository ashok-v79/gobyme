// main_test.go

package main

import "testing"

// MockPaymentProcessor is a mock implementation of the PaymentProcessor interface
type MockPaymentProcessor struct {
	mockResponse string
}

func (m MockPaymentProcessor) ProcessPayment(amount float64) string {
	return m.mockResponse
}

func TestProcessPayment(t *testing.T) {
	mockResponse := "Mock Payment Successful"
	mockProcessor := MockPaymentProcessor{mockResponse: mockResponse}

	result := ProcessPayment(mockProcessor, 100.0)
	if result != mockResponse {
		t.Errorf("Expected %v, got %v", mockResponse, result)
	}
}
