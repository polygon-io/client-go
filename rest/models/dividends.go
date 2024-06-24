package models

// ListDividendsParams is the set of parameters for the ListDividends method.
type ListDividendsParams struct {
	// Return the dividends that contain this ticker.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Query by ex-dividend date with the format YYYY-MM-DD.
	ExDividendDateEQ  *Date `query:"ex_dividend_date"`
	ExDividendDateLT  *Date `query:"ex_dividend_date.lt"`
	ExDividendDateLTE *Date `query:"ex_dividend_date.lte"`
	ExDividendDateGT  *Date `query:"ex_dividend_date.gt"`
	ExDividendDateGTE *Date `query:"ex_dividend_date.gte"`

	// Query by record date with the format YYYY-MM-DD.
	RecordDateEQ  *Date `query:"record_date"`
	RecordDateLT  *Date `query:"record_date.lt"`
	RecordDateLTE *Date `query:"record_date.lte"`
	RecordDateGT  *Date `query:"record_date.gt"`
	RecordDateGTE *Date `query:"record_date.gte"`

	// Query by declaration date with the format YYYY-MM-DD.
	DeclarationDateEQ  *Date `query:"declaration_date"`
	DeclarationDateLT  *Date `query:"declaration_date.lt"`
	DeclarationDateLTE *Date `query:"declaration_date.lte"`
	DeclarationDateGT  *Date `query:"declaration_date.gt"`
	DeclarationDateGTE *Date `query:"declaration_date.gte"`

	// Query by pay date with the format YYYY-MM-DD.
	PayDateEQ  *Date `query:"pay_date"`
	PayDateLT  *Date `query:"pay_date.lt"`
	PayDateLTE *Date `query:"pay_date.lte"`
	PayDateGT  *Date `query:"pay_date.gt"`
	PayDateGTE *Date `query:"pay_date.gte"`

	// Query by the number of times per year the dividend is paid out. Possible values are 0 (one-time), 1 (annually), 2
	// (bi-annually), 4 (quarterly), and 12 (monthly).
	Frequency *Frequency `query:"frequency"`

	// Query by the cash amount of the dividend.
	CashAmountEQ  *float64 `query:"cash_amount"`
	CashAmountLT  *float64 `query:"cash_amount.lt"`
	CashAmountLTE *float64 `query:"cash_amount.lte"`
	CashAmountGT  *float64 `query:"cash_amount.gt"`
	CashAmountGTE *float64 `query:"cash_amount.gte"`

	// Query by the type of dividend. Dividends that have been paid and/or are expected to be paid on consistent
	// schedules are denoted as CD. Special Cash dividends that have been paid that are infrequent or unusual, and/or
	// can not be expected to occur in the future are denoted as SC.
	DividendType *DividendType `query:"dividend_type"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`
}

func (p ListDividendsParams) WithTicker(c Comparator, q string) *ListDividendsParams {
	switch c {
	case EQ:
		p.TickerEQ = &q
	case LT:
		p.TickerLT = &q
	case LTE:
		p.TickerLTE = &q
	case GT:
		p.TickerGT = &q
	case GTE:
		p.TickerGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithExDividendDate(c Comparator, q Date) *ListDividendsParams {
	switch c {
	case EQ:
		p.ExDividendDateEQ = &q
	case LT:
		p.ExDividendDateLT = &q
	case LTE:
		p.ExDividendDateLTE = &q
	case GT:
		p.ExDividendDateGT = &q
	case GTE:
		p.ExDividendDateGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithDeclarationDate(c Comparator, q Date) *ListDividendsParams {
	switch c {
	case EQ:
		p.DeclarationDateEQ = &q
	case LT:
		p.DeclarationDateLT = &q
	case LTE:
		p.DeclarationDateLTE = &q
	case GT:
		p.DeclarationDateGT = &q
	case GTE:
		p.DeclarationDateGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithPayDate(c Comparator, q Date) *ListDividendsParams {
	switch c {
	case EQ:
		p.PayDateEQ = &q
	case LT:
		p.PayDateLT = &q
	case LTE:
		p.PayDateLTE = &q
	case GT:
		p.PayDateGT = &q
	case GTE:
		p.PayDateGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithFrequency(q Frequency) *ListDividendsParams {
	p.Frequency = &q
	return &p
}

func (p ListDividendsParams) WithCashAmount(c Comparator, q float64) *ListDividendsParams {
	switch c {
	case EQ:
		p.CashAmountEQ = &q
	case LT:
		p.CashAmountLT = &q
	case LTE:
		p.CashAmountLTE = &q
	case GT:
		p.CashAmountGT = &q
	case GTE:
		p.CashAmountGTE = &q
	}
	return &p
}

func (p ListDividendsParams) WithDividendType(q DividendType) *ListDividendsParams {
	p.DividendType = &q
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

func (p ListDividendsParams) WithSort(q Sort) *ListDividendsParams {
	p.Sort = &q
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
	DeclarationDate Date    `json:"declaration_date,omitempty"`
	DividendType    string  `json:"dividend_type,omitempty"`
	ExDividendDate  string  `json:"ex_dividend_date,omitempty"`
	Frequency       int64   `json:"frequency,omitempty"`
	PayDate         Date    `json:"pay_date,omitempty"`
	RecordDate      Date    `json:"record_date,omitempty"`
	Ticker          string  `json:"ticker,omitempty"`
}
