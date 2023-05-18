package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	GetAllTickersSnapshotPath     = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers"
	GetTickerSnapshotPath         = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers/{ticker}"
	GetGainersLosersSnapshotPath  = "/v2/snapshot/locale/{locale}/markets/{marketType}/{direction}"
	GetOptionContractSnapshotPath = "/v3/snapshot/options/{underlyingAsset}/{optionContract}"
	ListOptionsChainSnapshotPath  = "/v3/snapshot/options/{underlyingAsset}"
	GetCryptoFullBookSnapshotPath = "/v2/snapshot/locale/global/markets/crypto/tickers/{ticker}/book"
	GetIndicesSnapshotPath        = "/v3/snapshot/indices"
	ListAssetSnapshotsPath        = "/v3/snapshot"
)

// SnapshotClient defines a REST client for the Polygon snapshot API.
type SnapshotClient struct {
	client.Client
}

// ListOptionsChainSnapshot retrieves the snapshot of all options contracts for an underlying ticker. For more details see
// https://polygon.io/docs/options/get_v3_snapshot_options__underlyingasset.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListOptionsChainSnapshot(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (ac *SnapshotClient) ListOptionsChainSnapshot(ctx context.Context, params *models.ListOptionsChainParams, options ...models.RequestOption) *iter.Iter[models.OptionContractSnapshot] {
	return iter.NewIter(ctx, ListOptionsChainSnapshotPath, params, func(uri string) (iter.ListResponse, []models.OptionContractSnapshot, error) {
		res := &models.ListOptionsChainSnapshotResponse{}
		err := ac.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetAllTickersSnapshot gets the current minute, day, and previous day's aggregate, as well as the last trade and quote
// for all symbols of a specified market type.
//
// Note: Snapshot data is cleared at 12am EST and gets populated as data is received from the exchanges. This can happen
// as early as 4am EST.
//
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks_tickers.
func (ac *SnapshotClient) GetAllTickersSnapshot(ctx context.Context, params *models.GetAllTickersSnapshotParams, opts ...models.RequestOption) (*models.GetAllTickersSnapshotResponse, error) {
	res := &models.GetAllTickersSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAllTickersSnapshotPath, params, res, opts...)
	return res, err
}

// GetTickerSnapshot gets the current minute, day, and previous day's aggregate, as well as the last trade and quote for
// a single traded symbol of a specified market type.
//
// Note: Snapshot data is cleared at 12am EST and gets populated as data is received from the exchanges. This can happen
// as early as 4am EST.
//
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks_tickers__stocksticker.
func (ac *SnapshotClient) GetTickerSnapshot(ctx context.Context, params *models.GetTickerSnapshotParams, opts ...models.RequestOption) (*models.GetTickerSnapshotResponse, error) {
	res := &models.GetTickerSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetTickerSnapshotPath, params, res, opts...)
	return res, err
}

// GetGainersLosersSnapshot gets the current top 20 gainers or losers of the day in a specific market type.
//
// Top gainers are those tickers whose price has increased by the highest percentage since the previous day's close. Top
// losers are those tickers whose price has decreased by the highest percentage since the previous day's close.
//
// Note: Snapshot data is cleared at 12am EST and gets populated as data is received from the exchanges.
//
// For more details see https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks__direction.
func (ac *SnapshotClient) GetGainersLosersSnapshot(ctx context.Context, params *models.GetGainersLosersSnapshotParams, opts ...models.RequestOption) (*models.GetGainersLosersSnapshotResponse, error) {
	res := &models.GetGainersLosersSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetGainersLosersSnapshotPath, params, res, opts...)
	return res, err
}

// GetOptionContractSnapshot gets the snapshot of an option contract for a stock equity. For more details see
// https://polygon.io/docs/options/get_v3_snapshot_options__underlyingasset___optioncontract.
func (ac *SnapshotClient) GetOptionContractSnapshot(ctx context.Context, params *models.GetOptionContractSnapshotParams, opts ...models.RequestOption) (*models.GetOptionContractSnapshotResponse, error) {
	res := &models.GetOptionContractSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetOptionContractSnapshotPath, params, res, opts...)
	return res, err
}

// GetCryptoFullBookSnapshot gets the current level 2 book of a single cryptocurrency ticker. This is the combined book
// from all of the exchanges.
//
// Note: Snapshot data is cleared at 12am EST and gets populated as data is received from the exchanges.
//
// For more details see
// https://polygon.io/docs/crypto/get_v2_snapshot_locale_global_markets_crypto_tickers__ticker__book.
func (ac *SnapshotClient) GetCryptoFullBookSnapshot(ctx context.Context, params *models.GetCryptoFullBookSnapshotParams, opts ...models.RequestOption) (*models.GetCryptoFullBookSnapshotResponse, error) {
	res := &models.GetCryptoFullBookSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetCryptoFullBookSnapshotPath, params, res, opts...)
	return res, err
}

func (ac *SnapshotClient) GetIndicesSnapshot(ctx context.Context, params *models.GetIndicesSnapshotParams, opts ...models.RequestOption) (*models.GetIndicesSnapshotResponse, error) {
	res := &models.GetIndicesSnapshotResponse{}
	err := ac.Call(ctx, http.MethodGet, GetIndicesSnapshotPath, params, res, opts...)
	return res, err
}

// ListAssetSnapshots retrieves the snapshots for the specified tickers for the specified time. For more details see:
// - https://staging.polygon.io/docs/stocks/get_v3_snapshot
// - https://staging.polygon.io/docs/options/get_v3_snapshot
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListAssetSnapshots(context, params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (ac *SnapshotClient) ListAssetSnapshots(ctx context.Context, params *models.ListAssetSnapshotsParams, options ...models.RequestOption) *iter.Iter[models.SnapshotResponseModel] {
	return iter.NewIter(ctx, ListAssetSnapshotsPath, params, func(uri string) (iter.ListResponse, []models.SnapshotResponseModel, error) {
		res := &models.ListAssetSnapshotsResponse{}
		err := ac.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
