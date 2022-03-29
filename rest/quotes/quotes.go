package quotes

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// Client defines a REST client for the Polygon quotes API.
type Client struct {
	client.Client
}

// QuotesIter defines a domain specific iterator for the quotes API.
type QuotesIter struct {
	client.Iter
}

// Quote returns the current result that the iterator points to.
func (it *QuotesIter) Quote() *models.Quote {
	if it.Item() != nil {
		return it.Item().(*models.Quote)
	}
	return nil
}

// ListQuotes retrieves quotes for a specified ticker. This method returns an iterator that should be used to
// access the results via this pattern:
//   iter := c.ListQuotes(context.TODO(), params, opts...)
//   for iter.Next() {
//       // Do something with the current value
//       log.Print(iter.Quote())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListQuotes(ctx context.Context, params models.ListQuotesParams, options ...client.Option) *QuotesIter {
	return &QuotesIter{
		Iter: client.GetIter(ctx, params.String(), func(url string) (client.ListResponse, []interface{}, error) {
			res := &models.QuotesResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetLastQuote retrieves the last quote (NBBO) for a specified ticker.
func (c *Client) GetLastQuote(ctx context.Context, params models.GetLastQuoteParams, options ...client.Option) (*models.LastQuoteResponse, error) {
	res := &models.LastQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastQuotePath, params, res, options...)
	return res, err
}

// GetLastForexQuote retrieves the last quote (BBO) for a forex currency pair.
func (c *Client) GetLastForexQuote(ctx context.Context, params models.LastForexQuoteParams, options ...client.Option) (*models.LastForexQuoteResponse, error) {
	res := &models.LastForexQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastForexQuotePath, params, res, options...)
	return res, err
}
