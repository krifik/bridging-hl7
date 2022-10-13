package exception

type UniqueEmailError struct {
	Message string
}

func (uniqueEmailError UniqueEmailError) Error() string {
	return uniqueEmailError.Message
}
