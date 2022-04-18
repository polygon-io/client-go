package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListTickersPath      = "/v3/reference/tickers"
	GetTickerDetailsPath = "/v3/reference/tickers/{ticker}"
	GetTickerTypesPath   = "/v3/reference/tickers/types"

	GetMarketHolidaysPath = "/v1/marketstatus/upcoming"
	GetMarketStatusPath   = "/v1/marketstatus/now"

	ListSplitsPath = "/v3/reference/splits"

	ListDividendsPath = "/v3/reference/dividends"

	ListConditionsPath = "/v3/reference/conditions"

	GetExchangesPath = "/v3/reference/exchanges"
)

// ReferenceClient defines a REST client for the Polygon reference API.
type ReferenceClient struct {
	client.Client
}

// ListTickers retrieves reference tickers.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListTickers(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListTickers(ctx context.Context, params *models.ListTickersParams, options ...models.RequestOption) *iter.Iter[models.Ticker] {
	return iter.NewIter(ctx, ListTickersPath, params, func(uri string) (iter.ListResponse, []models.Ticker, error) {
		res := &models.ListTickersResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetTickerDetails retrieves details for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
func (c *ReferenceClient) GetTickerDetails(ctx context.Context, params *models.GetTickerDetailsParams, options ...models.RequestOption) (*models.GetTickerDetailsResponse, error) {
	res := &models.GetTickerDetailsResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerDetailsPath, params, res, options...)
	return res, err
}

// GetTickerTypes retrieves all the possible ticker types that can be queried.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers_types.
func (c *ReferenceClient) GetTickerTypes(ctx context.Context, params *models.GetTickerTypesParams, options ...models.RequestOption) (*models.GetTickerTypesResponse, error) {
	res := &models.GetTickerTypesResponse{}
	err := c.Call(ctx, http.MethodGet, GetTickerTypesPath, params, res, options...)
	return res, err
}

// GetMarketHolidays retrieves upcoming market holidays and their open/close times.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_upcoming.
func (c *ReferenceClient) GetMarketHolidays(ctx context.Context, options ...models.RequestOption) (*models.GetMarketHolidaysResponse, error) {
	res := &models.GetMarketHolidaysResponse{}
	err := c.CallURL(ctx, http.MethodGet, GetMarketHolidaysPath, res, options...)
	return res, err
}

// GetMarketStatus retrieves the current trading status of the exchanges and overall financial markets.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_now.
func (c *ReferenceClient) GetMarketStatus(ctx context.Context, options ...models.RequestOption) (*models.GetMarketStatusResponse, error) {
	res := &models.GetMarketStatusResponse{}
	err := c.CallURL(ctx, http.MethodGet, GetMarketStatusPath, res, options...)
	return res, err
}

// ListSplits retrieves reference splits.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_splits.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListSplits(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListSplits(ctx context.Context, params *models.ListSplitsParams, options ...models.RequestOption) *iter.Iter[models.Split] {
	return iter.NewIter(ctx, ListSplitsPath, params, func(uri string) (iter.ListResponse, []models.Split, error) {
		res := &models.ListSplitsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListDividends retrieves reference dividends.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_dividends.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListDividends(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListDividends(ctx context.Context, params *models.ListDividendsParams, options ...models.RequestOption) *iter.Iter[models.Dividend] {
	return iter.NewIter(ctx, ListDividendsPath, params, func(uri string) (iter.ListResponse, []models.Dividend, error) {
		res := &models.ListDividendsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// ListConditions retrieves reference conditions.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_conditions.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListConditions(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListConditions(ctx context.Context, params *models.ListConditionsParams, options ...models.RequestOption) *iter.Iter[models.Condition] {
	return iter.NewIter(ctx, ListConditionsPath, params, func(uri string) (iter.ListResponse, []models.Condition, error) {
		res := &models.ListConditionsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetExchanges lists all exchanges that Polygon knows about.
func (c *ReferenceClient) GetExchanges(ctx context.Context, params *models.GetExchangesParams, options ...models.RequestOption) (*models.GetExchangesResponse, error) {
	res := &models.GetExchangesResponse{}
	err := c.Call(ctx, http.MethodGet, GetExchangesPath, params, res, options...)
	return res, err
}
