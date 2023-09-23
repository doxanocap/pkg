# errs

___
errs - is a Go package that provides error handling utilities. It allows you to create, wrap, format, and retrieve error messages with enhanced functionality. These utilities simplify error handling and allow for enhanced error messages and debugging information.

```shell
go get -u github.com/doxanocap/pkg
```

## Import
```go
import "github.com/doxanocap/pkg/errs"
```

## Usage
```go
package main

import (
	"github.com/doxanocap/pkg/errs"
	"log"
)

func main() {
	err := someService()
	if err != nil {
		err = errs.Wrap("wrapping err from service", err)
	}
	log.Println(err)
}

func someService() error {
	return errs.New("service err")
}

```