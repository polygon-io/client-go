package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListFinancialsPath     = "/vX/reference/financials"
	GetTickerEventsPath    = "/vX/reference/tickers/{id}/events"
	ListTreasuryYieldsPath = "/fed/vX/treasury-yields"
	ListIPOsPath           = "/vX/reference/ipos"
	ListShortInterestPath  = "/stocks/vX/short-interest"
	ListShortVolumePath    = "/stocks/vX/short-volume"
)

// VXClient defines a REST client for the Polygon VX (experimental) API.
type VXClient struct {
	client.Client
}

// ListStockFinancials retrieves historical financial data for a stock ticker. The financials data is extracted from XBRL from company SEC filings
// using the methodology outlined here: http://xbrl.squarespace.com/understanding-sec-xbrl-financi/.
//
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListStockFinancials(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (c *VXClient) ListStockFinancials(ctx context.Context, params *models.ListStockFinancialsParams, options ...models.RequestOption) *iter.Iter[models.StockFinancial] {
	return iter.NewIter(ctx, ListFinancialsPath, params, func(uri string) (iter.ListResponse, []models.StockFinancial, error) {
		res := &models.ListStockFinancialsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetTickerEvents retrieves a timeline of events for the entity associated with the given ticker, CUSIP, or Composite FIGI.
// // For more details see https://polygon.io/docs/stocks/get_vx_reference_tickers__id__events.
func (c *VXClient) GetTickerEvents(ctx context.Context, params *models.GetTickerEventsParams, options ...models.RequestOption) (*models.GetTickerEventsResponse, error) {
	res := &models.GetTickerEventsResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerEventsPath, params, res, options...)
	return res, err
}

// ListTreasuryYields retrieves treasury yield data for U.S. Treasury securities at various maturities.
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
func (c *VXClient) ListTreasuryYields(ctx context.Context, params *models.ListTreasuryYieldsParams, options ...models.RequestOption) *iter.Iter[models.TreasuryYield] {
	return iter.NewIter(ctx, ListTreasuryYieldsPath, params, func(uri string) (iter.ListResponse, []models.TreasuryYield, error) {
		res := &models.ListTreasuryYieldsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListIPOs retrieves detailed information about Initial Public Offerings (IPOs), including both upcoming and historical events.
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
func (c *VXClient) ListIPOs(ctx context.Context, params *models.ListIPOsParams, options ...models.RequestOption) *iter.Iter[models.IPOResult] {
	return iter.NewIter(ctx, ListIPOsPath, params, func(uri string) (iter.ListResponse, []models.IPOResult, error) {
		res := &models.ListIPOsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListShortInterest retrieves short interest data for stocks, including days to cover and average daily volume.
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
func (c *VXClient) ListShortInterest(ctx context.Context, params *models.ListShortInterestParams, options ...models.RequestOption) *iter.Iter[models.ShortInterest] {
	return iter.NewIter(ctx, ListShortInterestPath, params, func(uri string) (iter.ListResponse, []models.ShortInterest, error) {
		res := &models.ListShortInterestResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListShortVolume retrieves short volume data for stocks, including venue-specific volumes and short volume ratio.
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
func (c *VXClient) ListShortVolume(ctx context.Context, params *models.ListShortVolumeParams, options ...models.RequestOption) *iter.Iter[models.ShortVolume] {
	return iter.NewIter(ctx, ListShortVolumePath, params, func(uri string) (iter.ListResponse, []models.ShortVolume, error) {
		res := &models.ListShortVolumeResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
