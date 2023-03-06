package service

import (
	"github.com/ArvarohFX/example-project/internal/app/example/usecase"
	dao "github.com/ArvarohFX/example-project/internal/models/dao/user"
	"github.com/gofiber/fiber/v2"
)

type User interface {
	Create(c *fiber.Ctx) error
}

type userService struct {
	usecase usecase.Usecase
}

func (s *userService) Create(c *fiber.Ctx) error {

	var user *dao.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	s.usecase.User().Create()

	return c.JSON(user)
}
