// inmemorydb.go
package database

import "fmt"

// InMemoryDB is a simple in-memory database implementation for demonstration
type InMemoryDB struct {
	Users map[string]User
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Users: map[string]User{
			"123": {ID: "123", Name: "Rajanikanth"},
			"456": {ID: "456", Name: "Vijaykanth"},
		},
	}
}

func (db *InMemoryDB) GetUser(id string) (User, error) {
	user, exists := db.Users[id]
	if !exists {
		return User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

/*create a very basic in-memory database implementation that satisfies the Database interface.
In a real-world application, you would connect to an actual database here.*/
