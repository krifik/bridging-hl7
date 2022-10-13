package validation

import "mangojek-backend/exception"

func IsEmailHasBeenTaken(row int64) {
	if row != 0 {
		panic(exception.UniqueEmailError{
			Message: "Email has been taken",
		})
	}
}
