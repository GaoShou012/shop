package franchisees

import (
	"crypto/md5"
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiAuth struct {

}

func (api *ApiAuth) Login(ctx *gin.Context) {
	var params struct{
		Username string
		Password string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx,err)
		return
	}

	hash := md5.New()
	hash.Write([]byte(params.Password))
	passwordHash := hash.Sum(nil)

	userData := models.Franchisees{}

	db := models.DB.Table(models.FranchiseesTableName)
	res := db.Where("username = ? and password = ?",params.Username,passwordHash).First(&userData)
	if res.Error != nil {
		if res.RecordNotFound() {
			app.ResponseError(ctx,fmt.Errorf("账号或密码错误"))
		}else{
			app.ResponseError(ctx,res.Error)
		}
		return
	}
}

