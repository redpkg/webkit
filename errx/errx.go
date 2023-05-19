package errx

import (
	"fmt"
	"net/http"
)

// var _ error = (*Error)(nil)

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Internal   error  `json:"-"`
}

func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) SetStatusCode(statusCode int) *Error {
	e.StatusCode = statusCode
	return e
}

func (e *Error) SetInternal(err error) *Error {
	e.Internal = err
	return e
}

func (e *Error) Error() string {
	if e.Internal == nil {
		return fmt.Sprintf("(%s) %s", e.Code, e.Message)
	}
	return fmt.Sprintf("(%s) %s | %s", e.Code, e.Message, e.Internal.Error())
}

func (e *Error) Unwrap() error {
	return e.Internal
}

func New(code string, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func Codex(code string, err error) *Error {
	return New(code, "Internal server error").SetInternal(err)
}

func Flatten(err error) []error {
	errs := []error{}

	for {
		errs = append(errs, err)

		u, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			return errs
		}

		if err = u.Unwrap(); err == nil {
			return errs
		}
	}
}
