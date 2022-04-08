package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"github.com/polygon-io/client-go/rest/models"
)

const (
	APIURL            = "https://api.polygon.io"
	DefaultRetryCount = 3
)

// Client defines an HTTP client for the Polygon REST API.
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

	pe := form.NewEncoder()
	pe.SetMode(form.ModeExplicit)
	pe.SetTagName("path")
	pe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Time)).Format("2006-01-02T15:04:05.000Z"))}, nil
	}, models.Time{})
	pe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Date)).Format("2006-01-02"))}, nil
	}, models.Date{})
	pe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Millis)).UnixMilli())}, nil
	}, models.Millis{})
	pe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Nanos)).UnixNano())}, nil
	}, models.Nanos{})

	qe := form.NewEncoder()
	qe.SetMode(form.ModeExplicit)
	qe.SetTagName("query")
	qe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Time)).Format("2006-01-02T15:04:05.000Z"))}, nil
	}, models.Time{})
	qe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Date)).Format("2006-01-02"))}, nil
	}, models.Date{})
	qe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Millis)).UnixMilli())}, nil
	}, models.Millis{})
	qe.RegisterCustomTypeFunc(func(x interface{}) ([]string, error) {
		return []string{fmt.Sprint(time.Time(x.(models.Nanos)).UnixNano())}, nil
	}, models.Nanos{})

	return Client{
		HTTP:         c,
		validate:     v,
		pathEncoder:  pe,
		queryEncoder: qe,
	}
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, path string, params, response interface{}, opts ...models.RequestOption) error {
	uri, err := c.EncodeParams(path, params)
	if err != nil {
		return err
	}
	return c.CallURL(ctx, method, uri, response, opts...)
}

// CallURL makes an API call based on a request URI and options. The response is automatically unmarshaled.
func (c *Client) CallURL(ctx context.Context, method, uri string, response interface{}, opts ...models.RequestOption) error {
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	if options.APIKey != nil {
		req.SetAuthToken(*options.APIKey)
	}
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetResult(response).SetError(&models.ErrorResponse{})

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

// EncodeParams encodes path and query params and returns a valid request URI.
func (c *Client) EncodeParams(path string, params interface{}) (string, error) {
	if err := c.validateParams(params); err != nil {
		return "", err
	}

	uri, err := c.encodePath(path, params)
	if err != nil {
		return "", err
	}

	query, err := c.encodeQuery(params)
	if err != nil {
		return "", err
	} else if query != "" {
		uri += "?" + query
	}

	return uri, nil
}

func (c *Client) validateParams(params interface{}) error {
	if err := c.validate.Struct(params); err != nil {
		return fmt.Errorf("invalid request params: %w", err)
	}
	return nil
}

func (c *Client) encodePath(uri string, params interface{}) (string, error) {
	val, err := c.pathEncoder.Encode(&params)
	if err != nil {
		return "", fmt.Errorf("error encoding path params: %w", err)
	}

	pathParams := map[string]string{}
	for k, v := range val {
		pathParams[k] = v[0] // only accept the first one for a given key
	}

	for k, v := range pathParams {
		uri = strings.ReplaceAll(uri, fmt.Sprintf("{%s}", k), url.PathEscape(v))
	}

	return uri, nil
}

func (c *Client) encodeQuery(params interface{}) (string, error) {
	query, err := c.queryEncoder.Encode(&params)
	if err != nil {
		return "", fmt.Errorf("error encoding query params: %w", err)
	}
	return query.Encode(), nil
}

func mergeOptions(opts ...models.RequestOption) *models.RequestOptions {
	options := &models.RequestOptions{}
	for _, o := range opts {
		o(options)
	}

	return options
}
