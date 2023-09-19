package gohttp

import (
	"bytes"
	"encoding/json"
	"github.com/doxanocap/pkg/errs"
	"io"
	"net/url"
	"strings"
)

func contentTypeByFormat(format formatType) string {
	switch format {
	case FormatJSON:
		return "application/json"
	case FormatURLEncoded:
		return "application/x-www-form-urlencoded"
	default:
		return "application/json"
	}
}

func decodeResponseBody(responseBody io.Reader, format formatType, dst interface{}) error {
	switch format {
	case FormatJSON:
		raw, err := io.ReadAll(responseBody)
		if err != nil {
			return errs.Wrap("reading response body: %v", err)
		}

		return json.Unmarshal(raw, dst)
	}
	return nil
}

func payloadByFormat(format formatType, v interface{}) (payload io.Reader, err error) {
	switch format {
	case FormatJSON:
		if v != nil {
			var raw []byte
			raw, err = json.Marshal(v)
			if err != nil {
				err = errs.Wrap("marshal payload: %v", err)
				return
			}
			payload = bytes.NewReader(raw)
		}
	case FormatURLEncoded:
		if v != nil {
			switch v := v.(type) {
			case url.Values:
				payload = strings.NewReader(v.Encode())
			case *url.Values:
				payload = strings.NewReader(v.Encode())
			case string:
				payload = strings.NewReader(v)
			}
		}
	}

	return
}
