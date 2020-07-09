package backend

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"github.com/proto"
	"net/http"
	"time"
)

type ApiAdmins struct{}

func (api *ApiAdmins) Login(ctx *gin.Context) {
	var params struct {
		Username string
		Password string
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	fmt.Println("i am login")

	// 数据库校验密码
	passwordHash := app.PasswordHash(params.Password)
	user := models.Admin{}
	res := models.DB.Model(&models.Admin{}).Where("username=? and password=?", params.Username, passwordHash).First(&user)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	// JWT 生成Token
	jwt := proto.NewJwtService("micro.service.jwt", app.ServiceClient())
	rsp, err := jwt.Encode(context.TODO(), &proto.JwtEncodeRequest{
		User: &proto.JwtUser{
			Id:        *user.ID,
			Username:  *user.Username,
			Nickname:  *user.Nickname,
			LoginTime: time.Now().Unix(),
		},
	})
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"token":   rsp.Token,
	})
}
