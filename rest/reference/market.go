package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/models"
)

// GetMarketHolidays retrieves upcoming market holidays and their open/close times.
// For more details see https://polygon.io/docs/stocks/get_v1_marketstatus_upcoming.
func (c *Client) GetMarketHolidays(ctx context.Context, params models.GetMarketHolidaysParams, options ...models.RequestOption) (*models.GetMarketHolidaysResponse, error) {
	res := &models.GetMarketHolidaysResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetMarketHolidaysPath, params, res, options...)
	return res, err
}
