package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	APIURL            = "https://api.polygon.io"
	DefaultRetryCount = 3
)

// Params defines an interface that parameter types must implement.
type Params interface {
	// Path returns a map of URL path parameters.
	Path() map[string]string
	// Query returns query string parameters as URL values.
	Query() url.Values
}

// ListParams defines an interface that list parameter types must implement.
type ListParams interface {
	Params
	// String returns a URL string that includes any path and query parameters that are set.
	String() string
}

// ListResponse defines an interface that list API responses must implement.
type ListResponse interface {
	// NextPageURL returns a URL for retrieving the next page of list results.
	NextPageURL() string
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
		req.SetQueryParamsFromValues(params.Query())
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

// PaginationHooks are links to next and/or previous pages. Embed this struct into an API response if
// the endpoint supports pagination.
type PaginationHooks struct {
	NextURL string `json:"next_url,omitempty"`
}

func (p PaginationHooks) NextPageURL() string {
	return p.NextURL
}

// ErrorResponse represents an API response with an error status code.
type ErrorResponse struct {
	StatusCode int
	BaseResponse
}

// Error returns the details of an error response.
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code '%d': message '%s': request ID '%s': internal status: '%s'", e.StatusCode, e.ErrorMessage, e.RequestID, e.Status)
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
