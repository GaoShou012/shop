package main

import (
	"github.com/app"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("micro.service.transport"),
		micro.Version("latest"),
		micro.Registry(app.EtcdRegistry()),
		)
}
