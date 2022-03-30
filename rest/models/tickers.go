package models

import (
	"net/url"
	"strconv"
	"time"

	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListTickersPath      = "/v3/reference/tickers"
	GetTickerDetailsPath = "/v3/reference/tickers/{ticker}"
	GetTickerTypesPath   = "/v3/reference/tickers/types"
)

// TickerDetails contains detailed information on a specified ticker symbol.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
type TickerDetails struct {
	Ticker                      string         `json:"ticker"`
	Name                        string         `json:"name"`
	Market                      string         `json:"market"`
	Locale                      string         `json:"locale"`
	PrimaryExchange             string         `json:"primary_exchange,omitempty"`
	Type                        string         `json:"type,omitempty"`
	Active                      bool           `json:"active"`
	CurrencySymbol              string         `json:"currency_symbol,omitempty"`
	CurrencyName                string         `json:"currency_name,omitempty"`
	BaseCurrencySymbol          string         `json:"base_currency_symbol,omitempty"`
	BaseCurrencyName            string         `json:"base_currency_name,omitempty"`
	CUSIP                       string         `json:"cusip,omitempty"`
	CIK                         string         `json:"cik,omitempty"`
	CompositeFIGI               string         `json:"composite_figi,omitempty"`
	ShareClassFIGI              string         `json:"share_class_figi,omitempty"`
	LastUpdatedUTC              int64          `json:"last_updated_utc,omitempty"`
	DelistedUTC                 int64          `json:"delisted_utc,omitempty"`
	MarketCap                   float64        `json:"market_cap,omitempty"`
	PhoneNumber                 string         `json:"phone_number,omitempty"`
	Address                     CompanyAddress `json:"address,omitempty"` // todo: ptr?
	Description                 string         `json:"description,omitempty"`
	SICCode                     string         `json:"sic_code,omitempty"`
	SICDescription              string         `json:"sic_description,omitempty"`
	TickerRoot                  string         `json:"ticker_root,omitempty"`
	TickerSuffix                string         `json:"ticker_suffix,omitempty"`
	HomepageURL                 string         `json:"homepage_url,omitempty"`
	TotalEmployees              int32          `json:"total_employees,omitempty"`
	ListDate                    string         `json:"list_date,omitempty"`
	Branding                    Branding       `json:"branding,omitempty"` // todo: ptr?
	ShareClassSharesOutstanding int64          `json:"share_class_shares_outstanding,omitempty"`
	WeightedSharesOutstanding   int64          `json:"weighted_shares_outstanding,omitempty"`
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

// TickersResponse contains a list of reference tickers.
type TickersResponse struct {
	client.BaseResponse
	Results []*TickerDetails `json:"results,omitempty"`
}

// ListTickersParams is the set of path and query parameters that are used to request reference tickers.
type ListTickersParams struct {
	QueryParams ListTickersQueryParams
}

// ListTickersQueryParams is the set of query parameters for requesting reference tickers.
type ListTickersQueryParams struct {
	TickerEQ  *string
	TickerLT  *string
	TickerLTE *string
	TickerGT  *string
	TickerGTE *string

	// todo: which ones should be enums?
	Type     *string
	Market   *MarketType // todo: this endpoint apparently expects fx instead of forex
	Exchange *string
	CUSIP    *string
	CIK      *string
	Date     *time.Time
	Active   *bool

	Sort  *Sort
	Order *Order
	Limit *int

	PageMarker *string
	Search     *string

	Cursor *string
}

// Path maps the path parameters to their respective keys.
func (p ListTickersParams) Path() map[string]string {
	return map[string]string{}
}

// Query maps the query parameters to their respective keys.
func (p ListTickersParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.TickerEQ != nil {
		q.Set("ticker", *p.QueryParams.TickerEQ)
	}

	if p.QueryParams.TickerLT != nil {
		q.Set("ticker.lt", *p.QueryParams.TickerLT)
	}

	if p.QueryParams.TickerLTE != nil {
		q.Set("ticker.lte", *p.QueryParams.TickerLTE)
	}

	if p.QueryParams.TickerGT != nil {
		q.Set("ticker.gt", *p.QueryParams.TickerGT)
	}

	if p.QueryParams.TickerGTE != nil {
		q.Set("ticker.gte", *p.QueryParams.TickerGTE)
	}

	if p.QueryParams.Type != nil {
		q.Set("type", *p.QueryParams.Type)
	}

	if p.QueryParams.Market != nil {
		q.Set("market", string(*p.QueryParams.Market))
	}

	if p.QueryParams.Exchange != nil {
		q.Set("exchange", *p.QueryParams.Exchange)
	}

	if p.QueryParams.CUSIP != nil {
		q.Set("cusip", *p.QueryParams.CUSIP)
	}

	if p.QueryParams.CIK != nil {
		q.Set("cik", *p.QueryParams.CIK)
	}

	if p.QueryParams.Date != nil {
		q.Set("date", p.QueryParams.Date.Format("2006-01-02"))
	}

	if p.QueryParams.Active != nil {
		q.Set("active", strconv.FormatBool(*p.QueryParams.Active))
	} else {
		q.Set("active", "true")
	}

	if p.QueryParams.Sort != nil {
		q.Set("sort", string(*p.QueryParams.Sort))
	}

	if p.QueryParams.Order != nil {
		q.Set("order", string(*p.QueryParams.Order))
	}

	if p.QueryParams.Limit != nil {
		q.Set("limit", strconv.FormatInt(int64(*p.QueryParams.Limit), 10))
	}

	if p.QueryParams.PageMarker != nil {
		q.Set("page_marker", *p.QueryParams.PageMarker)
	}

	if p.QueryParams.Search != nil {
		q.Set("search", *p.QueryParams.Search)
	}

	if p.QueryParams.Cursor != nil {
		q.Set("cursor", *p.QueryParams.Cursor)
	}

	return q
}

// String returns a URL string that includes any path and query parameters that are set.
func (p ListTickersParams) String() string {
	path := ListTickersPath

	q := p.Query().Encode()
	if q != "" {
		path += "?" + q
	}

	return path
}

// GetTickerDetailsParams is the set of path and query parameters that are used to request reference ticker details.
type GetTickerDetailsParams struct {
	Ticker      string
	QueryParams GetTickerDetailsQueryParams
}

// GetTickerDetailsQueryParams is the set of query parameters for requesting reference ticker details.
type GetTickerDetailsQueryParams struct {
	Date *time.Time
}

// Path maps the path parameters to their respective keys.
func (p GetTickerDetailsParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// Query maps the query parameters to their respective keys.
func (p GetTickerDetailsParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Date != nil {
		q.Set("date", p.QueryParams.Date.Format("2006-01-02"))
	}

	return q
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
}

// TickerTypesResponse contains a list of ticker types.
type TickerTypesResponse struct {
	client.BaseResponse
	Results []*TickerType `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of path and query parameters that are used to request ticker types.
type GetTickerTypesParams struct {
	QueryParams GetTickerTypesQueryParams
}

// GetTickerTypesQueryParams is the set of query parameters for requesting ticker types.
type GetTickerTypesQueryParams struct {
	AssetClass *string // todo: this is similar but slightly different than market type (also we offer four options but only one returns results)

	Locale *MarketLocale

	AfterPrimary   *string // todo: these aren't typically documented, what is it for?
	AfterSecondary *string

	Cursor *string
}

// Path maps the path parameters to their respective keys.
func (p GetTickerTypesParams) Path() map[string]string {
	return map[string]string{}
}

// Query maps the query parameters to their respective keys.
func (p GetTickerTypesParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.AssetClass != nil {
		q.Set("asset_class", *p.QueryParams.AssetClass)
	}

	if p.QueryParams.Locale != nil {
		q.Set("locale", string(*p.QueryParams.Locale))
	}

	if p.QueryParams.AfterPrimary != nil {
		q.Set("ap", *p.QueryParams.AfterPrimary)
	}

	if p.QueryParams.AfterSecondary != nil {
		q.Set("as", *p.QueryParams.AfterSecondary)
	}

	if p.QueryParams.Cursor != nil {
		q.Set("cursor", *p.QueryParams.Cursor)
	}

	return q
}
