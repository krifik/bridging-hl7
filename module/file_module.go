package module

import (
	"girhub.com/krifik/bridging-hl7/controller"
	"girhub.com/krifik/bridging-hl7/repository"
	"girhub.com/krifik/bridging-hl7/service"
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
