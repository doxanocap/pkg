package tools

import (
	"github.com/google/uuid"
	"regexp"
)

const (
	EmailRegexp       = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	PhoneNumberRegexp = `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
)

var (
	emailRegexpFn       = regexp.MustCompile(EmailRegexp)
	phoneNumberRegexpFn = regexp.MustCompile(PhoneNumberRegexp)
)

func GetPtr[T any](v T) *T {
	return &v
}

func IsValidPassword(password string) bool {
	return len(password) > 8 && len(password) < 72
}

func IsUUID(str string) bool {
	_, err := uuid.Parse(str)
	return err == nil
}

func IsValidEmail(e string) bool {
	return emailRegexpFn.MatchString(e)
}

func IsValidPhoneNumber(e string) bool {
	return phoneNumberRegexpFn.MatchString(e)
}
