package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/protoBuf"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	// Create a new service
	service := micro.NewService(
		//micro.Name("micro.greeter.client"),
		micro.Registry(etcdReg),
		//micro.Address(":9300"),
		//micro.Address(":9300"),
		)
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeter := protoBuf.NewGreeterService("micro.greeter1", service.Client())

	// Create new greeter2 client
	greeter2 := protoBuf.NewGreeter2Service("micro.greeter2",service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &protoBuf.HelloRequest{Name:"John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print response
	fmt.Println(rsp.Greeting)

	// Call the greeter2
	rsp1,err := greeter2.Val(context.TODO(),&protoBuf.ValRequest{Name:"123123"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp1.Name)
}
