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

func (controller *UserControllerImpl) Insert(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)
	response := controller.UserService.Insert(request)
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
