package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"time"

	//"github.com/micro/go-plugins/broker/kafka"

	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
	//_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	//_ "github.com/micro/go-plugins/transport/nats/v2"



	//"github.com/micro/go-plugins/transport/rabbitmq"
)


func main() {
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
		//micro.Name("greeter"),
		micro.Server(s),
		micro.Broker(b),
		micro.Registry(registry),
		//micro.Transport(transport),
	)

	go func() {
		b := service.Server().Options().Broker
		for i :=0;i< 100; i++ {
			b.Publish("topic",&broker.Message{
				Header: nil,
				Body:   []byte(fmt.Sprintf("id%v",i)),
			})
			time.Sleep(time.Second*1)
		}
	}()

	service.Init()
	service.Run()
}

