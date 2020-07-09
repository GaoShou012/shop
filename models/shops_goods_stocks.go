package models

type ShopsGoodsStocks struct {
	Model
	ShopId      *uint64
	GoodsTypeId *uint64
	GoodsId     *uint64
	Stock       *uint64
}

func (m *ShopsGoodsStocks) Table() string {
	return "shops_goods_stocks"
}