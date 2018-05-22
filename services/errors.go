package services

// NotFoundError ...
type NotFoundError struct {
	message string
}

// NewNotFoundError ...
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		message: message,
	}
}

func (e *NotFoundError) Error() string {
	return e.message
}
