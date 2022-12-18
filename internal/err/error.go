package err

type ErrorType int16

const (
	ValidationError ErrorType = iota
	BadRequestError
	AuthorizationError
)

type ApplicationError struct {
	ErrType ErrorType
	Err     error
}

func (e *ApplicationError) Error() string {
	return e.ErrString() + ":" + e.Err.Error()
}

func (e *ApplicationError) ErrString() string {
	switch e.ErrType {
	case ValidationError:
		return "Validation"
	case BadRequestError:
		return "Bad Request"
	case AuthorizationError:
		return "Authorization"
	default:
		return "Unknown"
	}
}
