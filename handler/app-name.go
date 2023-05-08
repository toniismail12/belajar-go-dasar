package handler

import (
	"github.com/gofiber/fiber/v2"
)

func AppName(c *fiber.Ctx) error {

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Belajar Dasar Golang",
	})

}
