package models

type FranchiseesOrdersDetails struct {
	Model
	GoodsId                   *uint64  `json:"goodsId"`
	GoodsName                 *string  `json:"goodsName"`
	GoodsImage                *string  `json:"goodsImage"`
	GoodsImportPrice          *float64 `json:"goodsImportPrice"`
	GoodsExportPrice          *float64 `json:"goodsExportPrice"`
	GoodsExportPriceAdvise    *float64 `json:"goodsExportPriceAdvise"`
	PlatformOrderNo           *uint64  `json:"platformOrderNo"`
	PlatformCommissionPercent *uint64  `json:"platformCommissionPercent"`
	PlatformCommission        *float64  `json:"platformCommission"`
	FranchiseeUsername        *string  `json:"franchiseeUsername"`
	FranchiseeCommission      *float64 `json:"franchiseeCommission"`
}
