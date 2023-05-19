package models

// Deprecated: Please use UniversalSnapshot types instead of AssetSnapshot types.
type ListAssetSnapshotsParams ListUniversalSnapshotsParams

func (p ListAssetSnapshotsParams) WithTickerAnyOf(q string) *ListAssetSnapshotsParams {
	p.TickerAnyOf = &q
	return &p
}

func (p ListAssetSnapshotsParams) WithTicker(q string) *ListAssetSnapshotsParams {
	p.Ticker = &q
	return &p
}

func (p ListAssetSnapshotsParams) WithType(q string) *ListAssetSnapshotsParams {
	p.Type = &q
	return &p
}

func (p ListAssetSnapshotsParams) WithTickersByComparison(c Comparator, q string) *ListAssetSnapshotsParams {
	switch c {
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

// Deprecated: Please use UniversalSnapshot types instead of AssetSnapshot types.
type ListAssetSnapshotsResponse ListUniversalSnapshotsResponse
