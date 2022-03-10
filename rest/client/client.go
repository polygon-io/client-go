package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	HeaderRequestID = "X-Request-ID"
)

// todo: add comments for godoc
// todo: could use internal if we don't want outside packages to use this directly

// Params defines an interface that path parameter and query parameter types must implement.
type Params interface {
	Path() map[string]string
	Query() map[string]string
}

// HTTPBase provides functionality to make API requests via HTTP.
type HTTPBase struct {
	rc *resty.Client
}

type HTTPBaseConfig struct {
	URL        string
	Key        string
	MaxRetries int
}

func New(config HTTPBaseConfig) HTTPBase {
	rc := resty.New()
	rc.SetBaseURL(config.URL)
	rc.SetAuthToken(config.Key)
	rc.SetRetryCount(config.MaxRetries)
	rc.SetTimeout(10 * time.Second)

	return HTTPBase{
		rc: rc,
	}
}

func (b *HTTPBase) Call(method, url string, params Params, response json.Unmarshaler, opts ...Option) error {
	req := b.newRequest(params, response, opts...)
	res, err := req.Execute(method, url)
	if err != nil {
		return err
	} else if res.IsError() {
		return newErrorResponse(res)
	}

	return nil
}

func (b *HTTPBase) newRequest(params Params, response json.Unmarshaler, opts ...Option) *resty.Request {
	options := mergeOptions(opts...)

	req := b.rc.R().SetContext(options.Ctx)
	if params != nil {
		req.SetPathParams(params.Path())
		req.SetQueryParams(params.Query())
	}

	if options.RequestID != nil {
		req.SetHeader(HeaderRequestID, *options.RequestID)
	}

	if options.APIKey != nil {
		req.SetAuthToken(*options.APIKey)
	}

	req.SetHeaderMultiValues(options.Headers)

	req.SetResult(response).SetError(&BaseResponse{})

	return req
}

// BaseResponse has all possible attributes that any response can use. It's intended to be embedded in a domain specific
// response struct.
// Ex:
//		type User struct {
//			Name string `json:"name"`
//			Age  int    `json:"age"`
//		}
//
//		type UserAPIResponse struct {
//			apis.BaseResponse
//			User User `json:"results,omitempty"`
// 		}
type BaseResponse struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count,omitempty"`
	Error     string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`

	PaginationHooks
}

// PaginationHooks are links to next and/or previous pages.
// Embed this struct into your API response if your endpoint is going to support pagination.
type PaginationHooks struct {
	NextURL     string `json:"next_url,omitempty"`
	PreviousURL string `json:"previous_url,omitempty"`
}

// ErrorResponse is returned from the client if an undesirable status is returned.
type ErrorResponse struct {
	StatusCode int
	Status     string
	RequestID  string
	Message    string
}

func newErrorResponse(res *resty.Response) *ErrorResponse {
	statusErr := &ErrorResponse{
		StatusCode: res.StatusCode(),
		RequestID:  res.Header().Get(HeaderRequestID),
	}

	errRes := &BaseResponse{}
	err := json.Unmarshal(res.Body(), errRes)
	if err != nil {
		return statusErr // always return status
	}

	statusErr.Status = errRes.Status
	statusErr.Message = errRes.Error

	if statusErr.RequestID == "" && errRes.RequestID != "" {
		statusErr.RequestID = errRes.RequestID
	}

	return statusErr
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("bad status with code: %d; message: %s; request ID: %s; internal status: %s", e.StatusCode, e.Message, e.RequestID, e.Status)
}

// Options are used to configure client calls.
type Options struct {
	// If set, the RequestID will be passed downstream, otherwise the downstream service will create its own RequestID
	RequestID *string

	// APIKey to pass with the request
	APIKey *string

	// Headers to apply to the request
	Headers http.Header

	// Ctx is a user provided context to allow the caller to control the request better
	Ctx context.Context
}

// Option changes the configuration of Options.
type Option func(o *Options)

// WithRequestID sets the RequestID for an Option.
func WithRequestID(id string) Option {
	return func(o *Options) {
		o.RequestID = &id
	}
}

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

// WithContext sets the context for an Option.
func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		switch c := ctx.(type) {
		case *gin.Context:
			o.Ctx = c.Request.Context()
		default:
			o.Ctx = ctx
		}
	}
}

func mergeOptions(opts ...Option) *Options {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}

	return options
}
