// Package polygon defines a REST client for the Polygon API.
package polygon

import (
	"github.com/polygon-io/client-go/rest/client"
)

// Client defines a client to the Polygon REST API.
type Client struct {
	client.Client
	AggsClient
	QuotesClient
	ReferenceClient
	TradesClient
	SnapshotClient
	VX VXClient
}

// New creates a client for the Polygon REST API.
func New(apiKey string) *Client {
	c := client.New(apiKey)
	return &Client{
		Client:          c,
		AggsClient:      AggsClient{Client: c},
		QuotesClient:    QuotesClient{Client: c},
		ReferenceClient: ReferenceClient{Client: c},
		TradesClient:    TradesClient{Client: c},
		SnapshotClient:  SnapshotClient{Client: c},
		VX:              VXClient{Client: c},
	}
}
