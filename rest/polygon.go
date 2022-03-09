package polygon

import (
	"github.com/polygon-io/client-golang/rest/aggregates"
	"github.com/polygon-io/client-golang/rest/client"
)

// todo: add comments for godoc

type polygonClient struct {
	Aggregates *aggregates.Client
	// todo: Trades, Quotes, etc
}

func New(config client.HTTPBaseConfig) *polygonClient {
	c := client.New(config)
	return &polygonClient{
		Aggregates: &aggregates.Client{HTTPBase: c},
	}
}
