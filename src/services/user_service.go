package services

import (
	"go-mvc/src/domain"
	"go-mvc/src/utils"
)

type userService struct{}

var (
	UserService userService
)

func (s *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
