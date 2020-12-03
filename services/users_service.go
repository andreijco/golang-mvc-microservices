package services

import (
	"golang/micro/domain"
	"golang/micro/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}