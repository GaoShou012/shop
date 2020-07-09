package backend

import (
	"context"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
	"net/http"
)

type ApiAuth struct {
}

func (api *ApiAuth) Franchisees(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    255,
			"message": "用户信息已经失效",
		})
		ctx.Abort()
		return false
	}

	service := micro.NewService(micro.Registry(app.EtcdRegistry()))
	service.Init()

	jwt := proto.NewJwtService("jwt", service.Client())
	rsp, err := jwt.Decode(context.TODO(), &proto.JwtDecodeRequest{Token: token})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		ctx.Abort()
		return false
	}

	ctx.Set("operator", rsp.User)
	return true
}

func (api *ApiAuth) Verify(ctx *gin.Context) bool {
	// 检查token是否存在
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    255,
			"message": "用户信息已经失效",
		})
		ctx.Abort()
		return false
	}

	// 解压token信息
	service := micro.NewService(micro.Registry(app.EtcdRegistry()))
	service.Init()

	jwt := proto.NewJwtService("micro.service.jwt", service.Client())
	rsp, err := jwt.Decode(context.TODO(), &proto.JwtDecodeRequest{Token: token})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		ctx.Abort()
		return false
	}

	ctx.Set("operator", rsp.User)
	return true
}

func (api *ApiAuth) Ping(ctx *gin.Context) {
	operator := app.Operator(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"id":        operator.Id,
		"username":  operator.Username,
		"nickname":  operator.Nickname,
		"loginTime": operator.LoginTime,
	})
}
