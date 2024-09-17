package exceptions

type ErrNotFound struct {
	message string
}

func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{
		message: message,
	}
}

func (self *ErrNotFound) Error() string {
	return self.message
}
