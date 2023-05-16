package module

import (
	"bridging-hl7/controller"
	"bridging-hl7/repository"
	"bridging-hl7/service"

	"gorm.io/gorm"
)

var fileRepository = repository.NewFileRepositoryImpl()

func NewFileModule(database *gorm.DB) controller.FileController {
	// Setup Repository

	// Setup Service
	fileService := service.NewFileServiceImpl(fileRepository)

	// Setup Controller
	fileController := controller.NewFileControllerImpl(fileService)
	return fileController

}

func UseService() service.FileService {

	return service.NewFileServiceImpl(fileRepository)
}
