package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListFinancialsPath = "/vX/reference/financials"
)

// VXClient defines a REST client for the Polygon VX (experimental) API.
type VXClient struct {
	client.Client
}

// Get historical financial data for a stock ticker. The financials data is extracted from XBRL from company SEC filings
// using the methodology outlined here: http://xbrl.squarespace.com/understanding-sec-xbrl-financi/.
//
// Note: this method utilizes an experimental API and could experience breaking changes or deprecation.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListStockFinancials(context.TODO(), params, opts...)
//	for iter.Next() {
//	    log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//	    return iter.Err()
//	}
func (c *VXClient) ListStockFinancials(ctx context.Context, params *models.ListStockFinancialsParams, options ...models.RequestOption) *iter.Iter[models.StockFinancial] {
	return iter.NewIter(ctx, ListFinancialsPath, params, func(uri string) (iter.ListResponse, []models.StockFinancial, error) {
		res := &models.ListStockFinancialsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
