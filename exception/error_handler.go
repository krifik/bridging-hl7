package exception

import (
	"mangojek-backend/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errorNotFound(c, err) {
		return err
	}
	if uniqueEmailError(c, err) {
		return err
	}
	if validationErrors(c, err) {
		return err
	}
	return c.Status(500).JSON(model.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.Error(),
	})
}

func errorNotFound(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		webResponse := model.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error(),
		}
		c.Status(404).JSON(webResponse)
		return true
	} else {
		return false
	}
}
func uniqueEmailError(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(UniqueEmailError)
	if ok {
		webResponse := model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		c.Status(400).JSON(webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(ValidationError)
	if ok {
		webResponse := model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		c.Status(400).JSON(webResponse)
		return true
	} else {
		return false
	}
}
