// userservice.go
package userservice

import (
	"github.com/akv/me13/database" // Update with the actual module name
)

type UserService struct {
	db database.Database
}

/*
This code defines a function called NewUserService that takes a database.Database parameter and returns a pointer
to a UserService struct.

	The UserService struct has a field called db that is initialized with the value of the db parameter.
*/
func NewUserService(db database.Database) *UserService {
	return &UserService{db: db}
}

/*This code defines a method called GetUserName in the UserService struct. It takes an id as input and
returns the corresponding user's name as a string.
If there is an error retrieving the user from the database, it returns an empty string and the error.
*/

func (s *UserService) GetUserName(id string) (string, error) {
	user, err := s.db.GetUser(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
