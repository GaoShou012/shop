package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"net/http"
)

/*
	商品种类，并不是提供给商店使用
	而是辅助，查找和浏览，目前商品
*/
type ApiGoodsType struct{}

func (api *ApiGoodsType) Insert(ctx *gin.Context) {
	var params struct {
		Sort uint64
		Name string
		Icon string
		Desc string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	goodsType := models.GoodsTypes{
		//Model:     models.Model{},
		Sort:      &params.Sort,
		Name:      &params.Name,
		Icon:      &params.Icon,
		Desc:      &params.Name,
		Reference: new(uint64),
	}

	res := models.DB.Create(&goodsType)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "添加商品种类成功")
}

func (api *ApiGoodsType) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "接口暂未开放",
	})
	return

	var params struct {
		Id uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	goodsType := models.GoodsTypes{
		Model:     models.Model{ID: &params.Id},
		Sort:      nil,
		Name:      nil,
		Icon:      nil,
		Desc:      nil,
		Reference: nil,
	}

	res := models.DB.Delete(&goodsType)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected <= 0 {
		app.ResponseError(ctx, fmt.Errorf("删除失败，可能ID(%d)不存在", params.Id))
		return
	}

	app.ResponseSuccess(ctx, "删除成功")
}

func (api *ApiGoodsType) Update(ctx *gin.Context) {
	var params struct {
		Id   uint64
		Sort uint64
		Name string
		Icon string
		Desc string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.GoodsTypes{
		//Model:     models.Model{},
		Sort:      &params.Sort,
		Name:      &params.Name,
		Icon:      &params.Icon,
		Desc:      &params.Desc,
		Reference: nil,
	}

	db := models.DB.Model(&models.GoodsTypes{})
	db = db.Where("id=?", params.Id)
	res := db.Debug().Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected <= 0 {
		app.ResponseError(ctx, fmt.Errorf("更新失败，可能ID(%d)不存在", params.Id))
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *ApiGoodsType) Search(ctx *gin.Context) {
	var total uint64
	var rows []models.GoodsTypes
	res := models.DB.Table(models.GoodsTypesTable).Count(&total)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	res = models.DB.Table(models.GoodsTypesTable).Order("sort desc").Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, len(rows), rows)
}
