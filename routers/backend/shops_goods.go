package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

/*
	商品复制，从商品库，复制数据到 商店的商品库
	不能进行新增，不能进行删除，只能复制

	搜索
	更新库存
	更新零售价
*/
type ApiShopsGoods struct{}

/*
	从商品库中复制所有商品到本店铺
*/
func (api *ApiShopsGoods) Copy(ctx *gin.Context) {
	var params struct {
		ShopId uint64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	// 商店ID是否存在
	exists, err := models.IsExistsV1(models.DB, &models.Shops{}, false, params.ShopId)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}
	if !exists {
		app.ResponseError(ctx, fmt.Errorf("商店ID(%d)不存在", params.ShopId))
		return
	}

	// 商品库查询所有的商品
	var allGoods []models.Goods
	res := models.DB.Model(models.Goods{}).Find(&allGoods)
	if res.Error != nil {
		app.ResponseError(ctx, err)
		return
	}

	{
		tx := models.DB.Begin()
		for _, goods := range allGoods {
			shopGoods := models.ShopsGoods{
				//Model:            models.Model{},
				ShopId:           &params.ShopId,
				GoodsId:          goods.ID,
				GoodsName:        goods.Name,
				GoodsImage:       goods.Image,
				GoodsExportPrice: goods.ExportPriceAdvise,
				GoodsStock:       nil,
			}

			// 当商品不存在的时候，才进行拷贝
			count := 0
			res := tx.Model(models.ShopsGoods{}).Where("shop_id=? and goods_id=?", params.ShopId, goods.ID).Count(&count)
			if res.Error != nil {
				tx.Rollback()
				app.ResponseError(ctx, res.Error)
				return
			}
			if count == 0 {
				res := tx.Create(&shopGoods)
				if res.Error != nil {
					tx.Rollback()
					app.ResponseError(ctx, res.Error)
					return
				}
			}
		}
		tx.Commit()
	}

	app.ResponseSuccess(ctx, "商品拷贝完成")
}

// 更新售卖价格
func (api *ApiShopsGoods) Update(ctx *gin.Context) {
	var params struct {
		ShopId           uint64
		GoodsId          uint64
		GoodsExportPrice float64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.ShopsGoods{
		GoodsExportPrice: &params.GoodsExportPrice,
	}
	db := models.DB.Model(models.ShopsGoods{})
	db = db.Where("shop_id = ? and goods_id = ?", params.ShopId, params.GoodsId)
	res := db.Updates(&newData)
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

/*
	更新库存
*/
func (api *ApiShopsGoods) UpdateStock(ctx *gin.Context) {
	var params struct {
		ShopId  uint64
		GoodsId uint64
		Stock   uint64
		Amount  int64
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	stockTmp := int64(params.Stock) + params.Amount
	if stockTmp < 0 {
		app.ResponseError(ctx, fmt.Errorf("库存不能少于0"))
		return
	}
	stock := uint64(stockTmp)

	newData := models.ShopsGoods{
		GoodsStock: &stock,
	}
	res := models.DB.Table(models.ShopsGoodsTable).Where("shop_id=? and goods_id=? and goods_stock=?", params.ShopId, params.GoodsId, params.Stock).Update(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("更新库存失败，请刷新，注意数据变化"))
		return
	}

	app.ResponseSuccess(ctx, "修改库存成功")
}

/*
	商品的货品搜索
	没有分类的搜索，只能通过名字进行搜索
*/
func (api *ApiShopsGoods) Search(ctx *gin.Context) {
	var params struct {
		ShopId    uint64
		GoodsName string
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
	var rows []models.BackendViewShopsGoods
	db := models.DB.Table(models.BackendViewShopsGoodsTable)
	db = db.Where("shop_id = ?", params.ShopId)
	if params.GoodsName != "" {
		db = db.Where("goods_name = ?", params.GoodsName)
	}
	db.Count(&total)
	res := db.Offset(params.Offset()).Limit(params.PageSize).Find(&rows)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSearchSuccess(ctx, total, rows)
}
