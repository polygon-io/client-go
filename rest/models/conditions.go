package models

// ListConditionsParams is the set of parameters for the ListConditions method.
type ListConditionsParams struct {
	// Filter for conditions within a given asset class.
	AssetClass *AssetClass `query:"asset_class,omitempty"`

	// Filter by data type.
	DataType *DataType `query:"data_type,omitempty"`

	// Filter for conditions with a given ID.
	ID *int64 `query:"id,omitempty"`

	// Filter by SIP. If the condition contains a mapping for that SIP, the condition will be returned.
	SIP *SIP `query:"sip,omitempty"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`
}

func (p ListConditionsParams) WithAssetClass(q AssetClass) *ListConditionsParams {
	p.AssetClass = &q
	return &p
}

func (p ListConditionsParams) WithDataType(q DataType) *ListConditionsParams {
	p.DataType = &q
	return &p
}

func (p ListConditionsParams) WithID(q int64) *ListConditionsParams {
	p.ID = &q
	return &p
}

func (p ListConditionsParams) WithSIP(q SIP) *ListConditionsParams {
	p.SIP = &q
	return &p
}

func (p ListConditionsParams) WithOrder(q Order) *ListConditionsParams {
	p.Order = &q
	return &p
}

func (p ListConditionsParams) WithLimit(q int) *ListConditionsParams {
	p.Limit = &q
	return &p
}

func (p ListConditionsParams) WithSort(q Sort) *ListConditionsParams {
	p.Sort = &q
	return &p
}

// ListConditionsResponse is the response returned by the ListConditions method.
type ListConditionsResponse struct {
	BaseResponse
	Results []Condition `json:"results,omitempty"`
}

// Condition contains detailed information on a specified condition.
type Condition struct {
	Abbreviation string   `json:"abbreviation,omitempty"`
	AssetClass   string   `json:"asset_class,omitempty"`
	DataTypes    []string `json:"data_types,omitempty"`
	Description  string   `json:"description,omitempty"`
	Exchange     int64    `json:"exchange,omitempty"`
	ID           int64    `json:"id,omitempty"`
	Legacy       bool     `json:"legacy"`
	Name         string   `json:"name,omitempty"`
	SIPMapping   struct {
		CTA  string `json:"CTA,omitempty"`
		OPRA string `json:"OPRA,omitempty"`
		UTP  string `json:"UTP,omitempty"`
	} `json:"sip_mapping,omitempty"`
	Type        string `json:"type,omitempty"`
	UpdateRules struct {
		Consolidated struct {
			UpdatesHighLow   bool `json:"updates_high_low,omitempty"`
			UpdatesOpenClose bool `json:"updates_open_close,omitempty"`
			UpdatesVolume    bool `json:"updates_volume,omitempty"`
		} `json:"consolidated,omitempty"`
		MarketCenter struct {
			UpdatesHighLow   bool `json:"updates_high_low,omitempty"`
			UpdatesOpenClose bool `json:"updates_open_close,omitempty"`
			UpdatesVolume    bool `json:"updates_volume,omitempty"`
		} `json:"market_center,omitempty"`
	} `json:"update_rules,omitempty"`
}
