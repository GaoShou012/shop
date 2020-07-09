package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiShopsGoodsAssign struct {
}

/*
	把店铺内的商品，分配到指定的商品类型下
	可以查询店铺内商品分类下的 商品列表
*/
func (api *ApiShopsGoodsAssign) Insert(ctx *gin.Context) {
	var params struct {
		ShopId      uint64
		GoodsTypeId uint64
		GoodsId     uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 查询店铺是否存在
	{
		exists, err := models.IsExistsV1(models.DB, &models.Shops{}, true, params.ShopId)
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		if !exists {
			app.ResponseError(ctx, fmt.Errorf("店铺ID(%d)不存在", params.ShopId))
			return
		}
	}
	// 查询分类是否存在
	{
		exists, err := models.IsExistsV1(models.DB, &models.ShopsGoodsTypes{}, true, params.GoodsTypeId)
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		if !exists {
			app.ResponseError(ctx, fmt.Errorf("分类ID(%d)不存在", params.GoodsTypeId))
			return
		}
	}

	assign := models.ShopsGoodsAssigns{
		//Model:       models.Model{},
		ShopId:      &params.ShopId,
		GoodsTypeId: &params.GoodsTypeId,
		GoodsId:     &params.GoodsId,
		//Sort:        nil,
		//Status:      nil,
	}
	res := models.DB.Create(&assign)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "分配成功")
}

/*
	把指定类型下的商品移除
*/
func (api *ApiShopsGoodsAssign) Delete(ctx *gin.Context) {
	var params struct {
		ShopId      uint64
		GoodsTypeId uint64
		GoodsId     uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	res := models.DB.Where("shop_id=? and goods_type_id=? and goods_id=?", params.ShopId, params.GoodsTypeId, params.GoodsId).Unscoped().Delete(&models.ShopsGoodsAssigns{})
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("删除失败，可能记录已经不存在"))
		return
	}

	app.ResponseSuccess(ctx, "删除成功")
}

/*

 */
func (api *ApiShopsGoodsAssign) Update(ctx *gin.Context) {
	var params struct {
		ShopId      uint64
		GoodsTypeId uint64
		GoodsId     uint64
		Sort        uint64
		Visible     bool
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.ShopsGoodsAssigns{
		//Model:       models.Model{},
		//ShopId:      nil,
		//GoodsTypeId: nil,
		//GoodsId:     nil,
		Sort:    &params.Sort,
		Visible: &params.Visible,
	}
	res := models.DB.Model(models.ShopsGoodsAssigns{}).Where("shop_id=? and goods_type_id=? and goods_id=?", params.ShopId, params.GoodsTypeId, params.GoodsId).Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("更新失败"))
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

/*
	查询的时候，需要带有店铺ID，商品分类ID
	就可以查询出，所有的商品(店铺内)
*/
func (api *ApiShopsGoodsAssign) Search(ctx *gin.Context) {
	var params struct {
		ShopId      uint64
		GoodsTypeId uint64
		GoodsName   string
		app.PageParams
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}
	err = params.CheckPageParams()
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	var total uint64
	var rows []models.BackendViewShopsGoodsAssign
	db := models.DB.Model(models.BackendViewShopsGoodsAssign{})
	db = db.Where("shop_id = ? and goods_type_id=?", params.ShopId, params.GoodsTypeId)
	if params.GoodsName != "" {
		db = db.Where("goods_name = ?", params.GoodsName)
	}
	db.Count(&total)
	res := db.Limit(params.PageSize).Offset(params.Offset()).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}
