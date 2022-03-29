package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

type Client struct {
	client.Client
}

// TickersIter defines a domain specific iterator for the reference tickers API.
type TickersIter struct {
	client.Iter
}

// Ticker returns the current result that the iterator points to.
func (it *TickersIter) Ticker() *models.TickerDetails {
	if it.Item() != nil {
		return it.Item().(*models.TickerDetails)
	}
	return nil
}

// ListTickers retrieves reference tickers. This method returns an iterator that should be used to
// access the results via this pattern:
//   iter := c.ListTickers(context.TODO(), params, opts...)
//   for iter.Next() {
//       // Do something with the current value
//       log.Print(iter.Ticker())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListTickers(ctx context.Context, params models.ListTickersParams, options ...client.Option) *TickersIter {
	return &TickersIter{
		Iter: client.GetIter(ctx, params.String(), func(url string) (client.ListResponse, []interface{}, error) {
			res := &models.TickersResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}
