package polygon

import (
	"github.com/polygon-io/client-go/rest/aggs"
	"github.com/polygon-io/client-go/rest/client"
)

type polygonClient struct {
	Aggs *aggs.Client
	// todo: Trades, Quotes, etc
}

// New creates a new client for the Polygon REST API.
func New(apiKey string) *polygonClient {
	c := client.New(apiKey)
	return &polygonClient{
		Aggs: &aggs.Client{Client: c},
	}
}

// Bool returns a pointer to a bool value.
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer to an int value.
func Int(v int) *int {
	return &v
}

// AggsSort returns a pointer to an aggs sort value.
func AggsSort(v aggs.Sort) *aggs.Sort {
	return &v
}
