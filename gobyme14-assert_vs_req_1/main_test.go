// main_test.go
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessData(t *testing.T) {
	// Test case with valid data
	t.Run("Valid Data", func(t *testing.T) {
		inputData := "good data"
		fmt.Println("Testing with valid data:", inputData)

		valid, err := Validate(inputData)
		require.NoError(t, err, "Validation should not return an error")
		require.True(t, valid, "Data should be valid after validation")
		fmt.Println("Validation passed for valid data")

		processedData, err := Process(inputData)
		assert.NoError(t, err, "Processing should not return an error")
		assert.NotNil(t, processedData, "Processed data should not be nil")
		fmt.Println("Processing passed for valid data")

		err = SaveToDatabase(processedData)
		assert.NoError(t, err, "Saving to database should not return an error")
		fmt.Println("Data saved to database successfully")
	})

	// Test case with invalid input
	t.Run("Invalid Input", func(t *testing.T) {
		inputData := ""
		fmt.Println("Testing with invalid input:", inputData)

		valid, err := Validate(inputData)
		fmt.Println("Performing validation for invalid data")
		require.Error(t, err, "Validation should return an error for invalid data")
		require.False(t, valid, "Data should be invalid after validation")
		fmt.Println("Validation correctly identified invalid data")
	})

	// Test case with invalid processing
	t.Run("Invalid Processing", func(t *testing.T) {
		inputData := "invalid"
		fmt.Println("Testing with data that causes processing error:", inputData)

		valid, err := Validate(inputData)
		require.NoError(t, err, "Validation should not return an error")
		require.True(t, valid, "Data should be valid after validation")
		fmt.Println("Validation passed for data that causes processing error")

		_, err = Process(inputData)
		fmt.Println("Performing processing for data that should cause an error")
		assert.Error(t, err, "Processing should return an error")
		fmt.Println("Processing correctly failed as expected")
	})
}
