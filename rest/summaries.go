package massive

import (
	"context"
	"net/http"

	"github.com/massive-com/client-go/v2/rest/client"
	"github.com/massive-com/client-go/v2/rest/models"
)

const (
	GetSummariesPath = "/v1/summaries"
)

// SummariesClient defines a REST client for the Massive Snapshot Summary API.
type SummariesClient struct {
	client.Client
}

// GetSummaries retrieves summaries for the ticker list with the given params.
// For more details see https://massive.com/docs/stocks/get_v1_summaries.
func (ic *SummariesClient) GetSummaries(ctx context.Context, params *models.GetSummaryParams, opts ...models.RequestOption) (*models.GetSummaryResponse, error) {
	res := &models.GetSummaryResponse{}
	err := ic.Call(ctx, http.MethodGet, GetSummariesPath, params, res, opts...)
	return res, err
}
