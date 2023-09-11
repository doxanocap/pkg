# gohttp
___
gohttp -  is a lightweight and simple Go package designed to simplify the process of constructing and dispatching HTTP requests.

```shell
go get -u github.com/doxanocap/pkg
```

## Import
```go
import "github.com/doxanocap/pkg/gohttp"
```

## Usage
```go
package main

import (
	"context"
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/gohttp"
	"github.com/doxanocap/pkg/lg"
)

func main() {
	// You can change default http client params
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
```