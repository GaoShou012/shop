package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
	"log"
	"time"
)

func main() {
	service := micro.NewService(micro.Registry(app.EtcdRegistry()))
	service.Init()

	go func() {
		for i:=0;i<10;i++{
			orderNoService := proto.NewOrderNoService("order.no",service.Client())
			rsp,err := orderNoService.Gen(context.TODO(),&proto.OrderNoRequest{})
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("i am 1:%v\n",rsp.OrderNo)
		}
	}()
	go func() {
		for i:=0;i<10;i++{
			orderNoService := proto.NewOrderNoService("order.no",service.Client())
			rsp,err := orderNoService.Gen(context.TODO(),&proto.OrderNoRequest{})
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("i am 2:%v\n",rsp.OrderNo)
		}
	}()
	go func() {
		for i:=0;i<10;i++{
			orderNoService := proto.NewOrderNoService("order.no",service.Client())
			rsp,err := orderNoService.Gen(context.TODO(),&proto.OrderNoRequest{})
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("i am 3:%v\n",rsp.OrderNo)
		}
	}()

	time.Sleep(time.Second*5)
}
