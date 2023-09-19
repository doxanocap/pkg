package errs

import (
	"fmt"
	"runtime"
	"strings"
)

var divider = ": "

type err struct {
	s string
}

func (e *err) Error() string {
	return e.s
}

func New(text string) error {
	return &err{s: text}
}

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s%s%s", msg, divider, err.Error())
}

func Newf(format string, input any) error {
	return fmt.Errorf(format, input)
}

func Newl(text string) error {
	return fmt.Errorf(callerFunction() + ":" + text)
}

func Unwrap(err error) string {
	slice := strings.Split(err.Error(), divider)
	return slice[len(slice)-1]
}

func SetWrappingDivider(ch string) {
	divider = ch
}

func callerFunction() string {
	pc, _, _, _ := runtime.Caller(2)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]
	return funcName
}
