package service

import "girhub.com/krifik/bridging-hl7/model"

type FileService interface {
	GetContentFile(url string) model.Json
	GetFiles() []string
	CreateFileResult(request model.JSONRequest) (string, error)
	SearchFile() string
}
