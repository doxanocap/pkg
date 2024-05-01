package errs

import (
	"errors"
)

type HttpError struct {
	StatusCode int    `json:"-"`
	ErrorCode  string `json:"error_code,omitempty"`

	Message string `json:"message"`
}

func NewHttp(statusCode int, message string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) SetErrorCode(errorCode string) *HttpError {
	e.ErrorCode = errorCode
	return e
}

func UnmarshalError(err error) *HttpError {
	e := &HttpError{}
	if errors.As(err, &e) {
		return e
	}
	e.Message = err.Error()
	return e
}
