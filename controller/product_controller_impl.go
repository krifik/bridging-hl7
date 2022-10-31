package controller

import (
	"mangojek-backend/exception"
	"mangojek-backend/model"
	"mangojek-backend/service"

	"github.com/gofiber/fiber/v2"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductControllerImpl(productService service.ProductService) ProductController {
	return &ProductControllerImpl{ProductService: productService}
}
func (controller *ProductControllerImpl) GetProduct(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.ProductService.GetProduct(request.ID)
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
func (controller *ProductControllerImpl) GetProducts(c *fiber.Ctx) error {
	responses := controller.ProductService.GetProducts()
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
func (controller *ProductControllerImpl) Save(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	// pp.Print(request)
	controller.ProductService.Save(request)
	var response model.CreateProductResponse
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
