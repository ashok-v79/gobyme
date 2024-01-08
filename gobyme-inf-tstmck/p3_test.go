// main_test.go

package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockPaymentProcessor is a mock implementation of the PaymentProcessor interface
type MockPaymentProcessor struct {
	mock.Mock
}

func (m *MockPaymentProcessor) ProcessPayment(amount float64) string {
	args := m.Called(amount)
	return args.String(0)
}

func TestProcessPayment(t *testing.T) {
	const amount = 100.0
	const expectedResponse = "Processed $100.00 via Mock"

	// Create a new instance of our mock object
	mockProcessor := new(MockPaymentProcessor)

	// Setup expectations
	mockProcessor.On("ProcessPayment", amount).Return(expectedResponse)

	// Call the function we are testing
	result := ProcessPayment(mockProcessor, amount)

	// Assert that the expectations were met
	mockProcessor.AssertExpectations(t)

	if result != expectedResponse {
		t.Errorf("Expected %v, got %v", expectedResponse, result)
	}
}
