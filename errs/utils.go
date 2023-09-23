package errs

import (
	"runtime"
	"strings"
)

func callerFunction() string {
	pc, _, _, _ := runtime.Caller(2)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]
	return funcName
}
