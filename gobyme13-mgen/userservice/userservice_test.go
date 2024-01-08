// userservice_test.go
package userservice

import (
	"testing"

	"github.com/akv/me13/database" // Import the database package
	"github.com/akv/me13/mocks"    // Import the generated mocks
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/*
It tests the GetUserName method of a UserService by mocking a database object. It sets up a mock database,
expects a specific user to be queried, and then asserts that the returned username matches the expected value
*/
func TestUserService_GetUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)
	mockDB.EXPECT().GetUser("123").Return(database.User{ID: "123", Name: "Rajanikanth"}, nil)

	userService := NewUserService(mockDB)
	name, err := userService.GetUserName("123")

	assert.NoError(t, err)
	assert.Equal(t, "Rajanikanth", name)
}
