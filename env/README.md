# env

___

env - is a go package that loads environment variables of the OS and from a specified .env file, then unmarshals them into a provided struct. Simplifies configuration management by reading environment values from a file and mapping them to a Go struct.

```shell
go get -u github.com/doxanocap/pkg
```

## Import
```go
import "github.com/doxanocap/pkg/env"
```


## Usage

```shell
# example .env file
NAME="eldos"
PASSWORD="12345678"

TOKEN_SECRET="secret"
TOKEN_TTL="1d"
TOKEN_ID=4
```

```go
package main

import (
	"github.com/doxanocap/pkg/env"
    "log"
)

type Cfg struct {
	Name     string `env:"NAME"`
	Password string `env:"PASSWORD"`

	Token
}

type Token struct {
	Secret string `env:"TOKEN_SECRET"`
	TTL    string `env:"TOKEN_TTL"`
	ID     int    `env:"TOKEN_ID"`
}

func main() {
	err := env.LoadFile("test.env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Cfg{}
	err = env.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", cfg)
}
```
