package users

import (
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiUsers struct {
}

/*
	查询店铺内的商品类目栏
*/
func (api *ApiUsers) Sidecar(ctx *gin.Context) {
	var rows []models.ShopsGoodsTypes

	res := models.DB.Model(models.ShopsGoodsTypes{}).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, len(rows), rows)
}

/*
	查询某个商品类目下的商品列表
*/
func (api *ApiUsers) GoodsList(ctx *gin.Context) {
	var params struct {
		ShopId      uint64
		GoodsTypeId uint64
		app.PageParams
	}

	var total uint64
	var rows []models.ShopsGoods

	db := models.DB.Model(models.ShopsGoods{})
	db = db.Where("shop_id=?", params.ShopId)
	db = db.Where("goods_type_id=?", params.GoodsTypeId)
	res := db.Count(&total)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	res = db.Offset(params.Offset()).Limit(params.PageSize).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}

/*
	下单
*/
func (api *ApiUsers) Order(ctx *gin.Context) {

}
