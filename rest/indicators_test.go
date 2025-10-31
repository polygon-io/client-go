package massive_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
	"github.com/stretchr/testify/assert"
)

const (
	expectedGetSMAURL      = "https://api.massive.com/v1/indicators/sma/AAPL"
	expectedGetSMAResponse = `{
		"status": "OK",
		"request_id": "6a7e466379af0a71039d60cc78e72282",
		"next_url": "https://api.massive.com/v1/indicators/SMA/AAPL?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZ",
		"results": {
			"values":[
				{
					"timestamp": 1578027600,
					"value": 141.34
				},
				{
					"timestamp": 1578035600,
					"value": 139.33
				},
				{
					"timestamp": 1578049600,
					"value": 138.22
				}
			],
			"underlying": {
				"aggregates": [
					{
						"v": 135647456,
						"vw": 74.6099,
						"o": 74.06,
						"c": 75.0875,
						"h": 75.15,
						"l": 73.7975,
						"t": 1577941200000,
						"n": 1
					},
					{
						"v": 146535512,
						"vw": 74.7026,
						"o": 74.2875,
						"c": 74.3575,
						"h": 75.145,
						"l": 74.125,
						"t": 1578027600000,
						"n": 1
					}
				],
				"url": "https://api.massive.com/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?adjusted=true&limit=1&sort=desc"
			}
		}
	}`

	expectedGetSMAResponseNoAggsNoNextURL = `{
		"status": "OK",
		"request_id": "6a7e466379af0a71039d60cc78e72282",
		"results": {
			"values":[
				{
					"timestamp": 1578027600,
					"value": 141.34
				}
			],
			"underlying": {
				"url": "https://api.massive.com/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?adjusted=true&limit=1&sort=desc"
			}
		}
	}`
)

func TestGetSMA(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	registerResponder(expectedGetSMAURL, expectedGetSMAResponse)

	res, err := c.GetSMA(context.Background(), &models.GetSMAParams{
		Ticker: "AAPL",
	})
	assert.Nil(t, err)

	var expect models.GetSMAResponse
	err = json.Unmarshal([]byte(expectedGetSMAResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetSMAWithQueryParams(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	registerResponder("https://api.massive.com/v1/indicators/sma/AAPL?adjusted=false&order=asc&series_type=open&timespan=week&timestamp.gte=1041379200000&timestamp.lte=1658707200000&window=6", expectedGetSMAResponseNoAggsNoNextURL)

	res, err := c.GetSMA(context.Background(), models.GetSMAParams{
		Ticker: "AAPL",
	}.WithOrder(models.Asc).
		WithAdjusted(false).
		WithTimestamp(models.GTE, models.Millis(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC))).
		WithTimestamp(models.LTE, models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))).
		WithTimespan(models.Week).
		WithSeriesType(models.Open).
		WithWindow(6))

	assert.Nil(t, err)

	var expect models.GetSMAResponse
	err = json.Unmarshal([]byte(expectedGetSMAResponseNoAggsNoNextURL), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetEMA(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedEMAResponse := `{
		"status": "OK",
		"request_id": "6a7e466379af0a71039d60cc78e72282",
		"next_url": "https://api.massive.com/v1/indicators/EMA/X:BTC-USD?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZ",
		"results": {
			"values":[
				{
					"timestamp": 1578027600,
					"value": 141.34
				}
			],
			"underlying": {
				"aggregates": [
					{
						"v": 135647456,
						"vw": 74.6099,
						"o": 74.06,
						"c": 75.0875,
						"h": 75.15,
						"l": 73.7975,
						"t": 1577941200000,
						"n": 1
					}
				],
				"url": "https://api.massive.com/v2/aggs/ticker/X:BTC-USD/range/1/day/1626912000000/1629590400000?adjusted=true&limit=1&sort=desc"
			}
		}
	}`

	registerResponder("https://api.massive.com/v1/indicators/ema/X:BTC-USD?order=desc&timestamp.gte=1041379200000&window=10", expectedEMAResponse)

	res, err := c.GetEMA(context.Background(), models.GetEMAParams{
		Ticker: "X:BTC-USD",
	}.WithWindow(10).
		WithOrder(models.Desc).
		WithTimestamp(models.GTE, models.Millis(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC))))

	assert.Nil(t, err)

	var expect models.GetEMAResponse
	err = json.Unmarshal([]byte(expectedEMAResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetRSI(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedRSIResponse := `{
		"status": "OK",
		"request_id": "6a7e466379af0a71039d60cc78e72282",
		"next_url": "https://api.massive.com/v1/indicators/RSI/O:SPY241220P00720000?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZ",
		"results": {
			"underlying": {
				"url": "https://api.massive.com/v2/aggs/ticker/O:SPY241220P00720000/range/1/day/1626912000000/1629590400000?adjusted=true&limit=1&sort=desc"
			}
		}
	}`

	registerResponder("https://api.massive.com/v1/indicators/rsi/O:SPY241220P00720000?order=desc&timestamp.gte=1041379200000&window=10", expectedRSIResponse)

	res, err := c.GetRSI(context.Background(), models.GetRSIParams{
		Ticker: "O:SPY241220P00720000",
	}.WithWindow(10).
		WithOrder(models.Desc).
		WithTimestamp(models.GTE, models.Millis(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC))))

	assert.Nil(t, err)

	var expect models.GetRSIResponse
	err = json.Unmarshal([]byte(expectedRSIResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetMACD(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedMACDResponse := `{
		"status": "OK",
		"request_id": "6a7e466379af0a71039d60cc78e72282",
		"results": {
			"underlying": {
				"url": "https://api.massive.com/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?adjusted=true&limit=1&sort=desc"
			},
			"values":[
				{
					"timestamp": 1578027600,
					"value": 141.34,
					"signal": 143.34,
					"histogram": 145.34
				},
				{
					"timestamp": 1578028400,
					"value": 146.34,
					"signal": 148.34,
					"histogram": 150.34
				}
			]
		}
	}`

	registerResponder("https://api.massive.com/v1/indicators/macd/AAPL?long_window=26&order=desc&short_window=12&signal_window=9&timestamp.lte=1041379200000", expectedMACDResponse)

	res, err := c.GetMACD(context.Background(), models.GetMACDParams{
		Ticker: "AAPL",
	}.WithShortWindow(12).
		WithLongWindow(26).
		WithSignalWindow(9).
		WithOrder(models.Desc).
		WithTimestamp(models.LTE, models.Millis(time.Date(2003, 1, 1, 0, 0, 0, 0, time.UTC))))

	assert.Nil(t, err)

	var expect models.GetMACDResponse
	err = json.Unmarshal([]byte(expectedMACDResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}
