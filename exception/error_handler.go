package exception

import (
	"mangojek-backend/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errorNotFound(c, err) {
		return err
	}
	if validationErrors(c, err) {
		return err
	}
	return c.JSON(model.WebResponse{
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
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}
		c.JSON(webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(c *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(ValidationError)
	if ok {

		webResponse := model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error(),
		}

		c.JSON(webResponse)
		return true
	} else {
		return false
	}
}
