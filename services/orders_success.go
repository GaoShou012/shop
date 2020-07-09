package main

import (
	"context"
	"github.com/app"
	_ "github.com/app"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"github.com/models"
	"github.com/proto"
	"time"
)

type handler struct {
	DB     *gorm.DB
	Broker *broker.Broker
}

/*
	code = 2
	订单不存在
*/
func (h handler) Order(ctx context.Context, req *proto.UsersOrdersSuccessReq, rsp *proto.UsersOrdersSuccessRsp) error {
	order := models.UsersOrders{}

	// 查询订单当前状态
	res := h.DB.Model(models.UsersOrders{}).Where("platform_order_no=?", req.PlatformOrderNo).First(&order)
	if res.Error != nil {
		if res.RecordNotFound() {
			rsp.Code = 2
			rsp.Message = "订单不存在"
			return nil
		} else {
			return res.Error
		}
	}

	// 检查订单状态
	switch *order.Status {
	case 0:
		break
	case 1:
		rsp.Code = 3
		rsp.Message = "订单已经标记完成"
		return nil
	case 2:
		rsp.Code = 4
		rsp.Message = "订单已经超时，不能标记完成"
		return nil
	default:
		rsp.Code = 1
		rsp.Message = "订单未知状态"
		return nil
	}

	status := uint64(1)
	now := time.Now()
	res = h.DB.Model(models.UsersOrders{}).Where("platform_order_no=?", req.PlatformOrderNo).Updates(&models.UsersOrders{
		//Model:            models.Model{},
		PlatformOrderNo:  nil,
		UserId:           nil,
		UserThumb:        nil,
		Status:           &status,
		TotalImportPrice: nil,
		TotalExportPrice: nil,
		ShopCode:         nil,
		ShopAddr:         nil,
		ClientToken:      nil,
		PaidAt:           &now,
	})
	if res.Error != nil {
		return res.Error
	}

	rsp.Code = 0
	rsp.Message = "订单标记成功"
	return nil
}

func main() {
	serviceName := "micro.service.users.orders.success"

	err := app.LogrusRedisHook(serviceName)
	if err != nil {
		panic(err)
	}

	DB, err := app.Gorm()
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
	serviceHandler := handler{DB: DB, Broker: &b}
	err = proto.RegisterUsersOrdersSuccessHandler(service.Server(), &serviceHandler)
	if err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		panic(err)
	}
}
