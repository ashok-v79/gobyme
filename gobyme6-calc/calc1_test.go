package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("add(5, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(5, 3)
	expected := 15
	if result != expected {
		t.Errorf("multiply(5, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := divide(6, 3)
	expected := 2
	if result != expected {
		t.Errorf("divide(6, 3) = %d; want %d", result, expected)
	}
}

func TestGetOperation(t *testing.T) {
	tests := []struct {
		userChoice int
		expected   int
	}{
		{1, add(5, 2)},
		{2, subtract(5, 2)},
		{3, multiply(5, 2)},
		{4, divide(5, 2)},
	}

	for _, test := range tests {
		operation := getOperation(test.userChoice)
		result := operation(5, 2)
		if result != test.expected {
			t.Errorf("getOperation(%d)(5, 2) = %d; want %d", test.userChoice, result, test.expected)
		}
	}
}
