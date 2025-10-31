package massive

import (
	"context"
	"net/http"

	"github.com/massive-com/client-go/v2/rest/client"
	"github.com/massive-com/client-go/v2/rest/models"
)

const (
	GetSMAPath  = "/v1/indicators/sma/{ticker}"
	GetEMAPath  = "/v1/indicators/ema/{ticker}"
	GetMACDPath = "/v1/indicators/macd/{ticker}"
	GetRSIPath  = "/v1/indicators/rsi/{ticker}"
)

// IndicatorsClient defines a REST client for the Massive Technical Indicators API.
type IndicatorsClient struct {
	client.Client
}

// GetSMA retrieves simple moving average data over the given time range with the specified parameters.
// For example, if timespan = 'day' and window = '10', a 10-period simple moving average
// will be calculated using day aggregates for each period.
// For more details see https://massive.com/docs/stocks/get_v1_indicators_sma__stockticker.
func (ic *IndicatorsClient) GetSMA(ctx context.Context, params *models.GetSMAParams, opts ...models.RequestOption) (*models.GetSMAResponse, error) {
	res := &models.GetSMAResponse{}
	err := ic.Call(ctx, http.MethodGet, GetSMAPath, params, res, opts...)
	return res, err
}

// GetEMA retrieves exponential moving average data over the given time range with the specified parameters.
// For example, if timespan = 'day' and window = '10', a 10-period exponential moving average
// will be calculated using day aggregates for each period.
// For more details see https://massive.com/docs/stocks/get_v1_indicators_ema__stockticker.
func (ic *IndicatorsClient) GetEMA(ctx context.Context, params *models.GetEMAParams, opts ...models.RequestOption) (*models.GetEMAResponse, error) {
	res := &models.GetEMAResponse{}
	err := ic.Call(ctx, http.MethodGet, GetEMAPath, params, res, opts...)
	return res, err
}

// GetMACD retrieves moving average convergence divergence data over the given time range with the specified parameters.
// For example, if timespan = 'day', short_window = '12', long_window = '26' and signal_window = '9',
// the MACD will be calculated by taking the difference between a 26-period EMA and a 12-period EMA. The signal line values
// will be calculated by taking the 9-day ema of the difference, and the histogram values will be calculated by taking
// the difference between the MACD values and the signal line.
// For more details see https://massive.com/docs/stocks/get_v1_indicators_macd__stockticker.
func (ic *IndicatorsClient) GetMACD(ctx context.Context, params *models.GetMACDParams, opts ...models.RequestOption) (*models.GetMACDResponse, error) {
	res := &models.GetMACDResponse{}
	err := ic.Call(ctx, http.MethodGet, GetMACDPath, params, res, opts...)
	return res, err
}

// GetRSI retrieves relative strength index data over the given time range with the specified parameters.
// For example, if timespan = 'day' and window = '10', a 10-period relative strength index
// will be calculated using day aggregates for each period.
// For more details see https://massive.com/docs/stocks/get_v1_indicators_rsi__stockticker.
func (ic *IndicatorsClient) GetRSI(ctx context.Context, params *models.GetRSIParams, opts ...models.RequestOption) (*models.GetRSIResponse, error) {
	res := &models.GetRSIResponse{}
	err := ic.Call(ctx, http.MethodGet, GetRSIPath, params, res, opts...)
	return res, err
}
