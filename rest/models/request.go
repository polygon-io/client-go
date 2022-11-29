package models

import (
	"net/http"
	"net/url"
)

// RequestOptions are used to configure client calls.
type RequestOptions struct {
	// APIKey to pass with the request
	APIKey *string

	// Headers to apply to the request
	Headers http.Header

	// Query params to apply to the request
	QueryParams url.Values
}

// RequestOption changes the configuration of RequestOptions.
type RequestOption func(o *RequestOptions)

// APIKey sets an APIKey as an option.
func APIKey(id string) RequestOption {
	return func(o *RequestOptions) {
		o.APIKey = &id
	}
}

// Header sets a header as an option.
func Header(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(http.Header)
		}

		o.Headers.Add(key, value)
	}
}

// QueryParam sets a query param as an option.
func QueryParam(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.QueryParams == nil {
			o.QueryParams = make(url.Values)
		}

		o.QueryParams.Add(key, value)
	}
}

// Headers required to use the Launchpad product.
const (
	HeaderEdgeID = "X-Polygon-Edge-ID"
	HeaderEdgeIPAddress = "X-Polygon-Edge-IP-Address "
	HeaderEdgeUserAgent = "X-Polygon-Edge-User-Agent"
)

// EdgeID sets the Launchpad required header denoting the Edge User's ID.
func EdgeID(id string) RequestOption {
	return Header(HeaderEdgeID, id)
}

// EdgeIPAddress sets the Launchpad required header denoting the Edge User's IP Address.
func EdgeIPAddress(ip string) RequestOption {
	return Header(HeaderEdgeIPAddress, ip)
}

// EdgeUserAgent sets the Launchpad required header denoting the Edge User's UserAgent.
func EdgeUserAgent(userAgent string) RequestOption {
	return Header(HeaderEdgeUserAgent, userAgent)
}
