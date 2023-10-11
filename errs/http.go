package errs

import (
	"fmt"
	"net/http"
)

var BaseErrorHttp = &ErrorHttp{}

type ErrorHttp struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewHttp(code int, msg string) *ErrorHttp {
	return &ErrorHttp{
		Code:    code,
		Message: msg,
	}
}

func (e *ErrorHttp) Error() string {
	return fmt.Sprintf("code: %d | msg: %s", e.Code, e.Message)
}

func (e *ErrorHttp) NewCode(code int) *ErrorHttp {
	e.Code = code
	return e
}

func UnmarshalCode(err error) (code int) {
	var (
		msg = err.Error()
		n   = 0
	)

	if len(msg) < 8 {
		return
	}

	// len("code:") == 5 -> we start from 6th index
	for i := 6; true; i++ {
		if msg[i] < 48 || msg[i] > 57 {
			break
		}

		n = int(msg[i] - 48)
		code = code*10 + n
	}

	if n == 0 {
		return http.StatusInternalServerError
	}
	return
}

func GetMessage(err error) string {
	var (
		msg = err.Error()
		idx = 0
	)

	if len(msg) < 8 {
		return msg
	}

	for i := 6; true; i++ {
		if msg[i] < 48 || msg[i] > 57 {
			idx = i
			break
		}
	}
	// code: %d | msg: %s <- len of chars after %d till %s
	// equal to 8
	return msg[idx+8:]
}
