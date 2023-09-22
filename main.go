package main

import (
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/lg"
)

func main() {
	err := errs.New("<qweq>")
	lg.Info("qwe")
	lg.Fatalf("qweqqe %s", err)
}
