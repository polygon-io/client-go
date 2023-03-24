package polygonws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testLogger struct{}

func (l *testLogger) Debugf(template string, args ...interface{}) {}
func (l *testLogger) Infof(template string, args ...interface{})  {}
func (l *testLogger) Errorf(template string, args ...interface{}) {}

func TestConfig_Validate(t *testing.T) {
	t.Run("InvalidConfig", func(t *testing.T) {
		cfg := &Config{}
		err := cfg.validate()
		require.Error(t, err)
		assert.Equal(t, "API key is required", err.Error())
	})

	t.Run("ValidConfig", func(t *testing.T) {
		cfg := &Config{
			APIKey: "test-api-key",
		}
		err := cfg.validate()
		require.NoError(t, err)
	})
}

func TestMarket_Supports(t *testing.T) {
	t.Run("ValidSupport", func(t *testing.T) {
		m := Stocks
		topic := StocksTrades
		assert.True(t, m.supports(topic))
	})

	t.Run("InvalidSupport", func(t *testing.T) {
		m := Stocks
		topic := OptionsTrades
		assert.False(t, m.supports(topic))
	})
}

func TestTopicPrefix(t *testing.T) {
	testCases := []struct {
			name   string
			topic  Topic
			prefix string
	}{
			{"StocksSecAggs", StocksSecAggs, "A"},
			{"StocksMinAggs", StocksMinAggs, "AM"},
			{"StocksTrades", StocksTrades, "T"},
			{"StocksQuotes", StocksQuotes, "Q"},
			{"StocksImbalances", StocksImbalances, "NOI"},
			{"StocksLULD", StocksLULD, "LULD"},
			{"OptionsSecAggs", OptionsSecAggs, "A"},
			{"OptionsMinAggs", OptionsMinAggs, "AM"},
			{"OptionsTrades", OptionsTrades, "T"},
			{"OptionsQuotes", OptionsQuotes, "Q"},
			{"ForexMinAggs", ForexMinAggs, "CA"},
			{"ForexQuotes", ForexQuotes, "C"},
			{"CryptoMinAggs", CryptoMinAggs, "XA"},
			{"CryptoTrades", CryptoTrades, "XT"},
			{"CryptoQuotes", CryptoQuotes, "XQ"},
			{"CryptoL2Book", CryptoL2Book, "XL2"},
	}

	for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
					assert.Equal(t, tc.prefix, tc.topic.prefix())
			})
	}
}