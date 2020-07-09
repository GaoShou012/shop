package backend

import (
	"github.com/app"
	"github.com/gin-gonic/gin"
)

type ApiConsumers struct {

}

/*
	消费者，下单操作，生成一个订单，有商品的明细
*/
func (api *ApiConsumers) Order(ctx *gin.Context) {
	//type _goods struct {
	//	Id uint					`json:"id"`
	//	Name string				`json:"name"`
	//	Price float32			`json:"price"`
	//	Piece uint				`json:"piece"`
	//}
	//type _params struct {
	//	ShopId uint				`json:"shopId"`
	//	Goods []_goods			`json:"goods"`
	//}
	//params := &_params{}
	//err := ctx.BindJSON(&params)
	//if err != nil {
	//	app.ResponseError(ctx,err)
	//	return
	//}
	//
	//// 查询店铺是否存在
	//shopInfo := &models.Shops{}
	//res := models.DB.Model(&models.Shops{}).First(shopInfo,params.ShopId)
	//if res.Error != nil {
	//	app.ResponseError(ctx,res.Error)
	//	return
	//}
	//
	//// 计算商品数量
	//if len(params.Goods) <= 0 {
	//	app.ResponseError(ctx,fmt.Errorf("商品数量不能少于1件"))
	//	return
	//}
	//
	//// 用户信息
	//user := app.Operator(ctx)
	//// 商品的信息
	//var allGoods []*models.ShopsGoods
	//
	//for _,row := range params.Goods {
	//	goods := &models.ShopsGoods{}
	//
	//	//// 查询数据
	//	//res := models.DB.Model(&models.ShopsGoods{}).Where("shop_id=? and goods_id=?",params.ShopId,row.Id).First(goods)
	//	//if res.Error != nil {
	//	//	app.ResponseError(ctx,fmt.Errorf("%s %s",row.Name,res.Error.Error()))
	//	//	return
	//	//}
	//	//
	//	//// 校验商品数据
	//	//if row.Price != *goods.Price {
	//	//	app.ResponseError(ctx,fmt.Errorf("%s 商品参数有错误",row.Name))
	//	//	return
	//	//}
	//	//if row.Name != *goods.GoodsName {
	//	//	app.ResponseError(ctx,fmt.Errorf("%s 商品参数有错误",row.Name))
	//	//	return
	//	//}
	//
	//	// 记录所有商品信息
	//	allGoods = append(allGoods,goods)
	//}

	//cli := proto.NewOrderNoService("order.no",app.Service().Client())
	//rsp,err := cli.Gen(context.TODO(),&proto.OrderNoGenRequest{})
	//if err != nil {
	//	app.ResponseError(ctx,err)
	//	return
	//}
	//orderNo := rsp.No

	//{
	//	tx :=models.DB.Begin()
	//
	//	// 计算订单总额
	//	total := float32(0)
	//	for _,v := range allGoods {
	//		total += *v.Price
	//	}
	//
	//	// 创建订单
	//	order := &models.ConsumersOrders{
	//		ConsumerId: user.Id,
	//		OrderNo:    &orderNo,
	//		Status:     nil,
	//		Total:      &total,
	//		ShopId:     shopInfo.ID,
	//		ShopAddr:   shopInfo.ShopAddr,
	//		ShopCode:   shopInfo.ShopCode,
	//		PaymentAt:  nil,
	//	}
	//	res := tx.Table(models.ConsumersOrderTableName).Create(order)
	//	if res.Error != nil {
	//		app.ResponseError(ctx,res.Error)
	//		tx.Rollback()
	//	}
	//
	//	// 创建订单明细
	//	for _,v := range allGoods {
	//		detail := &models.ConsumersOrdersDetails{
	//			OrderNo:    &orderNo,
	//			GoodsId:    v.ID,
	//			GoodsName:  v.GoodsName,
	//			GoodsImage: v.GoodsImage,
	//		}
	//		res := models.DB.Table(models.ConsumersOrdersDetailsTableName).Create(detail)
	//		if res.Error != nil {
	//			app.ResponseError(ctx,res.Error)
	//			tx.Rollback()
	//			return
	//		}
	//	}
	//
	//	tx.Commit()
	//}

	app.ResponseSuccess(ctx,"创建订单成功")
}

func (api *ApiConsumers) OrderInfo(ctx *gin.Context) {

}