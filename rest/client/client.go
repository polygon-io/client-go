package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
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
}

// ListResponse defines an interface that list API responses must implement.
type ListResponse interface {
	// NextPageURL returns a URL for retrieving the next page of list results.
	NextPageURL() string
}

// Client provides functionality to make API requests via HTTP.
type Client struct {
	HTTP     *resty.Client
	validate *validator.Validate
	encoder  *form.Encoder
}

// New returns a new client with the specified API key and default settings.
func New(apiKey string) Client {
	c := resty.New()
	c.SetBaseURL(APIURL)
	c.SetAuthToken(apiKey)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(10 * time.Second)

	v := validator.New()

	e := form.NewEncoder()
	e.SetMode(form.ModeExplicit)
	e.SetTagName("query")
	e.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(x.(time.Time).UnixNano())}, nil
	}, time.Time{})

	return Client{
		HTTP:     c,
		validate: v,
		encoder:  e,
	}
}

// NewIter returns a new initialized iterator. This method automatically makes the first query to populate
// the results. List methods should use this helper method when building domain specific iterators.
func (c *Client) NewIter(ctx context.Context, uri string, params Params, query Query) (*Iter, error) {
	it := &Iter{
		ctx:   ctx,
		query: query,
	}

	if err := c.validate.Struct(params); err != nil {
		return nil, fmt.Errorf("invalid request params: %w", err)
	}

	path := params.Path()
	for k, v := range path {
		uri = strings.ReplaceAll(uri, fmt.Sprintf("{%s}", k), url.PathEscape(v))
	}

	q, err := c.encoder.Encode(&params)
	if err != nil {
		return nil, fmt.Errorf("error encoding request params: %w", err)
	} else if q.Encode() != "" {
		uri += "?" + q.Encode()
	}

	it.page, it.results, it.err = it.query(uri)
	return it, nil
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, url string, params Params, response interface{}, opts ...Option) error {
	req, err := c.newRequest(ctx, params, response, opts...)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	res, err := req.Execute(method, url)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	} else if res.IsError() {
		errRes := res.Error().(*ErrorResponse)
		errRes.StatusCode = res.StatusCode()
		return errRes
	}

	return nil
}

func (c *Client) newRequest(ctx context.Context, params Params, response interface{}, opts ...Option) (*resty.Request, error) {
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	if params != nil {
		if err := c.validate.Struct(params); err != nil {
			return nil, fmt.Errorf("invalid request params: %w", err)
		}

		req.SetPathParams(params.Path())

		query, err := c.encoder.Encode(&params)
		if err != nil {
			return nil, fmt.Errorf("error encoding request params: %w", err)
		}
		req.SetQueryParamsFromValues(query)
	}

	if options.APIKey != nil {
		req.SetAuthToken(*options.APIKey)
	}

	req.SetHeaderMultiValues(options.Headers)

	req.SetResult(response).SetError(&ErrorResponse{})

	return req, nil
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
