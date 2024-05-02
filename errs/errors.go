package errs

import (
	"fmt"
	"strings"
)

var divider = ": "

type customErr struct {
	s string
}

func (e *customErr) Error() string {
	return e.s
}

// New creates new error interface
func New(text string) error {
	return &customErr{s: text}
}

// Wrap wraps error with message, as a result you get -> "message: err.Error()"
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s%s%s", msg, divider, err.Error())
}

// WrapIfErr wraps to message if error is not nil
// better use with defer
func WrapIfErr(msg string, err *error) {
	if err != nil {
		if *err != nil {
			*err = Wrap(msg, *err)
		}
	}
}

// NewWithCaller creates new error with a function name from which it was called
// 2 steps ago
func NewWithCaller(text string) error {
	return fmt.Errorf(callerFunction() + ":" + text)
}

// Unwrap get initial message after errs.Wrap()
func Unwrap(err error) string {
	slice := strings.Split(err.Error(), divider)
	return slice[len(slice)-1]
}

// SetWrappingDivider default value = ": "
func SetWrappingDivider(ch string) {
	divider = ch
}
