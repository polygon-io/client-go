package models

// ListDividendsParams is the set of parameters for the ListDividends method.
type ListDividendsParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	ExDividendDateEQ  *Date `query:"ex_dividend_date"`
	ExDividendDateLT  *Date `query:"ex_dividend_date.lt"`
	ExDividendDateLTE *Date `query:"ex_dividend_date.lte"`
	ExDividendDateGT  *Date `query:"ex_dividend_date.gt"`
	ExDividendDateGTE *Date `query:"ex_dividend_date.gte"`

	RecordDateEQ  *Date `query:"record_date"`
	RecordDateLT  *Date `query:"record_date.lt"`
	RecordDateLTE *Date `query:"record_date.lte"`
	RecordDateGT  *Date `query:"record_date.gt"`
	RecordDateGTE *Date `query:"record_date.gte"`

	DeclarationDateEQ  *Date `query:"declaration_date"`
	DeclarationDateLT  *Date `query:"declaration_date.lt"`
	DeclarationDateLTE *Date `query:"declaration_date.lte"`
	DeclarationDateGT  *Date `query:"declaration_date.gt"`
	DeclarationDateGTE *Date `query:"declaration_date.gte"`

	PayDateEQ  *Date `query:"pay_date"`
	PayDateLT  *Date `query:"pay_date.lt"`
	PayDateLTE *Date `query:"pay_date.lte"`
	PayDateGT  *Date `query:"pay_date.gt"`
	PayDateGTE *Date `query:"pay_date.gte"`

	Frequency *Frequency `query:"frequency"`

	CashAmountEQ  *float64 `query:"cash_amount"`
	CashAmountLT  *float64 `query:"cash_amount.lt"`
	CashAmountLTE *float64 `query:"cash_amount.lte"`
	CashAmountGT  *float64 `query:"cash_amount.gt"`
	CashAmountGTE *float64 `query:"cash_amount.gte"`

	DividendType *DividendType `query:"dividend_type"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListDividendsParams) WithTicker(c Comparator, q string) *ListDividendsParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithExDividendDate(c Comparator, q Date) *ListDividendsParams {
	if c == EQ {
		p.ExDividendDateEQ = &q
	} else if c == LT {
		p.ExDividendDateLT = &q
	} else if c == LTE {
		p.ExDividendDateLTE = &q
	} else if c == GT {
		p.ExDividendDateGT = &q
	} else if c == GTE {
		p.ExDividendDateGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithDeclarationDate(c Comparator, q Date) *ListDividendsParams {
	if c == EQ {
		p.DeclarationDateEQ = &q
	} else if c == LT {
		p.DeclarationDateLT = &q
	} else if c == LTE {
		p.DeclarationDateLTE = &q
	} else if c == GT {
		p.DeclarationDateGT = &q
	} else if c == GTE {
		p.DeclarationDateGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithFrequency(q Frequency) *ListDividendsParams {
	p.Frequency = &q
	return &p
}

func (p ListDividendsParams) WithCashAmount(c Comparator, q float64) *ListDividendsParams {
	if c == EQ {
		p.CashAmountEQ = &q
	} else if c == LT {
		p.CashAmountLT = &q
	} else if c == LTE {
		p.CashAmountLTE = &q
	} else if c == GT {
		p.CashAmountGT = &q
	} else if c == GTE {
		p.CashAmountGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithDividendType(q DividendType) *ListDividendsParams {
	p.DividendType = &q
	return &p
}

func (p ListDividendsParams) WithSort(q Sort) *ListDividendsParams {
	p.Sort = &q
	return &p
}

func (p ListDividendsParams) WithOrder(q Order) *ListDividendsParams {
	p.Order = &q
	return &p
}

func (p ListDividendsParams) WithLimit(q int) *ListDividendsParams {
	p.Limit = &q
	return &p
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
	DividendType    string  `json:"dividend_type,omitempty"`
	ExDividendDate  string  `json:"ex_dividend_date,omitempty"`
	Frequency       int64   `json:"frequency,omitempty"`
	PayDate         string  `json:"pay_date,omitempty"`
	RecordDate      string  `json:"record_date,omitempty"`
	Ticker          string  `json:"ticker,omitempty"`
}
