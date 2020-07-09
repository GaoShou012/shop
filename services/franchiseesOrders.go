package main

import (
	"github.com/app"
	_ "github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"github.com/models"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	serviceName := "orders.new.franchisees"

	err := app.LogrusRedisHook(serviceName)
	if err != nil {
		panic(err)
	}

	DB, err := app.Gorm()
	if err != nil {
		panic(err)
	}

	handler := func(event broker.Event) error {
		orderNo := string(event.Message().Body)

		// 查询订单号是否存在
		order := models.UsersOrders{}
		res := DB.Model(models.UsersOrders{}).Where("platform_order_no = ?", orderNo).First(&order)
		if res.Error != nil {
			if res.RecordNotFound() {
				logrus.WithFields(logrus.Fields{
					"platformOrderNo": orderNo,
				}).Error("订单号不存在")
			} else {
				logrus.Error(res.Error)
			}
			return res.Error
		}

		// 查询订单明细
		var orderDetails []models.UsersOrdersDetails
		res = DB.Model(models.UsersOrdersDetails{}).Where("platform_order_no=?", orderNo).Find(&orderDetails)
		if res.Error != nil {
			logrus.WithFields(logrus.Fields{
				"platformOrderNo": orderNo,
				"desc":            "查询订单明细失败",
			}).Error(res.Error)
			return res.Error
		}

		// 读取加盟商信息
		franchiseeShop := models.FranchiseesShops{}
		res = DB.Model(models.Franchisees{}).Where("shop_code=?", order.ShopCode).First(&franchiseeShop)
		if res.Error != nil {
			if res.RecordNotFound() {
				logrus.WithFields(logrus.Fields{
					"platformOrderNo": order.PlatformOrderNo,
					"shopCode":        order.ShopCode,
				}).Info("查询商店编码失败，没有找到相应的加盟商信息")
				return nil
			} else {
				logrus.Error(res.Error)
			}
			return res.Error
		}

		franchiseeOrder := models.FranchiseesOrders{
			//Model:                     models.Model{},
			Status:                    nil,
			ShopCode:                  order.ShopCode,
			UserId:                    order.UserId,
			UserThumb:                 order.UserThumb,
			TotalImportPrice:          order.TotalImportPrice,
			TotalExportPrice:          order.TotalExportPrice,
			FranchiseeUsername:        franchiseeShop.FranchiseeUsername,
			FranchiseeCommission:      new(float64),
			PlatformOrderNo:           order.PlatformOrderNo,
			PlatformCommissionPercent: franchiseeShop.PlatformCommissionPercent,
			PlatformCommission:        new(float64),
		}
		var franchiseeOrderDetails []models.FranchiseesOrdersDetails

		for _, row := range orderDetails {
			// 计算商品佣金分配

			// 平台佣金百分比
			percent := decimal.NewFromInt(int64(*franchiseeShop.PlatformCommissionPercent)).Div(decimal.NewFromInt(100))
			// 商品件数
			piece := decimal.NewFromInt(int64(*row.GoodsPiece))
			// 总进、出，货价
			totalImportPrice := decimal.NewFromFloat(*row.GoodsImportPrice).Mul(piece)
			totalExportPrice := decimal.NewFromFloat(*row.GoodsExportPrice).Mul(piece)
			// 总盈利
			profit := totalExportPrice.Sub(totalImportPrice)

			// 平台佣金
			commission := profit.Mul(percent)
			pCommission, _ := commission.Float64()
			// 加盟商佣金
			commission = profit.Sub(commission)
			fCommission, _ := commission.Float64()

			// 订单信息
			*franchiseeOrder.FranchiseeCommission += fCommission
			*franchiseeOrder.PlatformCommission += pCommission

			// 订单明细
			detail := models.FranchiseesOrdersDetails{
				//Model:                     models.Model{},
				GoodsId:                   row.GoodsId,
				GoodsName:                 row.GoodsName,
				GoodsImage:                row.GoodsImage,
				GoodsImportPrice:          row.GoodsImportPrice,
				GoodsExportPrice:          row.GoodsExportPrice,
				GoodsExportPriceAdvise:    row.GoodsExportPriceAdvise,
				PlatformOrderNo:           row.PlatformOrderNo,
				PlatformCommissionPercent: franchiseeShop.PlatformCommissionPercent,
				PlatformCommission:        &pCommission,
				FranchiseeUsername:        franchiseeShop.FranchiseeUsername,
				FranchiseeCommission:      &fCommission,
			}
			franchiseeOrderDetails = append(franchiseeOrderDetails, detail)
		}

		// 事务
		{
			tx := DB.Begin()

			res := tx.Create(&franchiseeOrder)
			if res.Error != nil {
				goto txFailed
			}

			for _, row := range franchiseeOrderDetails {
				res := tx.Create(&row)
				if res.Error != nil {
					goto txFailed
				}
			}

			tx.Commit()
			event.Ack()
			return nil
		txFailed:
			tx.Rollback()
			logrus.WithFields(logrus.Fields{
				"platformOrderNo": franchiseeOrder.PlatformOrderNo,
				"desc":            "事务-插入订单信息失败",
			}).Error(res.Error)
			return res.Error
		}
	}

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
	service.Server().Options().Broker.Subscribe("orders.new", handler,
		broker.Queue("orders.new.franchisees"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
