package main

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"net/http"
	"time"
)

func main() {
	// 初始化路由
	router := gin.Default()
	router.Handle("GET","/", func(ctx *gin.Context) {
		fmt.Println("i am hello")
		ctx.JSON(http.StatusOK,gin.H{
			"code":0,
			"message":"hello!",
		})
	})

	service := web.NewService(
		web.Name("api.service.hello"),
		web.Address(":8013"),
		web.Handler(router),
		web.Registry(app.EtcdRegistry()),
		web.RegisterTTL(time.Second*10),
		)

	service.Init()
	err := service.Init()
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
