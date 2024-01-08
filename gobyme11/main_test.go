// main_test.go (Assuming SendCustomerEmail is in a package named 'email')
package main

import (
	"errors"
	"testing"

	mocks "github.com/akv/mck5/mocks" // Replace with your actual module path for mocks
	"github.com/golang/mock/gomock"   // Replace with your actual module path
)

func TestSendCustomerEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailSender := mocks.NewMockEmailSender(ctrl)
	mockEmailSender.EXPECT().SendEmail("customer@example.com", "Welcome to My App!", "Thank you for signing up for our service.").Return(nil)

	if err := SendCustomerEmail(mockEmailSender, "customer@example.com"); err != nil {
		t.Errorf("Failed to send customer email: %v", err)
	}
}

func TestSendCustomerEmail_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailSender := mocks.NewMockEmailSender(ctrl)
	expectedError := errors.New("failed to send email")
	mockEmailSender.EXPECT().SendEmail("customer@example.com", "Welcome to My App!", "Thank you for signing up for our service.").Return(expectedError)

	err := SendCustomerEmail(mockEmailSender, "customer@example.com")
	if err == nil || err != expectedError {
		t.Errorf("Expected error '%v', got '%v'", expectedError, err)
	}
}
