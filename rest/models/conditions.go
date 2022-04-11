package models

// ListConditionsParams is the set of parameters for the ListConditions method.
type ListConditionsParams struct {
	AssetClass *AssetClass `query:"asset_class,omitempty"`
	DataType   *DataType   `query:"data_type,omitempty"`
	ID         *int64      `query:"id,omitempty"`
	SIP        *SIP        `query:"sip,omitempty"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
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

func (p ListConditionsParams) WithSort(q Sort) *ListConditionsParams {
	p.Sort = &q
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

// ListConditionsResponse is the response returned by the ListConditions method.
type ListConditionsResponse struct {
	BaseResponse
	Results []Condition `json:"results,omitempty"`
}

// Condition contains detailed information on a specified condition.
type Condition struct {
	Abbreviation string                     `json:"abbreviation,omitempty"`
	AssetClass   string                     `json:"asset_class,omitempty"`
	DataTypes    []string                   `json:"data_types,omitempty"`
	Description  string                     `json:"description,omitempty"`
	Exchange     int64                      `json:"exchange,omitempty"`
	ID           int64                      `json:"id,omitempty"`
	Legacy       bool                       `json:"legacy"`
	Name         string                     `json:"name,omitempty"`
	SIPMapping   map[string]string          `json:"sip_mapping,omitempty"`
	Type         string                     `json:"type,omitempty"`
	UpdateRules  map[string]map[string]bool `json:"update_rules,omitempty"`
}
