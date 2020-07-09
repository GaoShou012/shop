package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

func handler(srv server.Server) {
	fmt.Println(srv)
}

func main() {
	service := micro.NewService(
		micro.Name("micro.servicea"),
		micro.Version("last"),
		)

	service.Init()
	handler(service.Server())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}