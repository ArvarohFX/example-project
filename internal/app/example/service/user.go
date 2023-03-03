package service

import (
	dao "github.com/ArvarohFX/example-project/internal/models/dao/user"
	"github.com/gofiber/fiber/v2"
)

type User interface {
	Create(c *fiber.Ctx) error
}

type userService struct {
}

func (s *userService) Create(c *fiber.Ctx) error {

	var user *dao.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.JSON(user)
}
