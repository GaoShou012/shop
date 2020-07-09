package main

import (
	"fmt"
	"github.com/app"
	_ "github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"time"
)

/*
	from rabbitmq 得到用户的下单的订单号，然后给平台生成一份订单
*/

func main() {
	url := fmt.Sprintf("amqp://admin:admin@%s:%s/", app.Configs.RabbitMq.Host, app.Configs.RabbitMq.Port)
	rabbitmq.DefaultRabbitURL = url
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName(""),
	)
	if err := b.Init(); err != nil {
		panic(err)
	}
	if err := b.Connect(); err != nil {
		panic(err)
	}
	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		//micro.Name("micro.service.platform.orders"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
	)

	service.Server().Options().Broker.Subscribe("topic", func(event broker.Event) error {
		content := string(event.Message().Body)
		fmt.Println("content=", content)
		event.Ack()
		return nil
	}, broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
