package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
	"time"
)

type TestHandler struct {}

func (h *TestHandler) Ask(ctx context.Context,req *proto.TestReq,rsp *proto.TestRsp) error {
	fmt.Println("content=",req.Message)
	time.Sleep(time.Second*1)
	rsp.Message = req.Message
	return nil
}



func main() {
	service := micro.NewService(
		micro.Name("micro.service.test"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
		)

	err := proto.RegisterTestHandler(service.Server(),&TestHandler{})
	if err != nil {
		panic(err)
	}

	service.Init()
	service.Run()
}
