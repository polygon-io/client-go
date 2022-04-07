package polygon

import (
	"github.com/polygon-io/client-go/rest/aggs"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/quotes"
	"github.com/polygon-io/client-go/rest/reference"
	"github.com/polygon-io/client-go/rest/snapshot"
	"github.com/polygon-io/client-go/rest/trades"
)

type polygonClient struct {
	client.Client
	Aggs      *aggs.Client
	Quotes    *quotes.Client
	Reference *reference.Client
	Trades    *trades.Client
	Snapshot  *snapshot.Client
}

// New creates a new client for the Polygon REST API.
func New(apiKey string) *polygonClient {
	c := client.New(apiKey)
	return &polygonClient{
		Aggs:      &aggs.Client{Client: c},
		Client:    c,
		Quotes:    &quotes.Client{Client: c},
		Reference: &reference.Client{Client: c},
		Snapshot:  &snapshot.Client{Client: c},
		Trades:    &trades.Client{Client: c},
	}
}
