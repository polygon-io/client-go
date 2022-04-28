package models

// ListTickersParams is the set of parameters for the ListTickers method.
type ListTickersParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	Type     *string     `query:"type"`
	Market   *AssetClass `query:"market"`
	Exchange *int        `query:"exchange"`
	CUSIP    *int        `query:"cusip"`
	CIK      *int        `query:"cik"`
	Date     *Date       `query:"date"`
	Active   *bool       `query:"active"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`

	PageMarker *string `query:"page_marker"`
	Search     *string `query:"search"`
}

func (p ListTickersParams) WithTicker(c Comparator, q string) *ListTickersParams {
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

func (p ListTickersParams) WithType(q string) *ListTickersParams {
	p.Type = &q
	return &p
}

func (p ListTickersParams) WithMarket(q AssetClass) *ListTickersParams {
	p.Market = &q
	return &p
}

func (p ListTickersParams) WithExchange(q int) *ListTickersParams {
	p.Exchange = &q
	return &p
}

func (p ListTickersParams) WithCUSIP(q int) *ListTickersParams {
	p.CUSIP = &q
	return &p
}

func (p ListTickersParams) WithCIK(q int) *ListTickersParams {
	p.CIK = &q
	return &p
}

func (p ListTickersParams) WithDate(q Date) *ListTickersParams {
	p.Date = &q
	return &p
}

func (p ListTickersParams) WithActive(q bool) *ListTickersParams {
	p.Active = &q
	return &p
}

func (p ListTickersParams) WithSort(q Sort) *ListTickersParams {
	p.Sort = &q
	return &p
}

func (p ListTickersParams) WithOrder(q Order) *ListTickersParams {
	p.Order = &q
	return &p
}

func (p ListTickersParams) WithLimit(q int) *ListTickersParams {
	p.Limit = &q
	return &p
}

func (p ListTickersParams) WithPageMarker(q string) *ListTickersParams {
	p.PageMarker = &q
	return &p
}

func (p ListTickersParams) WithSearch(q string) *ListTickersParams {
	p.Search = &q
	return &p
}

// ListTickersResponse is the response returned by the ListTickers method.
type ListTickersResponse struct {
	BaseResponse
	Results []Ticker `json:"results,omitempty"`
}

// GetTickerDetailsParams is the set of parameters for the GetTickerDetails method.
type GetTickerDetailsParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Date *Date `query:"date"`
}

func (p GetTickerDetailsParams) WithDate(q Date) *GetTickerDetailsParams {
	p.Date = &q
	return &p
}

// GetTickerDetailsResponse is the response returned by the GetTickerDetails method.
type GetTickerDetailsResponse struct {
	BaseResponse
	Results Ticker `json:"results,omitempty"`
}

// ListTickerNewsParams is the set of parameters for the ListTickerNews method.
type ListTickerNewsParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	PublishedUtcEQ  *Millis `query:"published_utc"`
	PublishedUtcLT  *Millis `query:"published_utc.lt"`
	PublishedUtcLTE *Millis `query:"published_utc.lte"`
	PublishedUtcGT  *Millis `query:"published_utc.gt"`
	PublishedUtcGTE *Millis `query:"published_utc.gte"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListTickerNewsParams) WithTicker(c Comparator, q string) *ListTickerNewsParams {
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

func (p ListTickerNewsParams) WithPublishedUTC(c Comparator, q Millis) *ListTickerNewsParams {
	if c == EQ {
		p.PublishedUtcEQ = &q
	} else if c == LT {
		p.PublishedUtcLT = &q
	} else if c == LTE {
		p.PublishedUtcLTE = &q
	} else if c == GT {
		p.PublishedUtcGT = &q
	} else if c == GTE {
		p.PublishedUtcGTE = &q
	}
	return &p
}

func (p ListTickerNewsParams) WithSort(q Sort) *ListTickerNewsParams {
	p.Sort = &q
	return &p
}

func (p ListTickerNewsParams) WithOrder(q Order) *ListTickerNewsParams {
	p.Order = &q
	return &p
}

func (p ListTickerNewsParams) WithLimit(q int) *ListTickerNewsParams {
	p.Limit = &q
	return &p
}

// ListTickerNewsResponse is the response returned by the ListTickerNews method.
type ListTickerNewsResponse struct {
	BaseResponse
	Results []TickerNews `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of parameters for the GetTickerTypes method.
