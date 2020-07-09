package models

const (
	ConsumersOrdersDetailsTableName = "consumers_orders_details"
)

type ConsumersOrdersDetails struct {
	Model
	OrderNo    *uint64
	GoodsId    *uint64
	GoodsName  *string
	GoodsImage *string
}

type UsersOrdersDetails struct {
	Model
	PlatformOrderNo        *uint64
	GoodsId                *uint64
	GoodsName              *string
	GoodsImage             *string
	GoodsExportPrice       *float64
	GoodsImportPrice       *float64
	GoodsExportPriceAdvise *float64
	GoodsPiece             *uint64
}
