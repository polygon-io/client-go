package models

const GetExchangesPath = "/v3/reference/exchanges"

// GetExchangesParams is the set of parameters for the GetExchanges method.
type GetExchangesParams struct {
	AssetClass *AssetClass   `query:"asset_class,omitempty"`
	Locale     *MarketLocale `query:"locale,omitempty"`
}

func (p GetExchangesParams) WithAssetClass(q AssetClass) *GetExchangesParams {
	p.AssetClass = &q
	return &p
}

func (p GetExchangesParams) WithLocale(q MarketLocale) *GetExchangesParams {
	p.Locale = &q
	return &p
}

// GetExchangesResponse is the response returned by the GetExchanges method.
type GetExchangesResponse struct {
	BaseResponse
	Results []Exchange `json:"results,omitempty"`
}

// Exchange contains detailed information on a specified stock Exchange.
type Exchange struct {
	Acronym       string `json:"acronym,omitempty"`
	AssetClass    string `json:"asset_class,omitempty"`
	ID            int64  `json:"id,omitempty"`
	Locale        string `json:"locale,omitempty"`
	MIC           string `json:"mic,omitempty"`
	Name          string `json:"name,omitempty"`
	OperatingMIC  string `json:"operating_mic,omitempty"`
	ParticipantID string `json:"participant_id,omitempty"`
	Type          string `json:"type,omitempty"`
	URL           string `json:"url,omitempty"`
}
