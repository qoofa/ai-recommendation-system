package errors

type Kind string

const (
	BadRequest   Kind = "BAD_REQUEST"
	NotFound     Kind = "NOT_FOUND"
	Conflict     Kind = "CONFLICT"
	Unauthorized Kind = "UNAUTHORIZED"
	Forbidden    Kind = "FORBIDDEN"
	Internal     Kind = "INTERNAL"
)

type AppError struct {
	Kind    Kind
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func New(kind Kind, msg string) *AppError {
	return &AppError{
		Kind:    kind,
		Message: msg,
	}
}

func Wrap(kind Kind, msg string, err error) *AppError {
	return &AppError{
		Kind:    kind,
		Message: msg,
		Err:     err,
	}
}
