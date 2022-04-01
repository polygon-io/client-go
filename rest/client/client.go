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

// ListResponse defines an interface that list API responses must implement.
type ListResponse interface {
	// NextPageURL returns a URL for retrieving the next page of list results.
	NextPageURL() string
}

// Client provides functionality to make API requests via HTTP.
type Client struct {
	HTTP         *resty.Client
	validate     *validator.Validate
	pathEncoder  *form.Encoder
	queryEncoder *form.Encoder
}

// New returns a new client with the specified API key and default settings.
func New(apiKey string) Client {
	c := resty.New()
	c.SetBaseURL(APIURL)
	c.SetAuthToken(apiKey)
	c.SetRetryCount(DefaultRetryCount)
	c.SetTimeout(10 * time.Second)

	v := validator.New()

	// todo: implement some time types and create specific encoders for them

	pe := form.NewEncoder()
	pe.SetMode(form.ModeExplicit)
	pe.SetTagName("path")
	pe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(x.(time.Time).Format("2006-01-02"))}, nil
	}, time.Time{})

	qe := form.NewEncoder()
	qe.SetMode(form.ModeExplicit)
	qe.SetTagName("query")
	qe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(x.(time.Time).UnixNano())}, nil
	}, time.Time{})

	return Client{
		HTTP:         c,
		validate:     v,
		pathEncoder:  pe,
		queryEncoder: qe,
	}
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, url string, params interface{}, response interface{}, opts ...Option) error {
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

func (c *Client) EncodeParams(uri string, params interface{}) (string, error) {
	if err := c.validateParams(params); err != nil {
		return "", err
	}

	uri, err := c.encodePathString(uri, params)
	if err != nil {
		return "", err
	}

	query, err := c.encodeQueryString(params)
	if err != nil {
		return "", err
	} else if query != "" {
		uri += "?" + query
	}

	return uri, nil
}

func (c *Client) newRequest(ctx context.Context, params interface{}, response interface{}, opts ...Option) (*resty.Request, error) {
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	if params != nil {
		if err := c.validateParams(params); err != nil {
			return nil, err
		}

		path, err := c.encodePath(params)
		if err != nil {
			return nil, err
		}
		req.SetPathParams(path)

		query, err := c.encodeQuery(params)
		if err != nil {
			return nil, err
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

func (c *Client) validateParams(params interface{}) error {
	if err := c.validate.Struct(params); err != nil {
		return fmt.Errorf("invalid request params: %w", err)
	}
	return nil
}

func (c *Client) encodePath(params interface{}) (map[string]string, error) {
	val, err := c.pathEncoder.Encode(&params)
	if err != nil {
		return nil, fmt.Errorf("error encoding path params: %w", err)
	}

	pathParams := map[string]string{}
	for k, v := range val {
		pathParams[k] = v[0] // only accept the first one for a given key
	}
	return pathParams, nil
}

func (c *Client) encodePathString(uri string, params interface{}) (string, error) {
	pathParams, err := c.encodePath(params)
	if err != nil {
		return "", err
	}

	for k, v := range pathParams {
		uri = strings.ReplaceAll(uri, fmt.Sprintf("{%s}", k), url.PathEscape(v))
	}

	return uri, nil
}

func (c *Client) encodeQuery(params interface{}) (url.Values, error) {
	query, err := c.queryEncoder.Encode(&params)
	if err != nil {
		return nil, fmt.Errorf("error encoding query params: %w", err)
	}
	return query, nil
}

func (c *Client) encodeQueryString(params interface{}) (string, error) {
	query, err := c.encodeQuery(params)
	if err != nil {
		return "", nil
	}
	return query.Encode(), nil
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
