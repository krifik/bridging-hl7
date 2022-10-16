package controller

import (
	"mangojek-backend/exception"
	"mangojek-backend/helper"
	"mangojek-backend/model"
	"mangojek-backend/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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
	exception.PanicIfNeeded(err)
	response, err := controller.UserService.Login(request)
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

	claims := jwt.MapClaims{}
	claims["name"] = response.Name
	claims["email"] = response.Email
	claims["expired_at"] = time.Now().Add(60 * time.Minute).Unix()
	token, err := helper.GenerateJWT(&claims)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
func (controller *UserControllerImpl) TestRawSQL(c *fiber.Ctx) error {
	controller.UserService.TestRawSQL()

	return c.SendString("Berhasil")
}
