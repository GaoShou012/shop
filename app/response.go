package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicesErr struct {
	Id string
	Code uint64
	Detail string
	Status string
}

func ParseServicesErr(err error,v interface{}) error {
	if json.Unmarshal([]byte(err.Error()),v) != nil {
		return fmt.Errorf("解析服务error失败")
	}
	return nil
}

func Response(ctx *gin.Context,code uint64,message string) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":message,
	})
}

func ResponseError(context *gin.Context,err error) {
	if context == nil {
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK,gin.H{
		"code":1,
		"message":err.Error(),
	})
}

func ResponseSuccess(context *gin.Context,message string) {
	context.JSON(http.StatusOK,gin.H{
		"code":0,
		"message":message,
	})
}

func ResponseSearchSuccess(context *gin.Context,total interface{},values interface{}) {
	context.JSON(http.StatusOK,gin.H{
		"code":0,
		"message":"查询成功",
		"total":total,
		"list":values,
	})
}

func ResponseMicroServicesError(ctx *gin.Context,err error) {
	msErr := ServicesErr{}

	if json.Unmarshal([]byte(err.Error()),&msErr) != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code":500,
			"message":"解析服务error失败",
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":1,
		"message":msErr.Detail,
	})
}