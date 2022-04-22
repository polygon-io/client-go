package polygonws

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSupportsTopic(t *testing.T) {
	assert.Equal(t, true, supportsTopic("stocks", StocksMinAggs))
	assert.Equal(t, false, supportsTopic("stocks", stocksMax))
	assert.Equal(t, true, supportsTopic("options", OptionsSecAggs))
	assert.Equal(t, false, supportsTopic("options", StocksMinAggs))
	assert.Equal(t, true, supportsTopic("forex", ForexQuotes))
	assert.Equal(t, false, supportsTopic("forex", OptionsQuotes))
	assert.Equal(t, true, supportsTopic("crypto", CryptoL2Book))
	assert.Equal(t, false, supportsTopic("crypto", cryptoMin))
	assert.Equal(t, false, supportsTopic("testMarket", StocksImbalances))
}

func TestGetParams(t *testing.T) {
	p1, _ := getParams("stocks", StocksMinAggs, "AAPL", "GME", "HOOD")
	p2, _ := getParams("stocks", StocksMinAggs)
	s1 := "AM.AAPL,AM.GME,AM.HOOD"
	s2 := "AM.*"
	assert.Equal(t, p1, s1)
	assert.Equal(t, p2, s2)
}
func TestSubscriptions(t *testing.T) {
	c, _ := New(context.Background(), Config{
		APIKey:    "test",
		Feed:      RealTime,
		Market:    Stocks,
		ParseData: true,
		Log:       logrus.New(),
	})

	_ = c.Subscribe(StocksMinAggs, "AAPL", "TSLA")
	_, aapl := c.subscriptions["AM"]["AAPL"]
	_, tsla := c.subscriptions["AM"]["TSLA"]
	assert.Equal(t, true, aapl)
	assert.Equal(t, true, tsla)
	_ = c.Unsubscribe(StocksMinAggs, "AAPL", "NFLX")
	_, aapl = c.subscriptions["AM"]["AAPL"]
	assert.Equal(t, false, aapl)
	_ = c.Subscribe(StocksMinAggs)
	_, all := c.subscriptions["AM"]["*"]
	_, tsla = c.subscriptions["AM"]["TSLA"]
	assert.Equal(t, false, tsla)
	assert.Equal(t, true, all)
	_ = c.Unsubscribe(StocksMinAggs, "*")
	_, all = c.subscriptions["AM"]["*"]
	assert.Equal(t, false, all)
	_, trade := c.subscriptions["T"]
	assert.Equal(t, false, trade)
	_ = c.Unsubscribe(StocksTrades, "RDFN")
	_, trade = c.subscriptions["T"]
	assert.Equal(t, true, trade)
	_ = c.Subscribe(StocksTrades, "FB")
	_, fb := c.subscriptions["T"]["FB"]
	assert.Equal(t, true, fb)
	_ = c.Unsubscribe(StocksTrades)
	_, fb = c.subscriptions["T"]["FB"]
	assert.Equal(t, false, fb)
}
