package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2/broker"

	//"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func simplepub() {
	rabbitmq.DefaultRabbitURL = app.RabbitMqUrl()
	b := rabbitmq.NewBroker(
		rabbitmq.ExchangeName("broker"),
		rabbitmq.DurableExchange(),
	)
	b.Init()
	b.Connect()
	go func() {
		for i :=0;i< 100; i++ {
			b.Publish("topic",&broker.Message{
				Header: nil,
				Body:   []byte(fmt.Sprintf("id%v",i)),
			},rabbitmq.DeliveryMode(2))
			fmt.Printf("pub:%v\n",i)
			//time.Sleep(time.Second*1)
		}
	}()
	for{}
}

func main() {
	simplepub()
}
