// main.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Process Data Application")
}

func Validate(input string) (bool, error) {
	if input == "" {
		return false, errors.New("input is empty")
	}

	return true, nil
}

func Process(input string) (string, error) {
	if input == "invalid" {
		return "", errors.New("invalid processing")

	}
	return "Processed: " + input, nil
}

func SaveToDatabase(data string) error {
	if data == "Processed: bad" {
		return errors.New("failed to save")
	}
	return nil
}
