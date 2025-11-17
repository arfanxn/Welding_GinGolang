package httperror

var _ error = (*HttpError)(nil)

type ErrorsMap map[string][]string

type HttpError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Errors  ErrorsMap `json:"errors,omitempty"`
}

func New(code int, message string, errors ErrorsMap) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func Panic(code int, message string, errors ErrorsMap) {
	panic(New(code, message, errors))
}

func (e HttpError) Error() string {
	return e.Message
}
