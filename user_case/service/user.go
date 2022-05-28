package service

import (
	"loginUserGo/entities"
	"loginUserGo/infrastructure/model"
	"loginUserGo/user_case/request"
	"loginUserGo/user_case/response"
)

type UserRepository interface {
	Create(data entities.User) error
	Login(data entities.Login) (model.UserModel, error)
}

type JwtToken interface {
	Create(id, name, login, email, cellphone string) (string, error)
}

type UserService struct {
	UserRepository UserRepository
	JwtToken       JwtToken
}

func NewUserService(UserRepository UserRepository, JwtToken JwtToken) UserService {
	return UserService{UserRepository, JwtToken}
}

func (u UserService) Create(data request.User) error {
	var entity entities.User
	result := entity.ToDomain(data)
	return u.UserRepository.Create(result)
}

func (u UserService) Login(data request.Login) (response.UserLogin, error) {
	var entity entities.Login
	var login response.UserLogin
	result := entity.ToDomain(data)

	user, err := u.UserRepository.Login(result)
	if err != nil {
		return login, err
	}
	login.ID = user.ID
	login.Login = user.Login
	login.Email = user.Email
	login.Name = user.Name
	login.Cellphone = user.Password
	token, errToken := u.JwtToken.Create(login.ID, login.Name, login.Login, login.Email, login.Cellphone)
	if err != nil {
		return login, errToken
	}
	login.Token = token
	return login, nil
}
