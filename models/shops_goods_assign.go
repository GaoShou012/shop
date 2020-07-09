package models

const (
	ShopsGoodsAssignTable = "shops_goods_ass"
)

type ShopsGoodsAssigns struct {
	Model
	ShopId      *uint64
	GoodsTypeId *uint64
	GoodsId     *uint64
	Sort        *uint64 `gorm:"default:0"`
	Visible     *bool   `gorm:"default:false"`
}
