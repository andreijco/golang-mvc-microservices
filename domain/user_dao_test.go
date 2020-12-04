package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := UserDao.GetUser(0)

	// Validation:
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user is not found")
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T)  {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Andrew", user.FirstName)
	assert.EqualValues(t, "Lim", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}