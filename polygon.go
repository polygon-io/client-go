// Package polygon defines REST and WebSocket clients for the Polygon API.
package polygon

import (
	"github.com/polygon-io/client-go/rest/aggs"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/quotes"
	"github.com/polygon-io/client-go/rest/reference"
	"github.com/polygon-io/client-go/rest/snapshot"
	"github.com/polygon-io/client-go/rest/trades"
)

type restClient struct {
	client.Client
	Aggs      *aggs.Client
	Quotes    *quotes.Client
	Reference *reference.Client
	Trades    *trades.Client
	Snapshot  *snapshot.Client
}

// NewClient creates a client for the Polygon REST API.
func NewClient(apiKey string) *restClient {
	c := client.New(apiKey)
	return &restClient{
		Client:    c,
		Aggs:      &aggs.Client{Client: c},
		Quotes:    &quotes.Client{Client: c},
		Reference: &reference.Client{Client: c},
		Trades:    &trades.Client{Client: c},
		Snapshot:  &snapshot.Client{Client: c},
	}
}
