package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	listTickersPath      = "/v3/reference/tickers"
	getTickerDetailsPath = "/v3/reference/tickers/{ticker}"
	getTickerTypesPath   = "/v3/reference/tickers/types"

	getMarketHolidaysPath = "/v1/marketstatus/upcoming"
	getMarketStatusPath   = "/v1/marketstatus/now"

	listSplitsPath = "/v3/reference/splits"

	listDividendsPath = "/v3/reference/dividends"

	listConditionsPath = "/v3/reference/conditions"

	getExchangesPath = "/v3/reference/exchanges"
)

// ReferenceClient defines a REST client for the Polygon reference API.
type ReferenceClient struct {
	client.Client
}

// ListTickersIter is an iterator for the ListTickers method.
type ListTickersIter struct {
	iter.Iter
}

// Ticker returns the current result that the iterator points to.
func (it *ListTickersIter) Ticker() models.Ticker {
	if it.Item() != nil {
		return it.Item().(models.Ticker)
	}
	return models.Ticker{}
}

// ListTickers retrieves reference tickers.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListTickers(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Ticker())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListTickers(ctx context.Context, params *models.ListTickersParams, options ...models.RequestOption) *ListTickersIter {
	return &ListTickersIter{
		Iter: iter.NewIter(ctx, listTickersPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListTickersResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetTickerDetails retrieves details for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
func (c *ReferenceClient) GetTickerDetails(ctx context.Context, params *models.GetTickerDetailsParams, options ...models.RequestOption) (*models.GetTickerDetailsResponse, error) {
	res := &models.GetTickerDetailsResponse{}
	err := c.Call(ctx, http.MethodGet, getTickerDetailsPath, params, res, options...)
	return res, err
}

// GetTickerTypes retrieves all the possible ticker types that can be queried.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers_types.
func (c *ReferenceClient) GetTickerTypes(ctx context.Context, params *models.GetTickerTypesParams, options ...models.RequestOption) (*models.GetTickerTypesResponse, error) {
	res := &models.GetTickerTypesResponse{}
	err := c.Call(ctx, http.MethodGet, getTickerTypesPath, params, res, options...)
	return res, err
}

// GetMarketHolidays retrieves upcoming market holidays and their open/close times.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_upcoming.
func (c *ReferenceClient) GetMarketHolidays(ctx context.Context, options ...models.RequestOption) (*models.GetMarketHolidaysResponse, error) {
	res := &models.GetMarketHolidaysResponse{}
	err := c.CallURL(ctx, http.MethodGet, getMarketHolidaysPath, res, options...)
	return res, err
}

// GetMarketStatus retrieves the current trading status of the exchanges and overall financial markets.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_now.
func (c *ReferenceClient) GetMarketStatus(ctx context.Context, options ...models.RequestOption) (*models.GetMarketStatusResponse, error) {
	res := &models.GetMarketStatusResponse{}
	err := c.CallURL(ctx, http.MethodGet, getMarketStatusPath, res, options...)
	return res, err
}

// ListSplitsIter is an iterator for the ListSplits method.
type ListSplitsIter struct {
	iter.Iter
}

// Split returns the current result that the iterator points to.
func (it *ListSplitsIter) Split() models.Split {
	if it.Item() != nil {
		return it.Item().(models.Split)
	}
	return models.Split{}
}

// ListSplits retrieves reference splits.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_splits.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListSplits(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Split())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListSplits(ctx context.Context, params *models.ListSplitsParams, options ...models.RequestOption) *ListSplitsIter {
	return &ListSplitsIter{
		Iter: iter.NewIter(ctx, listSplitsPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListSplitsResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// ListDividendsIter is an iterator for the ListDividends method.
type ListDividendsIter struct {
	iter.Iter
}

// Dividend returns the current result that the iterator points to.
func (it *ListDividendsIter) Dividend() models.Dividend {
	if it.Item() != nil {
		return it.Item().(models.Dividend)
	}
	return models.Dividend{}
}

// ListDividends retrieves reference dividends.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_dividends.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListDividends(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Dividend())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListDividends(ctx context.Context, params *models.ListDividendsParams, options ...models.RequestOption) *ListDividendsIter {
	return &ListDividendsIter{
		Iter: iter.NewIter(ctx, listDividendsPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListDividendsResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// ListConditionsIter is an iterator for the ListConditions method.
type ListConditionsIter struct {
	iter.Iter
}

// Condition returns the current result that the iterator points to.
func (it *ListConditionsIter) Condition() models.Condition {
	if it.Item() != nil {
		return it.Item().(models.Condition)
	}
	return models.Condition{}
}

// ListConditions retrieves reference conditions.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_conditions.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListConditions(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Condition())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *ReferenceClient) ListConditions(ctx context.Context, params *models.ListConditionsParams, options ...models.RequestOption) *ListConditionsIter {
	return &ListConditionsIter{
		Iter: iter.NewIter(ctx, listConditionsPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListConditionsResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetExchanges lists all exchanges that Polygon knows about.
func (c *ReferenceClient) GetExchanges(ctx context.Context, params *models.GetExchangesParams, options ...models.RequestOption) (*models.GetExchangesResponse, error) {
	res := &models.GetExchangesResponse{}
	err := c.Call(ctx, http.MethodGet, getExchangesPath, params, res, options...)
	return res, err
}
