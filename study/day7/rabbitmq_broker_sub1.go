package main

import (
	"context"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"

	"github.com/micro/go-plugins/transport/nats/v2"
	"github.com/proto"

	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
)

type notification struct {
}

func (s *notification) Order(ctx context.Context, req *proto.BrokerMsgReq, rsp *proto.BrokerMsgRsp) error {
	rsp.Message = "content:" + req.Message
	return nil
}

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
		micro.Transport(nats.NewTransport()),
	)

	service.Init()

	proto.RegisterBrokerMsgHandler(service.Server(),new(notification))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
