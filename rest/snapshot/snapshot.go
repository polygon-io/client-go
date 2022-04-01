package snapshot

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// Client defines a REST client for the Polygon snapshot API.
type Client struct {
	client.Client
}

// ListSnapshotAllTickers lists the current minute, day, and previous day’s aggregate, as well as the last trade and quote for all symbols of a specified market type.
func (ac *Client) ListSnapshotAllTickers(ctx context.Context, params models.ListSnapshotAllTickersParams, opts ...client.Option) (*models.ListSnapshotAllTickersResponse, error) {
	res := &models.ListSnapshotAllTickersResponse{}
	err := ac.Call(ctx, http.MethodGet, models.ListSnapshotAllTickersPath, params, res, opts...)
	return res, err
}

// GetSnapshotTicker gets the current minute, day, and previous day’s aggregate, as well as the last trade and quote for a single traded symbol of a specified market type.
func (ac *Client) GetSnapshotTicker(ctx context.Context, params models.GetSnapshotTickerParams, opts ...client.Option) (*models.GetSnapshotTickerResponse, error) {
	res := &models.GetSnapshotTickerResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetSnapshotTickerPath, params, res, opts...)
	return res, err
}

// ListSnapshotGainersLosers lists the current top 20 gainers or losers of the day in a specific market type.
func (ac *Client) ListSnapshotGainersLosers(ctx context.Context, params models.ListSnapshotGainersLosersParams, opts ...client.Option) (*models.ListSnapshotGainersLosersResponse, error) {
	res := &models.ListSnapshotGainersLosersResponse{}
	err := ac.Call(ctx, http.MethodGet, models.ListSnapshotGainersLosersPath, params, res, opts...)
	return res, err
}

// GetSnapshotOptionContract gets the snapshot of an option contract for a stock equity.
func (ac *Client) GetSnapshotOptionContract(ctx context.Context, params models.GetSnapshotOptionContractParams, opts ...client.Option) (*models.GetSnapshotOptionContractResponse, error) {
	res := &models.GetSnapshotOptionContractResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetSnapshotOptionContractPath, params, res, opts...)
	return res, err
}

// ListSnapshotTickerFullBook Get the current level 2 book of a single cryptocurrency ticker. This is the combined book from all of the exchanges.
func (ac *Client) ListSnapshotTickerFullBook(ctx context.Context, params models.ListSnapshotTickerFullBookParams, opts ...client.Option) (*models.ListSnapshotTickerFullBookResponse, error) {
	res := &models.ListSnapshotTickerFullBookResponse{}
	err := ac.Call(ctx, http.MethodGet, models.ListSnapshotTickerFullBookPath, params, res, opts...)
	return res, err
}
