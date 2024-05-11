package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
