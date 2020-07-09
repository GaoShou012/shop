package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"time"
)

func simplesub() {
	rabbitmq.DefaultRabbitURL = app.RabbitMqUrl()
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName("orders"),
	)
	b.Init()
	b.Connect()
	b.Subscribe("topic", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		return nil
	})
	for {
	}
}

func main() {
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
		rabbitmq.ExchangeName("broker"),
	)
	b.Init()
	b.Connect()
	b.Subscribe("topic", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		event.Ack()
		time.Sleep(time.Second*1)
		return nil
	},
		broker.Queue("broker.orders"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)
	//for{}
	// 以上代码已经完成sub的所有代码，下面加入服务监控
	service := micro.NewService(
		micro.Name("subscribe"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
	)
	service.Init()
	service.Run()
}
