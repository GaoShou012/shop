package models

const (
	ShopsGoodsPublicTableName = "shops_goods_publics"
)

type ShopsGoodsPublic struct {
	Model
	Status *uint64					`json:"status,omitempty"`
	Sort *uint64					`json:"sort,omitempty"`
	GoodsTypeId *uint64				`json:"goodsTypeId,omitempty"`
	GoodsId *uint64					`json:"goodsId,omitempty"`
}

func (m *ShopsGoodsPublic) Table() string {
	return "shops_goods_public"
}