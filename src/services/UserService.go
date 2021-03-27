package services

import "go-mvc/src/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
