package err

type ValidationError struct {
	Err error
}

type BadRequestError struct {
	Err error
}

type AuthorizationError struct {
	Err error
}

func (e *ValidationError) Error() string {
	return "Validation:" + e.Err.Error()
}

func (e *BadRequestError) Error() string {
	return "Bad Request:" + e.Err.Error()
}

func (e *AuthorizationError) Error() string {
	return "Authorization:" + e.Err.Error()
}
