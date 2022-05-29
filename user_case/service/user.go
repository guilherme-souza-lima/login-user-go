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
	Validation(tokenString string) (entities.Token, error)
}

type CryptoPassword interface {
	Encrypt(password string) (string, error)
	Decrypt(crypt string) (string, error)
}

type UserService struct {
	UserRepository UserRepository
	JwtToken       JwtToken
	CryptoPassword CryptoPassword
}

func NewUserService(UserRepository UserRepository, JwtToken JwtToken, CryptoPassword CryptoPassword) UserService {
	return UserService{UserRepository, JwtToken, CryptoPassword}
}

func (u UserService) Create(data request.User) error {
	var entity entities.User
	result := entity.ToDomain(data)

	newPassword, err := u.CryptoPassword.Encrypt(data.Password)
	if err != nil {
		return err
	}
	result.Password = newPassword
	return u.UserRepository.Create(result)
}

func (u UserService) Login(data request.Login) (response.UserLogin, error) {
	var entity entities.Login
	var login response.UserLogin
	result := entity.ToDomain(data)

	newPassword, errCrypto := u.CryptoPassword.Encrypt(data.Password)
	if errCrypto != nil {
		return login, errCrypto
	}
	result.Password = newPassword

	user, err := u.UserRepository.Login(result)
	if err != nil {
		return login, err
	}
	login.ID = user.ID
	login.Login = user.Login
	login.Email = user.Email
	login.Name = user.Name
	login.Cellphone = user.Cellphone
	token, errToken := u.JwtToken.Create(login.ID, login.Name, login.Login, login.Email, login.Cellphone)
	if err != nil {
		return login, errToken
	}
	login.Token = token
	return login, nil
}

func (u UserService) Verify(data request.Verify) (bool, error) {
	result, err := u.JwtToken.Validation(data.Token)
	if err != nil {
		return false, err
	}
	if data.ID == result.ID && data.Login == result.Login && data.Name == result.Name &&
		data.Email == result.Email && data.Cellphone == result.Cellphone {
		return true, nil
	}
	return false, nil
}
