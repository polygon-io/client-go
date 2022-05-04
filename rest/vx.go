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

// VXClient defines a REST client for the Polygon VX API endpoints.
type VXClient struct {
	client.Client
}

// ListStockFinancials gets historical financial data for a stock ticker. This method utilizes an experimental API and could experience breaking changes or deprecation.
func (c *VXClient) ListStockFinancials(ctx context.Context, params *models.ListStockFinancialsParams, options ...models.RequestOption) *iter.Iter[models.StockFinancial] {
	return iter.NewIter(ctx, ListFinancialsPath, params, func(uri string) (iter.ListResponse, []models.StockFinancial, error) {
		res := &models.ListStockFinancialsResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}
