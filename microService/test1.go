package microService

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api"
	"github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
)

// protoc --go_out=. --micro_out=. echo.proto
// micro --registry etcd --registry_address 127.0.0.1:2379 list services
// micro --registry etcd --registry_address 127.0.0.1:2379 get service api.Echo.com.entry
// micro --registry etcd --registry_address 127.0.0.1:2379 call api.Echo.com.entry EchoService.Echo "{\"Input\":\"alwen\"}"

/* API Gateway (micro api)
Example:
         Path         Service                Method
        /foo/bar       go.micro.api.foo      Foo.Bar
        /foo/bar/baz    go.micro.api.foo      Bar.Baz
        /foo/bar/baz/cat  go.micro.api.foo.bar   Baz.Cat

查环境变量: micro api -h

MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 MICRO_API_NAMESPACE=api.Echo.com MICRO_API_HANDLER=rpc MICRO_API_ADDRESS=0.0.0.0:8888 micro api

MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 MICRO_API_NAMESPACE=micro.grpc.echo MICRO_API_HANDLER=rpc MICRO_API_ADDRESS=0.0.0.0:8888 micro api


Postman:
         POST http://127.0.0.1:8888/entry/EchoService/echo
         {"Input":"alwen"}

*/

type EchoServer struct{}
func (e *EchoServer) Echo(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Output = fmt.Sprintf("%v", req.Input)
	return nil
}

func GRPCEchoServiceETCDEntry1() {
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	service := micro.NewService(
		micro.Name("micro.grpc.echo.entry"),
		micro.Address(":8091"),
		micro.Registry(etcdRegistry),
		micro.RegisterTTL(time.Second*20),
	)
	err := RegisterEchoServiceHandler(service.Server(), new(EchoServer), api.WithEndpoint(&api.Endpoint{
		// The RPC method
		Name: "EchoService.Echo",
		// The HTTP paths. This can be a POSIX regex
		Path: []string{"/echo"},
		// The HTTP Methods for this endpoint
		Method: []string{"POST"},
		// The API handler to use
		Handler: rpc.Handler,
	}))
	if err != nil {
		panic(err)
	}
	service.Init()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}