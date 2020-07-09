package main

import (
	_ "github.com/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

	"context"
	"fmt"
	"github.com/app"
	"github.com/jinzhu/gorm"
	"github.com/models"
	"github.com/proto"

	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
)

type UsersOrdersHandler struct {
	Service micro.Service
	DB      *gorm.DB
	mutex []sync.Mutex
}

func (s *UsersOrdersHandler) Init() {
	s.mutex = make([]sync.Mutex,100)
}
func (s *UsersOrdersHandler) order() {

}

/*
	status = 2
	用户已经下单成功，不用重复下单
	status = 3
	订单号已经被占用，下单失败

	code = 1
	创建订单失败
	code = 2
	订单号已经存在
*/
func (s *UsersOrdersHandler) Order(ctx context.Context, req *proto.UsersOrdersReq, rsp *proto.UsersOrdersRsp) error {
	// 检查订单号是否已经存在
	// 需要使用 unscoped 查询已经被deleted的订单
	exists, err := models.IsExistsV1(s.DB, &models.UsersOrders{}, true, "platform_order_no=?", req.PlatformOrderNo)
	if err != nil {
		return err
	}
	if exists {
		rsp.Code = 2
		rsp.Message = "订单号已经存在，创建订单失败"
		return nil
	}

	// 查询店铺数据
	shop := models.Shops{}
	res := s.DB.Model(models.Shops{}).Where("shop_code=?", req.ShopCode).First(&shop)
	if res.Error != nil {
		if res.RecordNotFound() {
			rsp.Code = 1
			rsp.Message = fmt.Sprintf("商店(%s)不存在", req.ShopCode)
			return nil
		} else {
			return res.Error
		}
	}

	// 读取用户信息
	user := models.Users{}
	res = s.DB.Model(models.Users{}).First(&user, req.UserId)
	if res.Error != nil {
		if res.RecordNotFound() {
			rsp.Code = 1
			rsp.Message = "读取用户信息失败"
			return nil
		} else {
			return res.Error
		}
	}

	// 用户订单
	userOrder := models.UsersOrders{
		//Model:    models.Model{},
		PlatformOrderNo:  &req.PlatformOrderNo,
		ClientToken:      &req.ClientToken,
		UserId:           &req.UserId,
		Status:           nil,
		TotalImportPrice: new(float64),
		TotalExportPrice: new(float64),
		ShopCode:         shop.ShopCode,
		ShopAddr:         shop.ShopAddr,
		PaidAt:           nil,
	}
	userOrderDetails := make([]models.UsersOrdersDetails, 0, len(req.Goods))

	for _, row := range req.Goods {
		// 读取商品库中的商品信息
		goods := models.Goods{}
		res := s.DB.Model(goods).First(&goods, row.GoodsId)
		if res.Error != nil {
			if res.RecordNotFound() {
				rsp.Code = 1
				rsp.Message = fmt.Sprintf("商品ID(%d)不存在与商品库", row.GoodsId)
				return nil
			} else {
				return res.Error
			}
		}

		// 读取商品库中的商品信息
		shopGoods := models.ShopsGoods{}
		res = s.DB.Model(shopGoods).Where("shop_code=?", shop.ShopCode).First(&shopGoods)
		if res.Error != nil {
			if res.RecordNotFound() {
				rsp.Code = 1
				rsp.Message = fmt.Sprintf("商品ID(%d)不存在与商店商品库", row.GoodsId)
				return nil
			} else {
				return res.Error
			}
		}

		// 检查库存是否足够
		if *shopGoods.GoodsStock < row.GoodsPiece {
			rsp.Code = 1
			rsp.Message = fmt.Sprintf("%s 货存不足", *shopGoods.GoodsName)
			return nil
		}

		// 合计总售卖价
		// 合计总进货价
		goodsPiece := decimal.NewFromInt(int64(row.GoodsPiece))
		totalImportPrice, _ := decimal.NewFromFloat(*goods.ImportPrice).Mul(goodsPiece).Float64()
		totalExportPrice, _ := decimal.NewFromFloat(*shopGoods.GoodsExportPrice).Mul(goodsPiece).Float64()
		*userOrder.TotalImportPrice += totalImportPrice
		*userOrder.TotalExportPrice += totalExportPrice

		// 合成订单明细
		detail := models.UsersOrdersDetails{
			//Model:            models.Model{},
			PlatformOrderNo:        userOrder.PlatformOrderNo,
			GoodsId:                shopGoods.GoodsId,
			GoodsName:              shopGoods.GoodsName,
			GoodsImage:             shopGoods.GoodsImage,
			GoodsExportPrice:       shopGoods.GoodsExportPrice,
			GoodsImportPrice:       goods.ImportPrice,
			GoodsExportPriceAdvise: goods.ExportPriceAdvise,
			GoodsPiece:             &row.GoodsPiece,
		}

		userOrderDetails = append(userOrderDetails, detail)
	}

	// 事务操作
	{
		tx := s.DB.Begin()

		res := tx.Create(&userOrder)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		// 保存订单明细
		for _, detail := range userOrderDetails {
			res := tx.Create(detail)
			if res.Error != nil {
				tx.Rollback()
				return res.Error
			}
		}

		res = tx.Commit()
		if res.Error != nil {
			return res.Error
		}

		// Broker 投递任务
		b := s.Service.Options().Broker
		err := b.Publish("orders.new", &broker.Message{
			Header: nil,
			Body:   nil,
		}, rabbitmq.DeliveryMode(2))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"platformOrderNo": *userOrder.PlatformOrderNo,
				"message":         "用户下单，事务提交成功，投递order.new订单号失败",
			}).Error(err)
		}
	}

	rsp.PlatformOrderNo = *userOrder.PlatformOrderNo
	rsp.TotalExportPrice = *userOrder.TotalExportPrice
	return nil
}

func main() {
	serviceName := "micro.service.users"

	err := app.LogrusRedisHook(serviceName)
	if err != nil {
		panic(err)
	}

	db, err := app.Gorm()
	if err != nil {
		panic(err)
	}

	// broker
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
		rabbitmq.ExchangeName("orders"),
	)
	if err := b.Init(); err != nil {
		panic(err)
	}
	if err := b.Connect(); err != nil {
		panic(err)
	}

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Broker(b),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
	)
	serviceHandler := UsersOrdersHandler{Service: service, DB: db}
	err = proto.RegisterUsersOrdersHandler(service.Server(), &serviceHandler)
	if err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		panic(err)
	}
}
