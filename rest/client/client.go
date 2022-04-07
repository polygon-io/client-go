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
func (c *Client) Call(ctx context.Context, method, uri string, params, response interface{}) error {
	req, err := c.NewRequest(ctx, params, response)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

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

func (c *Client) NewRequest(ctx context.Context, params, response interface{}) (*resty.Request, error) {
	req := c.HTTP.R().SetContext(ctx)
	req.SetResult(response).SetError(&models.ErrorResponse{})
	if params == nil {
		return req, nil
	}

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
