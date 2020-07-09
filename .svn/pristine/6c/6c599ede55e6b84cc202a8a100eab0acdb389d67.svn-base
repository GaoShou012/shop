package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"net/http"
)

type ApiShopsOwner struct {}

func (api *ApiShopsOwner) Insert(context *gin.Context) {
	shopCode := context.PostForm("shopCode")
	franchiseeName := context.PostForm("franchiseeName")
	platformCommissionPercent,err := app.ParseUintFromPost(context,"platformCommissionPercent")
	if err != nil {
		app.ResponseError(context,err)
		return
	}

	// 查询商店编码是否存在
	res := models.DB.Model(models.Shops{}).Where("shop_code = ?",shopCode).Find(&models.Shops{})
	if res.Error != nil {
		app.ResponseError(context,fmt.Errorf("商店码 %v",res.Error))
		return
	}

	// 查询加盟商账户是否存在
	res = models.DB.Model(models.Franchisees{}).Where("username = ?",franchiseeName).Find(&models.Franchisees{})
	if res.Error != nil {
		app.ResponseError(context,fmt.Errorf("加盟商 %v",res.Error))
		return
	}

	newData := &models.ShopsOwners{
		ShopCode:                  &shopCode,
		FranchiseeName:            &franchiseeName,
		PlatformCommissionPercent: &platformCommissionPercent,
	}

	res = models.DB.Model(models.ShopsOwners{}).Create(newData)
	if res.Error != nil {
		app.ResponseError(context,res.Error)
		return
	}

	app.ResponseSuccess(context,"商店关联持有者成功")
}

func (api *ApiShopsOwner) Update(context *gin.Context) {
	id,err := app.ParseUintFromPost(context,"id")
	if err != nil {
		app.ResponseError(context,err)
		return
	}
	franchiseeName := context.PostForm("franchiseeName")
	platformCommissionPercent,err := app.ParseUintFromPost(context,"platformCommissionPercent")
	if err != nil {
		app.ResponseError(context,err)
		return
	}

	newData := &models.ShopsOwners{
		//Model:                     models.Model{},
		//ShopCode:                  nil,
		FranchiseeName:            &franchiseeName,
		PlatformCommissionPercent: &platformCommissionPercent,
	}

	res := models.DB.Model(models.ShopsOwners{}).Where("id = ?",id).Updates(newData)
	if res.Error != nil {
		app.ResponseError(context,err)
		return
	}
	if res.RowsAffected == 0 {
		app.ResponseError(context,fmt.Errorf("更新失败"))
		return
	}

	app.ResponseSuccess(context,"更新成功")
}

func (api *ApiShopsOwner) Search(context *gin.Context) {
	shopCode := context.PostForm("shopCode")
	franchiseeName := context.PostForm("franchiseeName")

	var total uint = 0
	var values []models.ShopsOwners
	db := models.DB.Model(models.ShopsOwners{})
	if shopCode != "" {
		db = db.Where("shop_code = ?",shopCode)
	}
	if franchiseeName != "" {
		db = db.Where("franchisee_name = ?",franchiseeName)
	}
	db.Count(&total)
	res := db.Debug().Find(&values)
	if res.Error != nil {
		app.ResponseError(context,res.Error)
		return
	}

	context.JSON(http.StatusOK,gin.H{
		"code":0,
		"message":"搜索成功",
		"total":total,
		"list":values,
	})
}