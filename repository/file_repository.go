package repository

import (
	"github.com/krifik/bridging-hl7/entity"
	"github.com/krifik/bridging-hl7/model"
)

type FileRepository interface {
	FindLatest() (entity.File, error)
	Insert(model.FileRequest) error
}
