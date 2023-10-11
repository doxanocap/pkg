package errs

import (
	"log"
	"testing"
)

func TestErrors(t *testing.T) {
	err := someService()
	if err != nil {
		err = Wrap("wrapping err from service", err)
	}
	log.Println(err)
}

func someService() error {
	return New("service err")
}
