package polygon

import (
	"github.com/polygon-io/client-golang/rest/aggs"
	"github.com/polygon-io/client-golang/rest/client"
)

// todo: add comments for godoc

type polygonClient struct {
	Aggs *aggs.Client
	// todo: Trades, Quotes, etc
}

func New(apiKey string) *polygonClient {
	c := client.New(apiKey)
	return &polygonClient{
		Aggs: &aggs.Client{Client: c},
	}
}
