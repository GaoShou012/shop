package models

const (
	ShopsGoodsTypesTable = "shops_goods_types"
)

type ShopsGoodsTypes struct {
	Model
	ShopId  *uint64 `json:"shopId"`
	Visible *bool   `json:"visible" gorm:"default:0"`
	Sort    *uint64 `json:"sort,omitempty" gorm:"default:0"`
	Name    *string `json:"name,omitempty"`
	Icon    *string `json:"icon,omitempty"`
	Desc    *string `json:"desc,omitempty"`
}

func (m *ShopsGoodsTypes) Table() string {
	return "shops_goods_types"
}
