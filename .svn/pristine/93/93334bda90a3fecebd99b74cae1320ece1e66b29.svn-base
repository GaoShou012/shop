package main

import (
	"github.com/micro/go-micro/v2"
)

func main() {
	//etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	service := micro.NewService(
		micro.Name("micro.servicea"),
		//micro.Registry(etcdReg),
		)
	service.Init()
}
