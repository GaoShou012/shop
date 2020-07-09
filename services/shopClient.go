package main

import (
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	registry := app.EtcdRegistry()

	// broker
	rabbitmq.DefaultRabbitURL = app.RabbitMqUrl()
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName("shops"),
		)
	if err := b.Init(); err != nil {
		panic(err)
	}
	if err := b.Connect(); err != nil {
		panic(err)
	}

	service := micro.NewService(
		micro.Name("micro.service.shops"),
		micro.Broker(b),
		micro.Registry(registry),
		)
	service.Init()
	
	if err := service.Run(); err != nil {
		panic(err)
	}
}
