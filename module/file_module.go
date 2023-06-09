package module

import (
	"github.com/krifik/bridging-hl7/controller"
	"github.com/krifik/bridging-hl7/repository"
	"github.com/krifik/bridging-hl7/service"
)

var fileRepository = repository.NewFileRepositoryImpl()

func NewFileModule() controller.FileController {
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
