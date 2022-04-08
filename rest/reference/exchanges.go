package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/models"
)

// GetExchanges lists all exchanges that Polygon.io knows about.
func (c *Client) GetExchanges(ctx context.Context, params *models.GetExchangesParams, options ...models.RequestOption) (*models.GetExchangesResponse, error) {
	res := &models.GetExchangesResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetExchangesPath, params, res, options...)
	return res, err
}
