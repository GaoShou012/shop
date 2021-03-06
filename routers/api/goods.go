package api

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type Goods struct{}

func (api *Goods) Insert(ctx *gin.Context) {
	type _params struct {
		TypeId            uint64
		Name              string
		ImportPrice       float64
		ExportPriceAdvise float64
		Image             string
		Desc              string
	}

	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 搜搜typeId是否存在
	goodsType := models.GoodsType{}
	res := models.DB.Find(&goodsType, params.TypeId)
	if res.Error != nil {
		if res.RecordNotFound() {
			app.ResponseError(ctx, fmt.Errorf("商品类型ID不存在"))
		} else {
			app.ResponseError(ctx, err)
		}
		return
	}

	// 添加商品
	goods := &models.Goods{
		Name:              &params.Name,
		ImportPrice:       &params.ImportPrice,
		ExportPriceAdvise: &params.ExportPriceAdvise,
		TypeId:            goodsType.ID,
		Image:             &params.Image,
		Desc:              &params.Desc,
	}
	res = models.DB.Create(goods)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "添加商品成功")
}

func (api *Goods) Update(ctx *gin.Context) {
	type _params struct {
		Id                uint64
		Name              string
		ImportPrice       float64
		ExportPriceAdvise float64
		Image             string
		Desc              string
	}

	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.Goods{
		Name:              &params.Name,
		ImportPrice:       &params.ImportPrice,
		ExportPriceAdvise: &params.ExportPriceAdvise,
		TypeId:            nil,
		Image:             &params.Image,
		Desc:              &params.Desc,
	}
	res := models.DB.Where("id = ?", params.Id).Updates(newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *Goods) Search(ctx *gin.Context) {
	type _params struct {
		Id   uint64
		Name string
		app.PageParams
	}

	params := _params{}
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
	var rows []models.Goods
	db := models.DB.Model(models.Goods{})
	if params.Id > 0 {
		db = db.Where("`id` = ?", params.Id)
	}
	if params.Name != "" {
		db = db.Where("`name` like ?", "%"+params.Name+"%")
	}
	db.Count(&total)
	res := db.Offset(params.Page * params.PageSize).Limit(params.PageSize).Debug().Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}

func (api *Goods) Delete(context *gin.Context) {

}
