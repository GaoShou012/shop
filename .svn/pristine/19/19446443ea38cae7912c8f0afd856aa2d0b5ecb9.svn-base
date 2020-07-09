package main

import (
	"github.com/app"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-plugins/client/http"
)

func main() {
	/*
		正规姿势
		使用 go-micro plugins 进行请求
		github.com/micro/go-plugins
	*/
	service := micro.NewService(
		micro.Name("api.service.hello"),
		micro.Version("latest"),
		)
	service.Init()

}
