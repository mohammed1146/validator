package httputils

import (
	"fmt"
	"net/http"
)

// Error copies errors.Error because it can be both json or proto
type Error struct {
	HTTPStatusCode int      `json:"http_status_code"`
	ErrorMessage   string   `json:"error_message"`
	ErrorMessages  []string `json:"error_messages,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf(`{"code": %d, "message": %s, "errors": %v}`,
		e.HTTPStatusCode, e.ErrorMessage, e.ErrorMessages)
}

// NewUnauthorizedError is a quick helper to create a custom httputils.Error
func NewUnauthorizedError(message string) *Error {
	return &Error{
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessage:   message,
	}
}

// NewNotFoundError is a quick helper to create a custom httputils.Error
func NewNotFoundError(message string) *Error {
	return &Error{
		HTTPStatusCode: http.StatusNotFound,
		ErrorMessage:   message,
	}
}

// NewBadRequestError is a helper to create a custom httputils.Error
func NewBadRequestError(message string) *Error {
	return &Error{
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage:   message,
	}
}

// NewRequestEntityTooLargeError is a helper to create a custom httputils.Error
func NewRequestEntityTooLargeError(err error) *Error {
	return &Error{
		HTTPStatusCode: http.StatusRequestEntityTooLarge,
		ErrorMessage:   err.Error(),
	}
}

// NewUnprocessableEntityError is a helper to create a custom httputils.Error
func NewUnprocessableEntityError(err error) *Error {
	return &Error{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		ErrorMessage:   err.Error(),
	}
}

// NewUnprocessableEntityErrors is a helper to create a custom httputils.Error
func NewUnprocessableEntityErrors(err error, errs []string) *Error {
	return &Error{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		ErrorMessage:   err.Error(),
		ErrorMessages:  errs,
	}
}

// NewUnexpectedError is a helper to create a custom httputils.Error
func NewUnexpectedError(err error) *Error {
	return &Error{
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage:   err.Error(),
	}
}

// NewGameOverError is a helper to create a custom httputils.Error
func NewGameOverError(err error) *Error {
	return &Error{
		HTTPStatusCode: http.StatusTeapot,
		ErrorMessage:   err.Error(),
	}
}
