package controller

import (
	"mangojek-backend/exception"
	"mangojek-backend/model"
	"mangojek-backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)
	// controller.UserService.Login(request)
	response, _ := controller.UserService.Register(request)
	// if result.Error != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"message": "Email has been taken",
	// 	})
	// }
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	responses, err := controller.UserService.FindAll()
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Login(request)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Wrong credential",
		})
	}
	return c.Status(200).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
