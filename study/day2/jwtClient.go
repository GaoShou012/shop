package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
)

func main() {
	service := micro.NewService(
		micro.Registry(app.EtcdRegistry()),
		)
	service.Init()

	jwtService := proto.NewJwtService("jwt",service.Client())

	rsp,err := jwtService.Encode(context.TODO(),&proto.JwtEncodeRequest{
		User:                 &proto.JwtUser{
			Id:                   123123,
			Username:             "aasdfasdf",
			Nickname:             "333333123",
			LoginTime:            11,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	rsp1,err := jwtService.Decode(context.TODO(),&proto.JwtDecodeRequest{
		Token:                rsp.Token,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp1.User)
}
