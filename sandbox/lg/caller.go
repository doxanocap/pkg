package lg

import (
	"os"
	"path"
	"runtime"
	"strings"
)

type CallerInfo struct {
	Host        string `json:"host"`
	PackageName string `json:"package_name"`
	FileName    string `json:"file"`
	FuncName    string `json:"func"`
	Line        int    `json:"line"`
}

func GetHost() string {
	host, err := os.Hostname()
	if err != nil {
		host = "UNDEFINED"
	}
	return host
}

func CallInfo(caller int) *CallerInfo {
	host := GetHost()
	pc, file, line, _ := runtime.Caller(caller)

	_, fileName := path.Split(file)

	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	pl := len(parts)

	packageName := ""
	funcName := parts[pl-1]
	if parts[pl-2] != "" {
		if parts[pl-2][0] == '(' {
			funcName = parts[pl-2] + "." + funcName
			packageName = strings.Join(parts[0:pl-2], ".")
		} else {
			packageName = strings.Join(parts[0:pl-1], ".")
		}
	} else {
		packageName = runtime.FuncForPC(pc).Name()
	}

	return &CallerInfo{
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		Line:        line,
		Host:        host,
	}
}
