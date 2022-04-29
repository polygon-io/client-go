package client

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/polygon-io/client-go/rest/encoder"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	APIURL            = "https://api.polygon.io"
	DefaultRetryCount = 3
)

// Client defines an HTTP client for the Polygon REST API.
type Client struct {
	HTTP    *resty.Client
	encoder *encoder.Encoder
}

// New returns a new client with the specified API key and default settings.
func New(apiKey string) Client {
	c := resty.New()
	c.SetBaseURL(APIURL)
	c.SetAuthToken(apiKey)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(10 * time.Second)
	c.SetHeader("User-Agent", "Go client")

	return Client{
		HTTP:    c,
		encoder: encoder.New(),
	}
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, path string, params, response any, opts ...models.RequestOption) error {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return err
	}
	return c.CallURL(ctx, method, uri, response, opts...)
}

// CallURL makes an API call based on a request URI and options. The response is automatically unmarshaled.
func (c *Client) CallURL(ctx context.Context, method, uri string, response any, opts ...models.RequestOption) error {
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	if options.APIKey != nil {
		req.SetAuthToken(*options.APIKey)
	}
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetResult(response).SetError(&models.ErrorResponse{})

	// todo: add additional headers (e.g. user agent)

	res, err := req.Execute(method, uri)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	} else if res.IsError() {
		errRes := res.Error().(*models.ErrorResponse)
		errRes.StatusCode = res.StatusCode()
		return errRes
	}

	return nil
}

func mergeOptions(opts ...models.RequestOption) *models.RequestOptions {
	options := &models.RequestOptions{}
	for _, o := range opts {
		o(options)
	}

	return options
}
