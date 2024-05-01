# lg - DEPRECATED

lg is a lightweight logging package for Go that provides simple and flexible logging capabilities for your applications. It allows you to log messages at different levels of severity, such as informational messages, warnings, errors, and fatal errors.

## Installation

To use lg in your Go project, you can simply import it and get started:

```shell
go get -u github.com/doxanocap/pkg/lg
```

## Import 
```go
import "github.com/doxanocap/pkg/lg"
```

## Usage

```go
package main

import (
	"github.com/doxanocap/pkg/lg"
)

func main() {
	// Customize the time format (optional)
	lg.SetTimeFormat("2006-01-02T15:04:05Z07:00")

	// Log messages
	lg.Info("This is an informational message")
	lg.Warn("This is a warning message")
	lg.Error(errors.New("An error occurred"))
	lg.Fatalf("Critical error: %s", err.Error())
}
```

