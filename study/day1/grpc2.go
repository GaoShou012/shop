package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/protoBuf"
	"time"
)

type Greeter2 struct{}

func (g *Greeter2) Val(ctx context.Context, req *protoBuf.ValRequest, rsp *protoBuf.ValResponse) error {
	rsp.Name = "Hello " + req.Name
	fmt.Println("received request")
	return nil
}

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("micro.greeter2"),
		//micro.Address(":9300"),
		micro.Registry(etcdReg),
		micro.RegisterTTL(time.Second*10),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	protoBuf.RegisterGreeter2Handler(service.Server(),new(Greeter2))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
