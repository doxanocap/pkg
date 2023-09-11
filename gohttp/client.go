package gohttp

import (
	"context"
	"github.com/doxanocap/pkg/errs"
	"net/http"
)

type Core struct {
	httpClient *http.Client

	url            string
	method         methodType
	requestFormat  formatType
	responseFormat formatType
	payload        interface{}
	result         interface{}
}

func SetDefaultClient(httpClient *http.Client) *Core {
	if httpClient == nil {
		return nil
	}
	defaultClient = httpClient
	return &Core{httpClient: defaultClient}
}

func NewRequest(client ...*http.Client) *Core {
	if client == nil && len(client) == 0 {
		return &Core{
			httpClient: defaultClient,
		}
	}
	return &Core{httpClient: client[0]}
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

func (c *Core) SetResponseFormat(format formatType) *Core {
	if c == nil {
		return nil
	}

	c.responseFormat = format
	return c
}

func (c *Core) SetPayload(payload interface{}) *Core {
	if c == nil {
		return nil
	}

	c.payload = payload
	return c
}

func (c *Core) SetResult(result interface{}) *Core {
	if c == nil {
		return nil
	}

	c.result = result
	return c
}

func (c *Core) Execute(ctx context.Context) (*http.Response, error) {
	if err := c.validateBuilder(); err != nil {
		return nil, errs.Wrap("build request: %v", err)
	}

	request, err := c.generateRequest(ctx)
	if err != nil {
		return nil, errs.Wrap("generate request: %v", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, errs.Wrap("execute request: %v", err)
	}
	defer response.Body.Close()

	err = decodeResponseBody(response.Body, c.responseFormat, c.result)
	if err != nil {
		return nil, errs.Wrap("decode response body: %v", err)
	}
	return response, nil
}

func (c *Core) generateRequest(ctx context.Context) (*http.Request, error) {
	requestBody, err := payloadByFormat(c.requestFormat, c.payload)
	if err != nil {
		return nil, errs.Wrap("create request body: %v", err)
	}

	request, err := http.NewRequestWithContext(ctx, string(c.method), c.url, requestBody)
	if err != nil {
		return nil, errs.Wrap("create request: %v", err)
	}

	contentType := contentTypeByFormat(c.requestFormat)
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}

	return request, nil
}

func (c *Core) validateBuilder() error {
	if c.url == "" {
		return ErrorEmptyURL
	}

	if !validateMethod(c.method) {
		return ErrorInvalidMethod
	}

	return nil
}
