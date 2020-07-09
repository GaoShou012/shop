package main

import (
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
		rabbitmq.ExchangeName("orders"),
		rabbitmq.DurableExchange(),
	)
	b.Init()
	b.Connect()
	b.Publish("orders.new", &broker.Message{
		Header: nil,
		Body:   []byte(fmt.Sprintf("adfasdf")),
	}, rabbitmq.DeliveryMode(2))
}
