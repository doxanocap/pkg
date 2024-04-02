package errs

import (
	"errors"
	"net/http"
)

const (
	defaultStatusCode = http.StatusInternalServerError
)

type translation struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type HttpError struct {
	Translations []translation `json:"-"`

	Lang    string `json:"lang,omitempty"`
	Message string `json:"message,omitempty"`

	StatusCode int    `json:"-"`
	ErrorCode  string `json:"error_code,omitempty"`
}

func NewHttp(code int, msg string) *HttpError {
	return &HttpError{
		StatusCode: code,
		Message:    msg,
	}
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) InLanguage(lang string) *HttpError {
	for _, t := range e.Translations {
		if t.Lang == lang {
			e.Lang = t.Lang
			e.Message = t.Value
			return e
		}
	}
	return e
}

func (e *HttpError) AddTranslation(lang string, value string) *HttpError {
	e.Translations = append(e.Translations, translation{lang, value})
	return e
}

func (e *HttpError) SetCode(code int) *HttpError {
	e.StatusCode = code
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
