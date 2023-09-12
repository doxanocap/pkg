package main

import (
	"context"
	"fmt"
	"github.com/doxanocap/pkg/gohttp"
	"github.com/doxanocap/pkg/lg"
)

func main() {

	req := gohttp.NewRequest()

	res, err := req.
		SetURL("http://localhost:8080/ping").
		SetMethod(gohttp.MethodGet).
		SetHeader("TOKEN", "test123456").
		SetRequestFormat(gohttp.FormatJSON).
		Execute(context.Background())
	if err != nil {
		lg.Fatalf("gohttp: %v", err)
	}
	fmt.Println(res.Status)
}
