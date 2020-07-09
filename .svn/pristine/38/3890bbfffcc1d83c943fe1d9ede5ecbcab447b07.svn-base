package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)


func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	//r := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"), etcd.Auth("root", "123456"))

	// 初始化路由
	router := gin.Default()

	s := web.NewService(
		web.Name("api.service.123"),
		//web.Id("api.service.123"),
		web.Address(":8123"),
		web.Handler(router),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*20),
		//web.Metadata(map[string]string{
		//	"protocol": "http",
		//}),
	)



	//service := micro.NewService(micro.Name("api.service"))
	//microService := micro.NewService(micro.Name("api.service"), micro.WrapClient(roundrobin.NewClientWrapper()))
	//productService := NewQueryService("productQueryService", service.Client())

	//router.Use(injectParameters(productService))

	//router.Handle("POST", "/productList", postHandler)

	err := s.Init()
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}

