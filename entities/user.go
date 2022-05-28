package entities

import (
	"github.com/gofrs/uuid"
	"loginUserGo/user_case/request"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
}

func (u User) ToDomain(data request.User) User {
	uuidGenerator, _ := uuid.NewV4()
	return User{
		ID:        uuidGenerator.String(),
		Name:      data.Name,
		Login:     data.Login,
		Password:  data.Password,
		Email:     data.Email,
		Cellphone: data.Cellphone,
	}
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (l Login) ToDomain(data request.Login) Login {
	return Login{
		Login:    data.Login,
		Password: data.Password,
	}
}
