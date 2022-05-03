package models

// ListFinancialsParams is the set of parameters for the ListFinancials method.
type ListFinancialsParams struct {
	Ticker *string `query:"ticker"`

	CIK *string `query:"cik"`

	CompanyNameFULL   *string `query:"company_name"`
	CompanyNameSEARCH *string `query:"company_name.search"`

	SIC *string `query:"sic"`

	FilingDateEQ  *Date `query:"filing_dividend_date"`
	FilingDateLT  *Date `query:"filing_dividend_date.lt"`
	FilingDateLTE *Date `query:"filing_dividend_date.lte"`
	FilingDateGT  *Date `query:"filing_dividend_date.gt"`
	FilingDateGTE *Date `query:"filing_dividend_date.gte"`

	PeriodOfReportDateEQ  *Date `query:"period_of_report_date"`
	PeriodOfReportDateLT  *Date `query:"period_of_report_date.lt"`
	PeriodOfReportDateLTE *Date `query:"period_of_report_date.lte"`
	PeriodOfReportDateGT  *Date `query:"period_of_report_date.gt"`
	PeriodOfReportDateGTE *Date `query:"period_of_report_date.gte"`

	Timeframe *Timeframe `query:"timeframe"`

	IncludeSources *bool `query:"include_sources"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListFinancialsParams) WithTicker(q string) *ListFinancialsParams {
	p.Ticker = &q
	return &p
}

func (p ListFinancialsParams) WithCIK(q string) *ListFinancialsParams {
	p.CIK = &q
	return &p
}

func (p ListFinancialsParams) WithCompanyName(c NameComparator, q string) *ListFinancialsParams {
	if c == FULL {
		p.CompanyNameFULL = &q
	} else if c == SEARCH {
		p.CompanyNameSEARCH = &q
	}
	return &p
}

func (p ListFinancialsParams) WithSIC(q string) *ListFinancialsParams {
	p.SIC = &q
	return &p
}

func (p ListFinancialsParams) WithFilingDate(c Comparator, q Date) *ListFinancialsParams {
	if c == EQ {
		p.FilingDateEQ = &q
	} else if c == LT {
		p.FilingDateLT = &q
	} else if c == LTE {
		p.FilingDateLTE = &q
	} else if c == GT {
		p.FilingDateGT = &q
	} else if c == GTE {
		p.FilingDateGTE = &q
	}
	return &p
}

func (p ListFinancialsParams) WithPeriodOfReportDate(c Comparator, q Date) *ListFinancialsParams {
	if c == EQ {
		p.PeriodOfReportDateEQ = &q
	} else if c == LT {
		p.PeriodOfReportDateLT = &q
	} else if c == LTE {
		p.PeriodOfReportDateLTE = &q
	} else if c == GT {
		p.PeriodOfReportDateGT = &q
	} else if c == GTE {
		p.PeriodOfReportDateGTE = &q
	}
	return &p
}

func (p ListFinancialsParams) WithTimeframe(q Timeframe) *ListFinancialsParams {
	p.Timeframe = &q
	return &p
}

func (p ListFinancialsParams) WithIncludeSources(q bool) *ListFinancialsParams {
	p.IncludeSources = &q
	return &p
}

func (p ListFinancialsParams) WithSort(q Sort) *ListFinancialsParams {
	p.Sort = &q
	return &p
}

func (p ListFinancialsParams) WithOrder(q Order) *ListFinancialsParams {
	p.Order = &q
	return &p
}

func (p ListFinancialsParams) WithLimit(q int) *ListFinancialsParams {
	p.Limit = &q
	return &p
}

// ListFinancialsResponse is the response returned by the ListFinancials method.
type ListFinancialsResponse struct {
	BaseResponse
	Results []Financial `json:"results,omitempty"`
}

// Financial contains detailed information on a specified stock financial.
type Financial struct {
	CIK                 string                 `json:"cik,omitempty"`
	CompanyName         string                 `json:"company_name,omitempty"`
	EndDate             string                 `json:"end_date,omitempty"`
	FilingDate          string                 `json:"filing_date,omitempty"`
	Financials          map[string]interface{} `json:"financials,omitempty"`
	FiscalPeriod        string                 `json:"fiscal_period,omitempty"`
	FiscalYear          string                 `json:"fiscal_year,omitempty"`
	SourceFilingFileUrl string                 `json:"source_filing_file_url,omitempty"`
	SourceFilingUrl     string                 `json:"source_filing_url,omitempty"`
	StartDate           string                 `json:"start_date,omitempty"`
}
