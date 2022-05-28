package handler

import (
	"github.com/gofiber/fiber/v2"
	"loginUserGo/user_case/request"
)

type UserService interface {
	Create(data request.User) error
}

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(UserService UserService) UserHandler {
	return UserHandler{UserService}
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user request.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	if err := u.UserService.Create(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error service user. Error: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("success")
}
