package client

import (
	"net/http"
)

type (
	Error struct {
		name       string
		statusCode int
		msg        string
	}
)

var (
	ErrBadRequest    error = &Error{name: "BadRequest", statusCode: http.StatusBadRequest}
	ErrUnauthorized  error = &Error{name: "Unauthorized", statusCode: http.StatusUnauthorized}
	ErrForbidden     error = &Error{name: "Forbidden", statusCode: http.StatusForbidden}
	ErrNotFound      error = &Error{name: "NotFound", statusCode: http.StatusNotFound}
	ErrConflict      error = &Error{name: "Conflict", statusCode: http.StatusConflict}
	ErrInternalError error = &Error{name: "InternalError", statusCode: http.StatusInternalServerError}
)

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Message(msg string) {
	e.msg = msg
}

func NewErrorFromStatusCode(status int) *Error {
	switch status {
	case http.StatusBadRequest:
		return ErrBadRequest.(*Error)
	case http.StatusUnauthorized:
		return ErrUnauthorized.(*Error)
	case http.StatusForbidden:
		return ErrForbidden.(*Error)
	case http.StatusNotFound:
		return ErrNotFound.(*Error)
	case http.StatusConflict:
		return ErrConflict.(*Error)
	case http.StatusInternalServerError:
		return ErrInternalError.(*Error)
	default:
		return ErrInternalError.(*Error)
	}
}
