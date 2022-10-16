package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("token")
	if token != "RAHASIA" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "UNAUTHORIZED",
		})
	}
	return c.Next()
}
