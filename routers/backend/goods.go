package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type Goods struct{}

func (api *Goods) Insert(ctx *gin.Context) {
	var params struct {
		TypeId            uint64
		Name              string
		ImportPrice       float64
		ExportPriceAdvise float64
		Image             string
		Desc              string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 添加商品
	goods := models.Goods{
		TypeId:            &params.TypeId,
		Name:              &params.Name,
		ImportPrice:       &params.ImportPrice,
		ExportPriceAdvise: &params.ExportPriceAdvise,
		Image:             &params.Image,
		Desc:              &params.Desc,
	}
	res := models.DB.Create(&goods)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "添加商品成功")
}

func (api *Goods) Delete(ctx *gin.Context) {
	var params struct {
		Id uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	oldData := models.Goods{
		Model:             models.Model{ID: &params.Id},
		TypeId:            nil,
		Name:              nil,
		ImportPrice:       nil,
		ExportPriceAdvise: nil,
		Image:             nil,
		Desc:              nil,
	}

	res := models.DB.Unscoped().Delete(&oldData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, err)
		return
	}

	app.ResponseSuccess(ctx, "删除成功")
}

func (api *Goods) Update(ctx *gin.Context) {
	var params struct {
		Id                uint64
		TypeId            uint64
		Name              string
		ImportPrice       float64
		ExportPrice       float64
		ExportPriceAdvise float64
		Image             string
		Desc              string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}
	newData := models.Goods{
		//Model:       models.Model{},
		TypeId:            &params.TypeId,
		Name:              &params.Name,
		ImportPrice:       &params.ImportPrice,
		ExportPriceAdvise: &params.ExportPriceAdvise,
		Image:             &params.Image,
		Desc:              &params.Desc,
	}
	res := models.DB.Model(models.Goods{}).Where("id = ?", params.Id).Debug().Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected <= 0 {
		app.ResponseError(ctx, fmt.Errorf("更新失败，可能ID不存在(%d)", params.Id))
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *Goods) Search(ctx *gin.Context) {
	var params struct {
		TypeId int
		Name   string
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
	var rows []models.BackendViewGoods
	db := models.DB.Model(models.BackendViewGoods{})
	if params.TypeId > 0 {
		db = db.Where("type_id = ?", params.TypeId)
	}
	if params.Name != "" {
		db = db.Where("name like ?", "%"+params.Name+"%")
	}
	db.Count(&total)
	res := db.Offset(params.Page * params.PageSize).Limit(params.PageSize).Debug().Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}
