package main

import (
	"context"
	"errors"
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/gohttp"
	"github.com/doxanocap/pkg/lg"
)

func main() {
	s := newService1("service1")
	err := s.Execute()
	if err != nil {
		lg.Error(err)
	}

	//gohttp.SetDefaultClient(&http.Client{})

	res, err := gohttp.NewRequest().
		SetURL("http://localhost:8080/healthcheck").
		SetMethod(gohttp.MethodGet).
		SetRequestFormat(gohttp.FormatJSON).
		Execute(context.Background())
	if err != nil {
		lg.Fatal(err)
	}
	lg.Info(res.Status)
}

type service1 struct {
	name  string
	error errs.CustomError
}

func newService1(name string) *service1 {
	return &service1{
		name:  name,
		error: errs.NewLayer("service"),
	}
}

func (s *service1) Execute() error {
	s.error.SetMethod("Execute")

	err := func() error {
		return errors.New("some error happened")
	}()
	if err != nil {
		return s.error.Wrapf("anonymous func err: %s", err.Error()).Wrap("HELLO")
	}

	return nil
}
