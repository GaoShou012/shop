package api

import (
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type Shops struct {}

/*
	客户端，微信小程序，可以通过此接口来查询商店的分类栏
	商品种类栏目
*/
func (api *Shops) GoodsTypes(ctx *gin.Context) {
	var rows []models.GoodsTypes
	res := models.DB.Select([]string{"id","name","icon"}).Order("sort desc").Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx,res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx,len(rows),rows)
}

func (api *Shops) GoodsList(ctx *gin.Context) {
	type _params struct {
		TypeId uint64
		app.PageParams
	}

	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx,err)
		return
	}

	err = params.CheckPageParams()
	if err != nil {
		app.ResponseError(ctx,err)
		return
	}

	// 商品列表库（公版）
	// 提取出相应的 商品ID
	var total uint64
	var rows []models.ShopsGoodsPublic
	db := models.DB.Where(models.ShopsGoodsPublic{GoodsTypeId:&params.TypeId})
	db.Count(&total)
	db = db.Select([]string{"id","goods_type_id","goods_id"})
	res := db.Order("sort desc").Offset(params.Page*params.PageSize).Limit(params.PageSize).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx,err)
		return
	}

	// 得到相应的商品ID，把商品信息搜索出来
	var goodsList []models.Goods
	for _,row := range rows {
		goods := models.Goods{}
		res := models.DB.Select([]string{"`id`","`name`","`export_price`","`image`","`desc`"}).First(&goods,*row.GoodsId)
		if res.Error != nil {
			app.ResponseError(ctx,res.Error)
			return
		}
		goodsList = append(goodsList,goods)
	}

	app.ResponseSearchSuccess(ctx,total,goodsList)
}


func (api *Shops) Search(ctx *gin.Context) {
	type _params struct {
		Name string
	}

	params := _params{}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx,err)
		return
	}

	var values []models.Goods
	res := models.DB.Select([]string{"`id`","`name`","`export_price`","`image`","`desc`"}).Where("`name` like ?", "%" + params.Name + "%").Limit(20).Find(&values)
	if res.Error != nil {
		app.ResponseError(ctx,err)
		return
	}

	app.ResponseSearchSuccess(ctx,len(values),values)
}

func (api *Shops) Update(context *gin.Context) {
	id,err := app.ParseUintFromPost(context,"id")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	status,err := app.ParseUintFromPost(context,"status")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	firstBusinessTime,err := app.ParseTimeFromPost(context,"firstBusinessTime")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	numOfDesk,err := app.ParseUintFromPost(context,"numOfDesk")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	numOfChair,err := app.ParseUintFromPost(context,"numOfChair")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	shopCode := context.PostForm("shopCode")
	shopAddr := context.PostForm("shopAddr")
	shopCoor := context.PostForm("shopCoor")
	desc := context.PostForm("desc")

	m := models.Shops{
		Status:            &status,
		FirstBusinessTime: firstBusinessTime,
		ShopCode:          &shopCode,
		ShopAddr:          &shopAddr,
		ShopCoor:          &shopCoor,
		NumOfDesk:         &numOfDesk,
		NumOfChair:        &numOfChair,
		Desc:              &desc,
	}

	res := models.DB.Model(models.Shops{Model:models.Model{ID:&id}}).Updates(m)
	if res.Error != nil {
		app.ResponseError(context,res.Error)
		return
	}

	app.ResponseSuccess(context,"更新成功")
}

func (api *Shops) Insert(context *gin.Context) {
	status,err := app.ParseUintFromPost(context,"status")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	firstBusinessTime,err := app.ParseTimeFromPost(context,"firstBusinessTime")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	numOfDesk,err := app.ParseUintFromPost(context,"numOfDesk")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	numOfChair,err := app.ParseUintFromPost(context,"numOfChair")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	shopCode := context.PostForm("shopCode")
	shopAddr := context.PostForm("shopAddr")
	shopCoor := context.PostForm("shopCoor")
	desc := context.PostForm("desc")

	m := models.Shops{
		Status:            &status,
		FirstBusinessTime: firstBusinessTime,
		ShopCode:          &shopCode,
		ShopAddr:          &shopAddr,
		ShopCoor:          &shopCoor,
		NumOfDesk:         &numOfDesk,
		NumOfChair:        &numOfChair,
		Desc:              &desc,
	}

	res := models.DB.Table(models.ShopTableName).Create(m)
	if res.Error != nil {
		app.ResponseError(context,res.Error)
		return
	}

	app.ResponseSuccess(context,"新增店铺成功")
}
