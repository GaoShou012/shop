package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"time"
)

func main() {
	// 启动broker消费端
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
		rabbitmq.ExchangeName("orders"),
	)
	b.Init()
	b.Connect()
	b.Subscribe("orders.new", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		event.Ack()
		time.Sleep(time.Second * 1)
		return nil
	},
		broker.Queue("orders.new.franchisees"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	// 启动服务
	service := micro.NewService(
		micro.Name("orders.new.franchisees"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
	)
	service.Init()
	service.Run()
}
