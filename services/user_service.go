package services

import "golang/micro/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}