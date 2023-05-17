package repository

import (
	"time"

	"girhub.com/krifik/bridging-hl7/entity"
	"girhub.com/krifik/bridging-hl7/model"

	amdqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	DB   *gorm.DB
	Conn *amdqp.Connection
	Ch   *amdqp.Channel
}

func NewFileRepositoryImpl() FileRepository {
	return &FileRepositoryImpl{}
}

func (f *FileRepositoryImpl) FindLatest() (entity.File, error) {
	var file entity.File
	f.DB.Last(&file)
	return file, nil
}

func (f *FileRepositoryImpl) Insert(request model.FileRequest) error {
	file := entity.File{
		FileName:  request.FileName,
		CreatedAt: time.Now(),
	}
	f.DB.Create(&file)
	return nil
}