type GetTickerTypesParams struct {
	AssetClass *AssetClass   `query:"asset_class"`
	Locale     *MarketLocale `query:"locale"`
}

func (p GetTickerTypesParams) WithAssetClass(q AssetClass) *GetTickerTypesParams {
	p.AssetClass = &q
	return &p
}

func (p GetTickerTypesParams) WithLocale(q MarketLocale) *GetTickerTypesParams {
	p.Locale = &q
	return &p
}

// GetTickerTypesResponse is the response returned by the GetTickerTypes method.
type GetTickerTypesResponse struct {
	BaseResponse
	Results []TickerType `json:"results,omitempty"`
}

// Ticker contains detailed information on a specified ticker symbol.
type Ticker struct {
	Ticker                      string          `json:"ticker,omitempty"`
	Name                        string          `json:"name,omitempty"`
	Market                      string          `json:"market,omitempty"`
	Locale                      string          `json:"locale,omitempty"`
	PrimaryExchange             string          `json:"primary_exchange,omitempty"`
	Type                        string          `json:"type,omitempty"`
	Active                      bool            `json:"active"`
	CurrencySymbol              string          `json:"currency_symbol,omitempty"`
	CurrencyName                string          `json:"currency_name,omitempty"`
	BaseCurrencySymbol          string          `json:"base_currency_symbol,omitempty"`
	BaseCurrencyName            string          `json:"base_currency_name,omitempty"`
	CUSIP                       string          `json:"cusip,omitempty"`
	CIK                         string          `json:"cik,omitempty"`
	CompositeFIGI               string          `json:"composite_figi,omitempty"`
	ShareClassFIGI              string          `json:"share_class_figi,omitempty"`
	LastUpdatedUTC              Time            `json:"last_updated_utc,omitempty"`
	DelistedUTC                 Time            `json:"delisted_utc,omitempty"`
	MarketCap                   float64         `json:"market_cap,omitempty"`
	PhoneNumber                 string          `json:"phone_number,omitempty"`
	Address                     *CompanyAddress `json:"address,omitempty"`
	Description                 string          `json:"description,omitempty"`
	SICCode                     string          `json:"sic_code,omitempty"`
	SICDescription              string          `json:"sic_description,omitempty"`
	TickerRoot                  string          `json:"ticker_root,omitempty"`
	TickerSuffix                string          `json:"ticker_suffix,omitempty"`
	HomepageURL                 string          `json:"homepage_url,omitempty"`
	TotalEmployees              int32           `json:"total_employees,omitempty"`
	ListDate                    string          `json:"list_date,omitempty"`
	Branding                    Branding        `json:"branding,omitempty"`
	ShareClassSharesOutstanding int64           `json:"share_class_shares_outstanding,omitempty"`
	WeightedSharesOutstanding   int64           `json:"weighted_shares_outstanding,omitempty"`
}

// CompanyAddress contains information on the physical address of a company.
type CompanyAddress struct {
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// Branding contains information related to a company's brand.
type Branding struct {
	LogoURL     string `json:"logo_url,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	AccentColor string `json:"accent_color,omitempty"`
	LightColor  string `json:"light_color,omitempty"`
	DarkColor   string `json:"dark_color,omitempty"`
}

type TickerNews struct {
	ID           string    `json:"id,omitempty"`
	Publisher    Publisher `json:"publisher,omitempty"`
	Title        string    `json:"title,omitempty"`
	Author       string    `json:"author,omitempty"`
	PublishedUTC Time      `json:"published_utc,omitempty"`
	ArticleURL   string    `json:"article_url,omitempty"`
	Tickers      []string  `json:"tickers,omitempty"`
	AMPURL       string    `json:"amp_url,omitempty"`
	ImageURL     string    `json:"image_url,omitempty"`
	Description  string    `json:"description,omitempty"`
	Keywords     []string  `json:"keywords,omitempty"`
}

type Publisher struct {
	Name        string `json:"name,omitempty"`
	HomepageURL string `json:"homepage_url,omitempty"`
	LogoURL     string `json:"logo_url,omitempty"`
	FaviconURL  string `json:"favicon_url,omitempty"`
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Locale      string `json:"locale,omitempty"`
}
