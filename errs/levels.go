package errs

import (
	"errors"
)

type CustomError interface {
	Error() string
	SetMethod(method string)
	Wrap(message string) CustomError
	New(input interface{}) CustomError
}

type levelErr struct {
	level  string
	method string
	msg    string
}

func (e *levelErr) Error() string {
	return e.msg
}

func NewLevel(level string) CustomError {
	return &levelErr{
		level: level,
	}
}

func ToCustom(err error) CustomError {
	var ce CustomError
	ok := errors.As(err, &ce)
	if ok {
		return ce
	}
	return &levelErr{
		msg: err.Error(),
	}
}

func (e *levelErr) SetMethod(method string) {
	e.method = method
}

func (e *levelErr) New(input interface{}) CustomError {
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

func (e *levelErr) Wrap(message string) CustomError {
	e.constructMsg(message + ": " + e.getMsg())
	return e
}

func (e *levelErr) constructMsg(text string) {
	e.msg = e.level + "." + e.method + " " + divider + " " + text
}

func (e *levelErr) getMsg() string {
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
	return e.msg[idx:]
}
