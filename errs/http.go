package errs

import (
	"fmt"
)

var BaseErrorHttp = &ErrorHttp{}

type ErrorHttp struct {
	Code int
	Msg  string
}

func NewHttp(code int, msg string) *ErrorHttp {
	return &ErrorHttp{
		Code: code,
		Msg:  msg,
	}
}

func (e *ErrorHttp) Error() string {
	return fmt.Sprintf("code: %d | msg: %s", e.Code, e.Msg)
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

	return
}
