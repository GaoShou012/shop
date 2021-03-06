package app

import (
	"github.com/gin-gonic/gin"
	"github.com/proto"
)

type OperatorInfo struct {
	Id        *uint64 `json:"id"`
	Username  *string `json:"username"`
	Nickname  *string `json:"nickname"`
	LoginTime *int64  `json:"loginTime"`
}

func Operator(ctx *gin.Context) *OperatorInfo {
	data, ok := ctx.Get("operator")
	if !ok {
		panic("获取用户信息失败")
	}
	info := data.(*proto.JwtUser)
	id := uint64(info.Id)
	return &OperatorInfo{
		Id:        &id,
		Username:  &info.Username,
		Nickname:  &info.Nickname,
		LoginTime: &info.LoginTime,
	}
}
