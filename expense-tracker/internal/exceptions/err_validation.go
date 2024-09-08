package exceptions

type ErrValidation struct {
	message string
}

func NewErrValidation(message string) *ErrValidation {
	return &ErrValidation{
		message: message,
	}
}

func (e *ErrValidation) Error() string {
	return e.message
}
