package models

const ListDividendsPath = "/v3/reference/dividends"

// ListDividendsParams is the set of parameters for the ListDividends method.
type ListDividendsParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	ExDividendDateEQ  *string `query:"ex_dividend_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	ExDividendDateLT  *string `query:"ex_dividend_date.lt"`
	ExDividendDateLTE *string `query:"ex_dividend_date.lte"`
	ExDividendDateGT  *string `query:"ex_dividend_date.gt"`
	ExDividendDateGTE *string `query:"ex_dividend_date.gte"`

	RecordDateEQ  *string `query:"record_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	RecordDateLT  *string `query:"record_date.lt"`
	RecordDateLTE *string `query:"record_date.lte"`
	RecordDateGT  *string `query:"record_date.gt"`
	RecordDateGTE *string `query:"record_date.gte"`

	DeclarationDateEQ  *string `query:"declaration_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	DeclarationDateLT  *string `query:"declaration_date.lt"`
	DeclarationDateLTE *string `query:"declaration_date.lte"`
	DeclarationDateGT  *string `query:"declaration_date.gt"`
	DeclarationDateGTE *string `query:"declaration_date.gte"`

	PayDateEQ  *string `query:"pay_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	PayDateLT  *string `query:"pay_date.lt"`
	PayDateLTE *string `query:"pay_date.lte"`
	PayDateGT  *string `query:"pay_date.gt"`
	PayDateGTE *string `query:"pay_date.gte"`

	Frequency *int64 `query:"frequency"` // Enum value

	CashAmountEQ  *float64 `query:"cash_amount"`
	CashAmountLT  *float64 `query:"cash_amount.lt"`
	CashAmountLTE *float64 `query:"cash_amount.lte"`
	CashAmountGT  *float64 `query:"cash_amount.gt"`
	CashAmountGTE *float64 `query:"cash_amount.gte"`

	DividendType *string `query:"dividend_type"` // Enum value

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

// ListDividendsResponse is the response returned by the ListDividends method.
type ListDividendsResponse struct {
	BaseResponse
	Results []Dividend `json:"results,omitempty"`
}

// Dividend contains detailed information on a specified stock dividend.
type Dividend struct {
	CashAmount      float64 `json:"cash_amount,omitempty"`
	DeclarationDate string  `json:"declaration_date,omitempty"`
	ExDividendDate  string  `json:"ex_dividend_date,omitempty"`
	Frequency       int64   `json:"frequency,omitempty"`
	PayDate         string  `json:"pay_date,omitempty"`
	RecordDate      string  `json:"record_date,omitempty"`
	Ticker          string  `json:"ticker,omitempty"`
}
