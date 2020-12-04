package services

import (
	"golang/micro/domain"
	"golang/micro/utils"
)

type userService struct {}

var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// return domain.UserDao.GetUser(userId)
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}