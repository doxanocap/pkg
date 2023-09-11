package gohttp

import (
	"crypto/tls"
	"github.com/doxanocap/pkg/errs"
	"net/http"
	"time"
)

type (
	formatType string
	methodType string
)

const (
	FormatJSON       formatType = "JSON"
	FormatURLEncoded formatType = "URLENCODED"

	MethodGet     methodType = "GET"
	MethodHead    methodType = "HEAD"
	MethodPost    methodType = "POST"
	MethodPut     methodType = "PUT"
	MethodPatch   methodType = "PATCH"
	MethodDelete  methodType = "DELETE"
	MethodConnect methodType = "CONNECT"
	MethodOptions methodType = "OPTIONS"
	MethodTrace   methodType = "TRACE"
)

var (
	defaultClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	ErrorEmptyURL      error = errs.New("empty url")
	ErrorEmptyMethod   error = errs.New("empty method")
	ErrorInvalidMethod error = errs.New("invalid method")
)

func validateMethod(method methodType) bool {
	return len(method) > 0 && (method == MethodGet ||
		method == MethodPost ||
		method == MethodPut ||
		method == MethodDelete ||
		method == MethodPatch ||
		method == MethodConnect ||
		method == MethodHead ||
		method == MethodOptions ||
		method == MethodTrace)
}
