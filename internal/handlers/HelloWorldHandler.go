package handlers

import "github.com/gofiber/fiber/v2"

func HelloWorldHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(resp)
}
