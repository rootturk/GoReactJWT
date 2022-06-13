package controllers

import (
	"github.com/gofiber/fiber"
)

func Login(c *fiber.Ctx) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(503).Send(err)
		return
	}

	c.JSON(data)
}
