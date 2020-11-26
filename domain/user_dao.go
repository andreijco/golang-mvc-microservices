package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Andrew", LastName: "Lim", Email: "myemail.@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	
	return nil, fmt.Errorf("user %v was not found", userId)
}