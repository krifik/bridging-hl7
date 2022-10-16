package middleware

import (
	"mangojek-backend/exception"
	"mangojek-backend/helper"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "UNAUTENTICATE",
		})
	}

	_, err := helper.VerifyToken(token)
	exception.PanicIfNeeded(err)
	claims, err := helper.DecodeToken(token)
	expired_at := int64(claims["expired_at"].(float64))
	if expired_at < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "EXPIRED TOKEN",
		})
	}
	return c.Next()
}
