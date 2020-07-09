package models

const (
	BackendViewShopsGoodsPublicTable = "backend_view_shops_goods_public"
	BackendViewShopsGoodsTable       = "backend_view_shops_goods"
	BackendViewShopsGoodsAssignTable = "backend_view_shops_goods_assign"
)

type BackendViewShopsGoodsPublic struct {
	Id               uint64  `json:"id"`
	Sort             uint64  `json:"sort"`
	GoodsTypeId      uint64  `json:"goodsTypeId"`
	GoodsId          uint64  `json:"goodsId"`
	GoodsName        string  `json:"goodsName"`
	GoodsImportPrice float32 `json:"goodsImportPrice"`
	GoodsExportPrice float32 `json:"goodsExportPrice"`
	GoodsImage       string  `json:"goodsImage"`
	GoodsDesc        string  `json:"goodsDesc"`
}

type BackendViewShopsGoodsStocks struct {
	Id          string  `json:"id"`
	Status      uint64  `json:"status"`
	Stock       uint64  `json:"stock"`
	Name        string  `json:"name"`
	ImportPrice float32 `json:"importPrice"`
	ExportPrice float32 `json:"exportPrice"`
	Image       string  `json:"image"`
	Desc        string  `json:"desc"`
}

type BackendViewShopsGoods struct {
	ShopId                 uint64  `json:"shopId"`
	GoodsId                uint64  `json:"goodsId"`
	GoodsName              string  `json:"goodsName"`
	GoodsImage             string  `json:"goodsImage"`
	GoodsStock             uint64  `json:"goodsStock"`
	GoodsExportPrice       float32 `json:"goodsExportPrice"`
	GoodsImportPrice       float32 `json:"goodsImportPrice"`
	GoodsExportPriceOrigin float32 `json:"goodsExportPriceOrigin"`
}

type BackendViewShopsGoodsAssign struct {
	Sort                   uint64  `json:"sort"`
	Visible                uint64  `json:"visible"`
	ShopId                 uint64  `json:"shopId"`
	GoodsTypeId            uint64  `json:"goodsTypeId"`
	GoodsId                uint64  `json:"goodsId"`
	GoodsName              string  `json:"goodsName"`
	GoodsImage             string  `json:"goodsImage"`
	GoodsStock             uint64  `json:"goodsStock"`
	GoodsExportPrice       float32 `json:"goodsExportPrice"`
	GoodsImportPrice       float32 `json:"goodsImportPrice"`
	GoodsExportPriceOrigin float32 `json:"goodsExportPriceOrigin"`
}

type BackendViewGoods struct {
	Model
	TypeName          *string   `json:"typeName"`
	TypeId            *uint64   `json:"typeId"`
	Name              *string   `json:"name"`
	ImportPrice       *float64 `json:"importPrice,omitempty"`
	ExportPriceAdvise *float64 `json:"exportPriceAdvise,omitempty"`
	Image             *string  `json:"image,omitempty"`
	Desc              *string  `json:"desc,omitempty"`
}
