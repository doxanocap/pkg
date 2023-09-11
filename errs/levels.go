package errs

import (
	"errors"
	"fmt"
)

type CustomError interface {
	Error() string
	SetMethod(method string)
	Wrap(message string) CustomError
	Wrapf(format string, input any) CustomError
	WrapNew(err error, message string) CustomError
	New(input interface{}) CustomError
}

type levelError struct {
	levelName string
	method    string
	msg       string
}

func (e *levelError) Error() string {
	return e.msg
}

func NewLayer(levelName string) CustomError {
	return &levelError{
		levelName: levelName,
	}
}

func ToCustom(err error) CustomError {
	var ce CustomError
	ok := errors.As(err, &ce)
	if ok {
		return ce
	}
	return &levelError{
		msg: err.Error(),
	}
}

func (e *levelError) SetMethod(method string) {
	e.method = method
}

func (e *levelError) New(input interface{}) CustomError {
	var msg string

	switch v := input.(type) {
	case error:
		msg = v.Error()
	case string:
		msg = input.(string)
	default:
		return nil
	}

	e.constructMsg(msg)
	return e
}

func (e *levelError) Wrap(message string) CustomError {
	e.constructMsg(message + divider + e.getMsg())
	return e
}

func (e *levelError) Wrapf(format string, input any) CustomError {
	e.constructMsg(fmt.Sprintf(format, input))
	return e
}

func (e *levelError) WrapNew(err error, message string) CustomError {
	e.constructMsg(message + ": " + err.Error())
	return e
}

func (e *levelError) constructMsg(text string) {
	e.msg = e.levelName + "." + e.method + divider + text
}

func (e *levelError) getMsg() string {
	var idx = 0
	for i := range e.msg {
		var ok = true
		for j := range divider {
			if i+j < len(e.msg) && divider[j] != e.msg[i+j] {
				ok = false
				break
			}
		}
		if ok {
			idx = i + len(divider) + 1
			break
		}
	}

	if len(e.msg) <= idx {
		return ""
	}
	return e.msg[idx-1:]
}
