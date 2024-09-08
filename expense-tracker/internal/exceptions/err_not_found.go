package exceptions

type errNotFound struct {
	message string
}

func NewErrNotFound(message string) *errNotFound {
	return &errNotFound{
		message: message,
	}
}

func (e *errNotFound) Error() string {
	return e.message
}
