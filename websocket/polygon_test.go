package polygonws

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSupportsTopic(t *testing.T) {
	assert.Equal(t, true, supportsTopic("stocks", Topic(12)))
	assert.Equal(t, false, supportsTopic("stocks", Topic(40)))
	assert.Equal(t, true, supportsTopic("options", Topic(31)))
	assert.Equal(t, false, supportsTopic("options", Topic(40)))
	assert.Equal(t, true, supportsTopic("forex", Topic(52)))
	assert.Equal(t, false, supportsTopic("forex", Topic(40)))
	assert.Equal(t, true, supportsTopic("crypto", Topic(74)))
	assert.Equal(t, false, supportsTopic("crypto", Topic(40)))
	assert.Equal(t, false, supportsTopic("fakeMarket", Topic(15)))
}

func TestGetParams(t *testing.T) {
	p1, _ := getParams("stocks", Topic(12), "AAPL", "GME", "HOOD")
	p2, _ := getParams("stocks", Topic(12))
	s1 := "AM.AAPL,AM.GME,AM.HOOD"
	s2 := "AM.*"
	assert.Equal(t, p1, s1)
	assert.Equal(t, p2, s2)
}
func TestSubscriptions(t *testing.T) {
	c, _ := New(Config{
		APIKey:    "test",
		Feed:      RealTime,
		Market:    Stocks,
		ParseData: true,
		Log:       logrus.New(),
	})

	c.setSubscriptions(Topic(12), "AAPL", "TSLA")
	_, aapl := c.subscriptions["AM"]["AAPL"]
	_, tsla := c.subscriptions["AM"]["TSLA"]
	assert.Equal(t, true, aapl)
	c.deleteSubscriptions(Topic(12), "AAPL", "NFLX")
	_, aapl = c.subscriptions["AM"]["AAPL"]
	assert.Equal(t, false, aapl)
	c.setSubscriptions(Topic(12), "*")
	_, all := c.subscriptions["AM"]["*"]
	_, tsla = c.subscriptions["AM"]["TSLA"]
	assert.Equal(t, false, tsla)
	assert.Equal(t, true, all)
	_, trade := c.subscriptions["T"]
	assert.Equal(t, false, trade)
	c.deleteSubscriptions(Topic(13), "RDFN")
	_, trade = c.subscriptions["T"]
	assert.Equal(t, true, trade)
}
