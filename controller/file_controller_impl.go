package controller

import (
	"github.com/krifik/bridging-hl7/model"
	"github.com/krifik/bridging-hl7/service"
	"github.com/krifik/bridging-hl7/utils"

	"github.com/gofiber/fiber/v2"
)

type FileControllerImpl struct {
	FileService service.FileService
}

func NewFileControllerImpl(fileService service.FileService) FileController {
	return &FileControllerImpl{FileService: fileService}
}

func (controller *FileControllerImpl) GetContentFile(c *fiber.Ctx) error {
	// var request model.CreateUserRequest
	// err := c.BodyParser(&request)
	// exception.PanicIfNeeded(err)
	response := controller.FileService.GetContentFile("default")
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *FileControllerImpl) GetFiles(c *fiber.Ctx) error {
	// var request model.CreateUserRequest
	// err := c.BodyParser(&request)
	// exception.PanicIfNeeded(err)
	response := controller.FileService.GetFiles()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// CreateFileResult creates a new file and returns a JSON response containing the result.
//
// Parameters:
// - c: The fiber context.
//
// Returns:
// - error: An error if one occurred during the creation of the file, otherwise nil.

func (controller *FileControllerImpl) CreateFileResult(c *fiber.Ctx) error {
	var request model.JSONRequest
	err := c.BodyParser(&request)
	// jsonChan := rabbitmq.Consume()
	if err != nil {
		utils.SendMessage("LINE 54\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
	}
	response, err := controller.FileService.CreateFileResult(request)
	if err != nil {
		utils.SendMessage("LINE 58\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
		return c.JSON(model.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
