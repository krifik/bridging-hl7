package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindAll(c *fiber.Ctx) error
	Insert(c *fiber.Ctx) error
}
