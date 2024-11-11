package failure

import (
	"net/http"
)

// Failure is a custom error wrapper for error messages and codes using standard HTTP response codes
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for the Failure type
func (f *Failure) Error() string {
	return f.Message
}

// BadRequest returns a new Failure with code for bad requests
func BadRequest(msg string) error {
	return &Failure{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

// Unauthorized returns a new Failure with code for unauthorized requests
func Unauthorized(msg string) error {
	return &Failure{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

// Forbidden returns a new Failure with code for forbidden requests
func Forbidden(msg string) error {
	return &Failure{
		Code:    http.StatusForbidden,
		Message: msg,
	}
}

// NotFound returns a new Failure with code for entity not found
func NotFound(msg string) error {
	return &Failure{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

// Conflict returns a new Failure with code for conflict situations
func Conflict(msg string) error {
	return &Failure{
		Code:    http.StatusConflict,
		Message: msg,
	}
}

// InternalError returns a new Failure with code for internal error
func InternalError(msg string) error {
	return &Failure{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

// GetCode returns the error code of a custom error interface
func GetCode(err error) int {
	failure, ok := err.(*Failure) // Assert to Failure type
	if ok {
		return failure.Code
	}

	// Handle the case where the error is not of type *Failure
	return http.StatusInternalServerError
}
