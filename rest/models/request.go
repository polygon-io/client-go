package models

import "net/http"

// RequestOptions are used to configure client calls.
type RequestOptions struct {
	// APIKey to pass with the request
	APIKey *string

	// Headers to apply to the request
	Headers http.Header

	// Query params to apply to the request
	QueryParams map[string]string
}

// RequestOption changes the configuration of RequestOptions.
type RequestOption func(o *RequestOptions)

// WithAPIKey sets the APIKey for an Option.
func WithAPIKey(id string) RequestOption {
	return func(o *RequestOptions) {
		o.APIKey = &id
	}
}

// WithHeader sets a Header for an Option.
func WithHeader(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(http.Header)
		}

		o.Headers.Add(key, value)
	}
}

// WithHeader sets a Header for an Option.
func WithQueryParam(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.QueryParams == nil {
			o.QueryParams = make(map[string]string)
		}

		o.QueryParams[key] = value
	}
}
