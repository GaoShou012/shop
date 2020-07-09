package franchisees

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiShops struct {
}

/*
	可以查看代理商加盟的店铺
*/
func (api *ApiShops) Search(ctx *gin.Context) {
	//operator := app.Operator(ctx)

}

/*
	查看店铺的交易明细
*/
func (api *ApiShops) Orders(ctx *gin.Context) {
	operator := app.Operator(ctx)

	var params struct {
		ShopCode string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 查询店铺是否属于此加盟商的
	shop := models.FranchiseesShops{}
	res := models.DB.Table(models.FranchiseesShopsTable).Where("franchisee_name=? and shop_code=?", operator.Username, params.ShopCode).First(&shop)
	if res.Error != nil {
		if res.RecordNotFound() {
			app.ResponseError(ctx, fmt.Errorf("店铺编码不存在"))
		} else {
			app.ResponseError(ctx, res.Error)
		}
		return
	}
}
