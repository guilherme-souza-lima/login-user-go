package service

import (
	"loginUserGo/entities"
	"loginUserGo/user_case/request"
)

type UserRepository interface {
	Create(data entities.User) error
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(UserRepository UserRepository) UserService {
	return UserService{UserRepository}
}

func (u UserService) Create(data request.User) error {
	var entity entities.User
	result := entity.ToDomain(data)
	return u.UserRepository.Create(result)
}
