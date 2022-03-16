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
	*client.Iter
}

// Quote returns the current result that the iterator points to.
func (it *QuotesIter) Quote() *models.Quote {
	if it.Current() != nil {
		return it.Current().(*models.Quote)
	}
	return nil
}

// QuotesList returns the current page of results.
func (it *QuotesIter) QuotesList() *models.QuotesResponse {
	if it.Page() != nil {
		return it.Page().(*models.QuotesResponse)
	}
	return nil
}

// ListQuotes retrieves quotes for a specified ticker. This method returns an iterator that should be used to
// access the results via this pattern:
//   iter := c.ListQuotes(context.TODO(), params, opts...)
//	 for iter.Next() {
//		 // Do something with the current value
//  	 log.Print(iter.Quote())
//	 }
//	 if iter.Err() != nil {
//	 	return err
//	 }
func (c *Client) ListQuotes(ctx context.Context, params models.ListQuotesParams, options ...client.Option) *QuotesIter {
	return &QuotesIter{
		Iter: client.GetIter(params.String(), func(url string) (client.ListResponse, []interface{}, error) {
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
