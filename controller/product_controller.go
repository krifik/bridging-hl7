package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	GetProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
}
