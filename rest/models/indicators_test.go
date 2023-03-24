package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

func TestGetSMAParams(t *testing.T) {
	sma := models.GetSMAParams{}
	adjusted := true
	order := models.Asc
	expandUnderlying := true
	ts := models.Millis(time.Now())
	allComparators := [5]models.Comparator{models.EQ, models.LT, models.LTE, models.GT, models.GTE}

	sma = *sma.WithAdjusted(adjusted)
	sma = *sma.WithOrder(order)
	sma = *sma.WithExpandUnderlying(expandUnderlying)

	if *sma.Adjusted != adjusted {
		t.Errorf("Adjusted not set correctly, got: %v, want: %v", *sma.Adjusted, adjusted)
	}

	if *sma.Order != order {
		t.Errorf("Order not set correctly, got: %v, want: %v", *sma.Order, order)
	}

	if *sma.ExpandUnderlying != expandUnderlying {
		t.Errorf("ExpandUnderlying not set correctly, got: %v, want: %v", *sma.ExpandUnderlying, expandUnderlying)
	}

	for _, c := range allComparators {
		sma = *sma.WithTimestamp(c, ts)
		switch c {
		case models.EQ:
			if *sma.TimestampEQ != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *sma.TimestampEQ, ts)
			}
		case models.LT:
			if *sma.TimestampLT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *sma.TimestampLT, ts)
			}
		case models.LTE:
			if *sma.TimestampLTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *sma.TimestampLTE, ts)
			}
		case models.GT:
			if *sma.TimestampGT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *sma.TimestampGT, ts)
			}
		case models.GTE:
			if *sma.TimestampGTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *sma.TimestampGTE, ts)
			}
		}
	}
}

func TestGetEMAParams(t *testing.T) {
	ema := models.GetEMAParams{}
	adjusted := true
	order := models.Asc
	expandUnderlying := true
	ts := models.Millis(time.Now())
	allComparators := [5]models.Comparator{models.EQ, models.LT, models.LTE, models.GT, models.GTE}

	ema = *ema.WithAdjusted(adjusted)
	ema = *ema.WithOrder(order)
	ema = *ema.WithExpandUnderlying(expandUnderlying)

	if *ema.Adjusted != adjusted {
		t.Errorf("Adjusted not set correctly, got: %v, want: %v", *ema.Adjusted, adjusted)
	}

	if *ema.Order != order {
		t.Errorf("Order not set correctly, got: %v, want: %v", *ema.Order, order)
	}

	if *ema.ExpandUnderlying != expandUnderlying {
		t.Errorf("ExpandUnderlying not set correctly, got: %v, want: %v", *ema.ExpandUnderlying, expandUnderlying)
	}

	for _, c := range allComparators {
		ema = *ema.WithTimestamp(c, ts)
		switch c {
		case models.EQ:
			if *ema.TimestampEQ != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *ema.TimestampEQ, ts)
			}
		case models.LT:
			if *ema.TimestampLT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *ema.TimestampLT, ts)
			}
		case models.LTE:
			if *ema.TimestampLTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *ema.TimestampLTE, ts)
			}
		case models.GT:
			if *ema.TimestampGT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *ema.TimestampGT, ts)
			}
		case models.GTE:
			if *ema.TimestampGTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *ema.TimestampGTE, ts)
			}
		}
	}
}

func TestGetRSIParams(t *testing.T) {
	rsi := models.GetRSIParams{}
	adjusted := true
	order := models.Asc
	expandUnderlying := true
	ts := models.Millis(time.Now())
	allComparators := [5]models.Comparator{models.EQ, models.LT, models.LTE, models.GT, models.GTE}

	rsi = *rsi.WithAdjusted(adjusted)
	rsi = *rsi.WithOrder(order)
	rsi = *rsi.WithExpandUnderlying(expandUnderlying)

	if *rsi.Adjusted != adjusted {
		t.Errorf("Adjusted not set correctly, got: %v, want: %v", *rsi.Adjusted, adjusted)
	}

	if *rsi.Order != order {
		t.Errorf("Order not set correctly, got: %v, want: %v", *rsi.Order, order)
	}

	if *rsi.ExpandUnderlying != expandUnderlying {
		t.Errorf("ExpandUnderlying not set correctly, got: %v, want: %v", *rsi.ExpandUnderlying, expandUnderlying)
	}

	for _, c := range allComparators {
		rsi = *rsi.WithTimestamp(c, ts)
		switch c {
		case models.EQ:
			if *rsi.TimestampEQ != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *rsi.TimestampEQ, ts)
			}
		case models.LT:
			if *rsi.TimestampLT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *rsi.TimestampLT, ts)
			}
		case models.LTE:
			if *rsi.TimestampLTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *rsi.TimestampLTE, ts)
			}
		case models.GT:
			if *rsi.TimestampGT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *rsi.TimestampGT, ts)
			}
		case models.GTE:
			if *rsi.TimestampGTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *rsi.TimestampGTE, ts)
			}
		}
	}
}

func TestGetMACDParams(t *testing.T) {
	macd := models.GetMACDParams{}
	adjusted := true
	order := models.Asc
	expandUnderlying := true
	ts := models.Millis(time.Now())
	allComparators := [5]models.Comparator{models.EQ, models.LT, models.LTE, models.GT, models.GTE}

	macd = *macd.WithAdjusted(adjusted)
	macd = *macd.WithOrder(order)
	macd = *macd.WithExpandUnderlying(expandUnderlying)

	if *macd.Adjusted != adjusted {
		t.Errorf("Adjusted not set correctly, got: %v, want: %v", *macd.Adjusted, adjusted)
	}

	if *macd.Order != order {
		t.Errorf("Order not set correctly, got: %v, want: %v", *macd.Order, order)
	}

	if *macd.ExpandUnderlying != expandUnderlying {
		t.Errorf("ExpandUnderlying not set correctly, got: %v, want: %v", *macd.ExpandUnderlying, expandUnderlying)
	}

	for _, c := range allComparators {
		macd = *macd.WithTimestamp(c, ts)
		switch c {
		case models.EQ:
			if *macd.TimestampEQ != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *macd.TimestampEQ, ts)
			}
		case models.LT:
			if *macd.TimestampLT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *macd.TimestampLT, ts)
			}
		case models.LTE:
			if *macd.TimestampLTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *macd.TimestampLTE, ts)
			}
		case models.GT:
			if *macd.TimestampGT != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *macd.TimestampGT, ts)
			}
		case models.GTE:
			if *macd.TimestampGTE != ts {
				t.Errorf("Timestamp not set correctly, got: %v, want: %v", *macd.TimestampGTE, ts)
			}
		}
	}
}
