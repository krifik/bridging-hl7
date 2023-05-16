package config

import (
	"bridging-hl7/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfiguration(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
