package errs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultStatusCode = http.StatusInternalServerError
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

func SetGinError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	_ = ctx.Error(err)
	httpError := UnmarshalError(err)
	if httpError.StatusCode != defaultStatusCode {
		ctx.JSON(httpError.StatusCode, httpError)
	} else {
		ctx.Status(httpError.StatusCode)
	}
	return
}

func SetGinErrorWithStatus(ctx *gin.Context, status int, err error) {
	if err == nil {
		return
	}
	_ = ctx.Error(err)
	ctx.Status(status)
	return
}
