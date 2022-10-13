package helper

import (
	"mangojek-backend/exception"

	"golang.org/x/crypto/bcrypt"
)

func ToHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	exception.PanicIfNeeded(err)
	return string(hashedPassword)
}
