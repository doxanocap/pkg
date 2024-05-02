package errs

import (
	"fmt"
	"log"
	"testing"
)

func TestErrors(t *testing.T) {
	err := someService()
	if err != nil {
		err = Wrap("wrapping err from service", err)
	}
	log.Println(err)

	v := 2
	v1 := 3
	p := &v
	p = &v1
	fmt.Println(v, *p)

	err = test()
	fmt.Println(err)

	WrapIfErr("repo.FindByID", nil)
}

func someService() error {
	return New("service err")
}

func test() (err error) {
	//defer func() {
	//	if err != nil {
	//		err = errs.Wrap("repo.FindByID", err)
	//	}
	//}()
	defer WrapIfErr("repo.FindByID", &err)

	err = New("not found")
	return
}
