package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2"
)

func test() {
	registry := app.EtcdRegistry()

	//url := fmt.Sprintf("amqp://admin:admin@192.168.0.200:5672/")
	//rabbitmq.DefaultRabbitURL = url
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName("exchangeName"),
	)
	b.Init()
	b.Connect()
	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Registry(registry),
	)

	service.Server().Options().Broker.Subscribe("topic", func(p broker.Event) error {
		content := string(p.Message().Body)
		fmt.Println("content = ", content)
		p.Ack()
		return nil
	}, broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
		//rabbitmq.AckOnSuccess(),
	)

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func ok1() {
	registry := app.EtcdRegistry()

	url := fmt.Sprintf("amqp://admin:admin@192.168.0.200:5672/")
	rabbitmq.DefaultRabbitURL = url
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName("exchangeName"),
	)
	b.Init()
	b.Connect()
	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
		micro.Registry(registry),
	)

	service.Server().Options().Broker.Subscribe("topic", func(p broker.Event) error {
		content := string(p.Message().Body)
		fmt.Println("content = ", content)
		p.Ack()
		return nil
	}, broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
		//rabbitmq.AckOnSuccess(),
	)

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func main() {
	test()
}
