package gohttp

import (
	"crypto/tls"
	"github.com/doxanocap/pkg/errs"
	"net/http"
	"time"
)

type (
	FormatType string
)

const (
	FormatJSON       FormatType = "JSON"
	FormatURLEncoded FormatType = "URLENCODED"
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

	ErrorInvalidURL    error = errs.New("invalid url")
	ErrorEmptyMethod   error = errs.New("empty method")
	ErrorInvalidMethod error = errs.New("invalid method")
)

func validateMethod(method string) bool {
	return len(method) > 0 && (method == http.MethodGet ||
		method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodDelete ||
		method == http.MethodPatch ||
		method == http.MethodConnect ||
		method == http.MethodHead ||
		method == http.MethodOptions ||
		method == http.MethodTrace)
}
