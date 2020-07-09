package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/proto"
	"io/ioutil"
	"net/http"
	"time"
)

// 获取验证码
// SessionId
func GetAuthCode(context *gin.Context) {

}

func Logout(context *gin.Context) {

}

type Auth struct {
}

func (api *Auth) WxLogin(ctx *gin.Context) {
	code := ctx.Request.URL.Query().Get("code")
	appId := ""
	appSecret := ""

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, code)
	rsp, _ := http.Get(url)
	defer rsp.Body.Close()

	body, _ := ioutil.ReadAll(rsp.Body)
	var info struct {
		OpenId     string `json:"openid"`
		SessionKey string `json:"session_key"`
	}
	err := json.Unmarshal(body, &info)
	if err != nil {
		// 输出日志
		return
	}

	token := ""
	{
		user := proto.JwtUser{
			Type:      "WeChat",
			Uuid:      info.OpenId,
			Id:        0,
			Username:  "",
			Nickname:  "",
			LoginTime: time.Now().Unix(),
		}
		jwt := proto.NewJwtService("micro.service.jwt", app.ServiceClient())
		rsp, err := jwt.Encode(context.TODO(), &proto.JwtEncodeRequest{
			User: &user,
		})
		if err != nil {
			app.ResponseError(ctx, err)
			return
		}
		token = rsp.Token
	}

	// 生成token
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"token":   token,
	})
}
