package services

import (
	"go-mvc/src/domain"
	"go-mvc/src/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
