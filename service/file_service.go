package service

import "bridging-hl7/model"

type FileService interface {
	GetContentFile(url string) model.Json
	GetFiles() []string
	CreateFileResult(request model.JSONRequest) (string, error)
	SearchFile() string
}
