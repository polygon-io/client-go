package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListQuotesPath        = "/v3/quotes/{ticker}"
	GetLastQuotePath      = "/v2/last/nbbo/{ticker}"
	GetLastForexQuotePath = "/v1/last_quote/currencies/{from}/{to}"
)

// QuotesClient defines a REST client for the Polygon quotes API.
type QuotesClient struct {
	client.Client
}

// ListQuotes retrieves quotes for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_quotes__stockticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListQuotes(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *QuotesClient) ListQuotes(ctx context.Context, params *models.ListQuotesParams, options ...models.RequestOption) *iter.Iter[models.Quote] {
	return iter.NewIter(ctx, ListQuotesPath, params, func(uri string) (iter.ListResponse, []models.Quote, error) {
		res := &models.ListQuotesResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetLastQuote retrieves the last quote (NBBO) for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_nbbo__stocksticker.
func (c *QuotesClient) GetLastQuote(ctx context.Context, params *models.GetLastQuoteParams, options ...models.RequestOption) (*models.GetLastQuoteResponse, error) {
	res := &models.GetLastQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, GetLastQuotePath, params, res, options...)
	return res, err
}

// GetLastForexQuote retrieves the last quote (BBO) for a forex currency pair.
// For more details see https://polygon.io/docs/forex/get_v1_last_quote_currencies__from___to.
func (c *QuotesClient) GetLastForexQuote(ctx context.Context, params *models.GetLastForexQuoteParams, options ...models.RequestOption) (*models.GetLastForexQuoteResponse, error) {
	res := &models.GetLastForexQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, GetLastForexQuotePath, params, res, options...)
	return res, err
}
