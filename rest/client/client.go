package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	APIURL            = "https://api.polygon.io"
	DefaultRetryCount = 3

	HeaderRequestID = "X-Request-ID"
)

// Params defines an interface that path parameter and query parameter types must implement.
type Params interface {
	Path() map[string]string
	Query() map[string]string
}

// Client provides functionality to make API requests via HTTP.
type Client struct {
	HTTP *resty.Client
}

// New returns a new client with the specified API key and default settings.
func New(apiKey string) Client {
	c := resty.New()
	c.SetBaseURL(APIURL)
	c.SetAuthToken(apiKey)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(10 * time.Second)

	return Client{
		HTTP: c,
	}
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (b *Client) Call(ctx context.Context, method, url string, params Params, response interface{}, opts ...Option) error {
	req := b.newRequest(ctx, params, response, opts...)
	res, err := req.Execute(method, url)
	if err != nil {
		return err
	} else if res.IsError() {
		errRes := res.Error().(*ErrorResponse)
		errRes.StatusCode = res.StatusCode()
		return errRes
	}

	return nil
}

func (b *Client) newRequest(ctx context.Context, params Params, response interface{}, opts ...Option) *resty.Request {
	options := mergeOptions(opts...)

	req := b.HTTP.R().SetContext(ctx)
	if params != nil {
		req.SetPathParams(params.Path())
		req.SetQueryParams(params.Query())
	}

	if options.APIKey != nil {
		req.SetAuthToken(*options.APIKey)
	}

	req.SetHeaderMultiValues(options.Headers)

	req.SetResult(response).SetError(&ErrorResponse{})

	return req
}

// BaseResponse has all possible attributes that any response can use. It's intended to be embedded in a
// domain specific response struct.
type BaseResponse struct {
	Status       string `json:"status"`
	RequestID    string `json:"request_id"`
	Count        int    `json:"count,omitempty"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error,omitempty"`

	PaginationHooks
}

type ErrorResponse struct {
	StatusCode int
	BaseResponse
}

// Error returns the details of an error response.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code '%d': message '%s': request ID '%s': internal status: '%s'", e.StatusCode, e.ErrorMessage, e.RequestID, e.Status)
}

// PaginationHooks are links to next and/or previous pages. Embed this struct into your API response if
// your endpoint is going to support pagination.
type PaginationHooks struct {
	NextURL     string `json:"next_url,omitempty"`
	PreviousURL string `json:"previous_url,omitempty"`
}

// Options are used to configure client calls.
type Options struct {
	// APIKey to pass with the request
	APIKey *string

	// Headers to apply to the request
	Headers http.Header
}

// Option changes the configuration of Options.
type Option func(o *Options)

// WithAPIKey sets the APIKey for an Option.
func WithAPIKey(id string) Option {
	return func(o *Options) {
		o.APIKey = &id
	}
}

// WithHeader sets a Header for an Option.
func WithHeader(key, value string) Option {
	return func(o *Options) {
		if o.Headers == nil {
			o.Headers = make(http.Header)
		}

		o.Headers.Add(key, value)
	}
}

func mergeOptions(opts ...Option) *Options {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}

	return options
}
