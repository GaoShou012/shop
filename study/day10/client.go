package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
	"time"
)

/*
	经过测试，并发调用服务
	是可以同时接待多请求
*/
func main() {
	service := micro.NewService(
		micro.Registry(app.EtcdRegistry()),
		)

	for i:=0;i<10;i++{
		go func(i int) {
			//service := micro.NewService(
			//	micro.Registry(app.EtcdRegistry()),
			//)
			cli := proto.NewTestService("micro.service.test",service.Client())
			rsp,err := cli.Ask(context.TODO(),&proto.TestReq{Message:fmt.Sprintf("id:%d %v",i,time.Now())})
			if err != nil {
				panic(err)
			}
			fmt.Println(rsp.Message,time.Now())
		}(i)
	}

	//go func() {
	//	//service := micro.NewService(
	//	//	micro.Registry(app.EtcdRegistry()),
	//	//	)
	//	cli := proto.NewTestService("micro.service.test",service.Client())
	//	rsp,err := cli.Ask(context.TODO(),&proto.TestReq{Message:"456"})
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(rsp.Message,time.Now())
	//}()

	time.Sleep(time.Second*10)
}
