package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

/*
	属于店铺的 商品类目栏
*/
type ApiShopsGoodsTypes struct{}

func (api *ApiShopsGoodsTypes) Insert(ctx *gin.Context) {
	var params struct {
		ShopId uint64
		Name   string
		Icon   string
		Desc   string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	total := 0
	res := models.DB.Model(models.Shops{}).Where("id=?", params.ShopId).Count(&total)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if total != 1 {
		app.ResponseError(ctx, fmt.Errorf("商店ID(%d)不存在", params.ShopId))
		return
	}

	newData := models.ShopsGoodsTypes{
		//Model:   models.Model{},
		ShopId:  &params.ShopId,
		Visible: nil,
		Sort:    nil,
		Name:    &params.Name,
		Icon:    &params.Icon,
		Desc:    &params.Desc,
	}
	res = models.DB.Debug().Create(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "添加成功")
}

func (api *ApiShopsGoodsTypes) Update(ctx *gin.Context) {
	var params struct {
		Id      uint64
		Visible bool
		Sort    uint64
		Name    string
		Icon    string
		Desc    string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.ShopsGoodsTypes{
		//Model:   models.Model{},
		ShopId:  nil,
		Visible: &params.Visible,
		Sort:    &params.Sort,
		Name:    &params.Name,
		Icon:    &params.Icon,
		Desc:    &params.Desc,
	}

	res := models.DB.Model(models.ShopsGoodsTypes{Model: models.Model{ID: &params.Id}}).Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *ApiShopsGoodsTypes) Search(ctx *gin.Context) {
	var params struct {
		ShopId uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	var total uint64
	var rows []models.ShopsGoodsTypes
	db := models.DB.Model(models.ShopsGoodsTypes{})
	db = db.Where("shop_id=?", params.ShopId)
	res := db.Count(&total)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	res = db.Order("sort desc").Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, len(rows), rows)
}
