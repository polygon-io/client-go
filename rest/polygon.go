// Package polygon defines a REST client for the Polygon API.
package polygon

import (
	"github.com/polygon-io/client-go/rest/client"
)

// PolygonClient defines a client to the Polygon REST API.
type PolygonClient struct {
	client.Client
	AggsClient
	QuotesClient
	ReferenceClient
	TradesClient
	SnapshotClient
}

// New creates a client for the Polygon REST API.
func New(apiKey string) *PolygonClient {
	c := client.New(apiKey)
	return &PolygonClient{
		Client:          c,
		AggsClient:      AggsClient{Client: c},
		QuotesClient:    QuotesClient{Client: c},
		ReferenceClient: ReferenceClient{Client: c},
		TradesClient:    TradesClient{Client: c},
		SnapshotClient:  SnapshotClient{Client: c},
	}
}
