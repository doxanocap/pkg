package gohttp

import (
	"context"
	"github.com/doxanocap/pkg/errs"
	"io"
	"net/http"
)

type Core struct {
	httpClient *http.Client

	url     string
	method  methodType
	headers map[string]string

	requestBody   io.Reader
	requestFormat formatType

	responseBody   interface{}
	responseFormat formatType

	request *http.Request
}

func SetDefaultClient(httpClient *http.Client) *Core {
	if httpClient == nil {
		return nil
	}
	defaultClient = httpClient
	return &Core{
		httpClient: defaultClient,
		headers:    map[string]string{},
	}
}

func NewRequest(client ...*http.Client) *Core {
	if client == nil && len(client) == 0 {
		return &Core{
			httpClient: defaultClient,
		}
	}
	return &Core{
		httpClient: client[0],
		headers:    map[string]string{},
	}
}

func (c *Core) SetURL(url string) *Core {
	if c == nil {
		return nil
	}

	c.url = url
	return c
}

func (c *Core) SetMethod(method methodType) *Core {
	if c == nil {
		return nil
	}

	c.method = method
	return c
}

func (c *Core) SetRequestFormat(format formatType) *Core {
	if c == nil {
		return nil
	}

	c.requestFormat = format
	return c
}

func (c *Core) SetHeader(key string, value string) *Core {
	if c.headers == nil {
		c.headers = map[string]string{}
	}
	c.headers[key] = value
	return c
}

func (c *Core) SetHeaders(headers map[string]string) *Core {
	c.headers = headers
	return c
}

func (c *Core) SetResponseFormat(format formatType) *Core {
	if c == nil {
		return nil
	}

	c.responseFormat = format
	return c
}

func (c *Core) SetRequestBody(requestBody io.Reader) *Core {
	if c == nil {
		return nil
	}

	c.requestBody = requestBody
	return c
}

func (c *Core) SetResponseBody(responseBody interface{}) *Core {
	if c == nil {
		return nil
	}

	c.responseBody = responseBody
	return c
}

func (c *Core) SetRequest(request *http.Request) *Core {
	if c == nil {
		return nil
	}
	c.request = request
	return c
}

func (c *Core) Execute(ctx context.Context) (*http.Response, error) {
	err := c.validateBuilder()
	if err != nil {
		return nil, errs.Wrap("build request: %v", err)
	}

	if c.request == nil {
		c.request, err = c.generateRequest(ctx)
		if err != nil {
			return nil, errs.Wrap("generate request: %v", err)
		}
	}

	response, err := c.httpClient.Do(c.request)
	if err != nil {
		return nil, errs.Wrap("execute request: %v", err)
	}

	if c.responseBody != nil {
		err = decodeResponseBody(response.Body, c.responseFormat, c.responseBody)
		if err != nil {
			return nil, errs.Wrap("decode response body: %v", err)
		}
	}
	return response, nil
}

func (c *Core) generateRequest(ctx context.Context) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, string(c.method), c.url, c.requestBody)
	if err != nil {
		return nil, errs.Wrap("create request: %v", err)
	}

	c.setHeaders(request)
	c.setContentType(request)
	return request, nil
}

func (c *Core) setHeaders(request *http.Request) {
	for key, value := range c.headers {
		request.Header.Set(key, string(value))
	}
}

func (c *Core) setContentType(request *http.Request) {
	contentType := contentTypeByFormat(c.requestFormat)
	request.Header.Set("Content-Type", contentType)
}

func (c *Core) validateBuilder() error {
	if c.url == "" {
		return ErrorEmptyURL
	}

	if !validateMethod(c.method) {
		return ErrorInvalidMethod
	}

	if c.requestFormat == "" {
		c.requestFormat = FormatJSON
	}
	if c.responseFormat == "" {
		c.responseFormat = FormatJSON
	}
	return nil
}
