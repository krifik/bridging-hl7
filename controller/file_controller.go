package controller

import "github.com/gofiber/fiber/v2"

type FileController interface {
	GetContentFile(c *fiber.Ctx) error
	GetFiles(c *fiber.Ctx) error
	CreateFileResult(c *fiber.Ctx) error
}
