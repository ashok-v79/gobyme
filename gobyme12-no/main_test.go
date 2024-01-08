package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockDB is a mock implementation of the sql.DB interface
type MockDB struct {
	mock.Mock
}

func (mdb *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Define the expected arguments and return values for the Query method
	argsMock := mdb.Called(query, args)
	return argsMock.Get(0).(*sql.Rows), argsMock.Error(1)
}

// TestInteractWithDatabase is a test function that mocks the database and tests the interaction code
func TestInteractWithDatabase(t *testing.T) {
	// Create a new instance of the mock DB
	mockDB := new(MockDB)

	// Define the expected return values for the Query method
	mockDB.On("Query", "SELECT * FROM users").Return(&sql.Rows{}, nil)

	// Call the function that interacts with the database
	interactWithDatabase(mockDB)

	// Assert that the Query method was called with the expected arguments
	mockDB.AssertCalled(t, "Query", "SELECT * FROM users")
}

func interactWithDatabase(db *sql.DB) {
	// Perform a simple query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate over the result rows
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("ID:", id, "Name:", name)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
