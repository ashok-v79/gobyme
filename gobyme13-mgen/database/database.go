// database.go
package database

type User struct {
	ID   string
	Name string
}

/*
This class definition defines a Database interface with a single method GetUser.
The method takes an id parameter and returns a User object and an erro
*/
type Database interface {
	GetUser(id string) (User, error)
}
