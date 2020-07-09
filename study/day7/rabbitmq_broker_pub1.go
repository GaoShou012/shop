package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/proto"
	//"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
	)
	b.Init()
	b.Connect()

	// new service
	service := micro.NewService(
		micro.Name("micro.service.broker"),
		micro.Version("latest"),
		micro.Broker(b),
		micro.Registry(app.EtcdRegistry()),
	)
	s := proto.NewBrokerMsgService("micro.service.broker",service.Client())
	rsp,err := s.Order(context.TODO(),&proto.BrokerMsgReq{Message:"i am message 123"})
	if err != nil {
		panic(err)
	}
	fmt.Println("i am rsp ",rsp.Message)
}
