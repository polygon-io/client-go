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

// GetAllTickersSnapshot gets the current minute, day, and previous day’s aggregate, as well as the last trade and quote for all symbols of a specified market type.
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks_tickers.
func (ac *Client) GetAllTickersSnapshot(ctx context.Context, params models.GetAllTickersSnapshotParams) (*models.GetAllTickersSnapshotResponse, error) {
	res := &models.GetAllTickersSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetAllTickersSnapshotPath, params, res)
	return res, err
}

// GetTickerSnapshot gets the current minute, day, and previous day’s aggregate, as well as the last trade and quote for a single traded symbol of a specified market type.
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks_tickers__stocksticker.
func (ac *Client) GetTickerSnapshot(ctx context.Context, params models.GetTickerSnapshotParams) (*models.GetTickerSnapshotResponse, error) {
	res := &models.GetTickerSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetTickerSnapshotPath, params, res)
	return res, err
}

// GetGainersLosersSnapshot gets the current top 20 gainers or losers of the day in a specific market type.
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks__direction.
func (ac *Client) GetGainersLosersSnapshot(ctx context.Context, params models.GetGainersLosersSnapshotParams) (*models.GetGainersLosersSnapshotResponse, error) {
	res := &models.GetGainersLosersSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetGainersLosersSnapshotPath, params, res)
	return res, err
}

// GetOptionContractSnapshot gets the snapshot of an option contract for a stock equity.
// For more details see https://polygon.io/docs/options/get_v3_snapshot_options__underlyingasset___optioncontract.
func (ac *Client) GetOptionContractSnapshot(ctx context.Context, params models.GetOptionContractSnapshotParams) (*models.GetOptionContractSnapshotResponse, error) {
	res := &models.GetOptionContractSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetOptionContractSnapshotPath, params, res)
	return res, err
}

// GetCryptoFullBookSnapshot gets the current level 2 book of a single cryptocurrency ticker. This is the combined book from all of the exchanges.
// For more details see https://polygon.io/docs/crypto/get_v2_snapshot_locale_global_markets_crypto_tickers__ticker__book.
func (ac *Client) GetCryptoFullBookSnapshot(ctx context.Context, params models.GetCryptoFullBookSnapshotParams) (*models.GetCryptoFullBookSnapshotResponse, error) {
	res := &models.GetCryptoFullBookSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetCryptoFullBookSnapshotPath, params, res)
	return res, err
}
