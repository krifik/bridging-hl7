package repository

import (
	"bridging-hl7/entity"
	"bridging-hl7/model"
)

type FileRepository interface {
	FindLatest() (entity.File, error)
	Insert(model.FileRequest) error
}
