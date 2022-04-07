package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/models"
)

// GetMarketHolidays retrieves upcoming market holidays and their open/close times.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_upcoming.
func (c *Client) GetMarketHolidays(ctx context.Context) (*models.GetMarketHolidaysResponse, error) {
	res := &models.GetMarketHolidaysResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetMarketHolidaysPath, nil, res)
	return res, err
}

// GetMarketStatus retrieves the current trading status of the exchanges and overall financial markets.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_now.
func (c *Client) GetMarketStatus(ctx context.Context) (*models.GetMarketStatusResponse, error) {
	res := &models.GetMarketStatusResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetMarketStatusPath, nil, res)
	return res, err
}
