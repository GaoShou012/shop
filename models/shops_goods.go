package models

const (
	ShopsGoodsTableName = "shops_goods"
	ShopsGoodsTable     = "shops_goods"
)

type ShopsGoods struct {
	Model
	ShopId           *uint64  `json:"shopId"`
	GoodsId          *uint64  `json:"goodsId"`
	GoodsName        *string  `json:"goodsName"`
	GoodsImage       *string  `json:"goodsImage"`
	GoodsExportPrice *float64 `json:"goodsPrice"`
	GoodsStock       *uint64  `json:"goodsStock" gorm:"default:0"`
}

func (m *ShopsGoods) Table() string {
	return "shops_goods"
}
