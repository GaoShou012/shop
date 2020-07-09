package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiFranchiseesShops struct {
}

/*
	绑定店铺
	店铺和加盟商账号进行绑定后，会员消费，才有利润分成
*/
func (api *ApiFranchiseesShops) Bind(ctx *gin.Context) {
	var params struct {
		FranchiseeUsername        string
		ShopCode                  string
		PlatformCommissionPercent uint64
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 查询店铺的ID
	// 查询店铺是否存在
	{
		exists, err := models.IsExistsV1(models.DB, &models.Shops{}, false, "shop_code=?", params.ShopCode)
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		if !exists {
			app.ResponseError(ctx, fmt.Errorf("商店编码(%s)不存在", params.ShopCode))
			return
		}
	}

	// 查询加盟商是否存在
	{
		exists, err := models.IsExistsV1(models.DB, &models.Franchisees{}, false, "username=?", params.FranchiseeUsername)
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		if !exists {
			app.ResponseError(ctx, fmt.Errorf("加盟商(%s)不存在", params.FranchiseeUsername))
			return
		}
	}

	// 绑定加盟商与店铺
	franchiseesShop := models.FranchiseesShops{
		//Model:                     models.Model{},
		//Enable:                    nil,
		FranchiseeUsername:        &params.FranchiseeUsername,
		ShopCode:                  &params.ShopCode,
		PlatformCommissionPercent: &params.PlatformCommissionPercent,
	}
	res := models.DB.Create(&franchiseesShop)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "绑定成功")
}

func (api *ApiFranchiseesShops) Unbind(ctx *gin.Context) {
	var params struct {
		FranchiseeUsername string
		ShopCode           string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	res := models.DB.Where("franchisee_username=? and shop_code=?", params.FranchiseeUsername, params.ShopCode).Unscoped().Delete(&models.FranchiseesShops{})
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("解绑失败，可能数据不存在"))
		return
	}

	app.ResponseSuccess(ctx, "解绑成功")
}

func (api *ApiFranchiseesShops) Update(ctx *gin.Context) {
	var params struct {
		FranchiseeUsername        string
		ShopCode                  string
		Enable                    bool
		PlatformCommissionPercent uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.FranchiseesShops{
		//Model:                     models.Model{},
		Enable: &params.Enable,
		//FranchiseeUsername:        nil,
		//ShopCode:                  nil,
		PlatformCommissionPercent: &params.PlatformCommissionPercent,
	}
	res := models.DB.Model(models.FranchiseesShops{}).Where("franchisee_username =? and shop_code=?", params.FranchiseeUsername, params.ShopCode).Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("更新失败，可能数据不存在"))
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *ApiFranchiseesShops) Search(ctx *gin.Context) {
	var params struct {
		FranchiseeUsername string
		ShopCode           string
		app.PageParams
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	db := models.DB.Model(models.FranchiseesShops{})
	if params.FranchiseeUsername != "" {
		db = db.Where("franchisee_username=?", params.FranchiseeUsername)
	}
	if params.ShopCode != "" {
		db = db.Where("shop_code=?", params.ShopCode)
	}

	var total uint
	var rows []models.FranchiseesShops

	res := db.Count(&total)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	res = db.Limit(params.PageSize).Offset(params.Offset()).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}
