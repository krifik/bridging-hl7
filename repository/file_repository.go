package repository

import (
	"girhub.com/krifik/bridging-hl7/entity"
	"girhub.com/krifik/bridging-hl7/model"
)

type FileRepository interface {
	FindLatest() (entity.File, error)
	Insert(model.FileRequest) error
}
