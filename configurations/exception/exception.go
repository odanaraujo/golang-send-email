package exception

import "net/http"

type Exception struct {
	Message string
	Err     string
	Code    int
	Causes  []Causes
}

type Causes struct {
	Field   string
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

func NewException(message, err string, code int, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewRequestValidationError(message string, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
