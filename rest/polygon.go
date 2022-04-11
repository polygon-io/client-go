// Package polygon defines a REST client for the Polygon API.
package polygon

import (
	"github.com/polygon-io/client-go/rest/client"
)

// PolygonClient defines a client to the Polygon REST API.
type PolygonClient struct {
	client.Client
	Aggs      *AggsClient
	Quotes    *QuotesClient
	Reference *ReferenceClient
	Trades    *TradesClient
	Snapshot  *SnapshotClient
}

// New creates a client for the Polygon REST API.
func New(apiKey string) *PolygonClient {
	c := client.New(apiKey)
	return &PolygonClient{
		Client:    c,
		Aggs:      &AggsClient{Client: c},
		Quotes:    &QuotesClient{Client: c},
		Reference: &ReferenceClient{Client: c},
		Trades:    &TradesClient{Client: c},
		Snapshot:  &SnapshotClient{Client: c},
	}
}
