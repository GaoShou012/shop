package backend

import (
	"fmt"
	"github.com/app"
	"github.com/gin-gonic/gin"
	"github.com/models"
)

type ApiFranchisees struct{}

func (api *ApiFranchisees) Insert(ctx *gin.Context) {
	var params struct {
		Username string
		Password string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	passwordHash := app.PasswordHash(params.Password)
	balance := float32(0)
	withdrawalTotalAmount := float32(0)

	newData := models.Franchisees{
		//Model:                 models.Model{},
		//Status:                nil,
		Username:              &params.Username,
		Password:              &passwordHash,
		Balance:               &balance,
		WithdrawalTotalAmount: &withdrawalTotalAmount,
		//WithdrawalTotalTimes:  nil,
	}
	res := models.DB.Create(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "创建加盟商账号成功")
}

func (api *ApiFranchisees) Update(ctx *gin.Context) {
	var params struct {
		Id       uint64
		Status   uint64
		Password string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	newData := models.Franchisees{
		//Model:                 models.Model{},
		Status: &params.Status,
		//Username:              nil,
		//Password:              nil,
		//Balance:               nil,
		//WithdrawalTotalAmount: nil,
		//WithdrawalTotalTimes:  nil,
	}

	// 如果密码不等于 空字符串 就更新密码
	if params.Password != "" {
		password := app.PasswordHash(params.Password)
		newData.Password = &password
	}

	res := models.DB.Model(models.Franchisees{}).Where("id = ?", params.Id).Updates(&newData)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	if res.RowsAffected != 1 {
		app.ResponseError(ctx, fmt.Errorf("更新失败"))
		return
	}

	app.ResponseSuccess(ctx, "更新成功")
}

func (api *ApiFranchisees) Search(ctx *gin.Context) {
	var params struct {
		Id       uint64
		Username string
	}

	var total uint64 = 0
	var values []models.Franchisees
	db := models.DB.Model(models.Franchisees{})
	if params.Id > 0 {
		db = db.Where("id = ?", params.Id)
	}
	if params.Username != "" {
		db = db.Where("username = ?", params.Username)
	}
	db.Count(&total)
	res := db.Find(&values)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}
	app.ResponseSearchSuccess(ctx, total, values)
}

/*
	绑定银行卡
*/
func (api *ApiFranchisees) BindBankCard(ctx *gin.Context) {
	var params struct {
		Username   string
		RealName   string
		BankName   string
		BankBranch string
		CardNo     string
		Phone      string
	}

	err := ctx.BindJSON(&params)
	if err != nil {
		app.ResponseError(ctx, err)
		return
	}

	bankCard := models.FranchiseesBankCard{
		Model:      models.Model{},
		Username:   &params.Username,
		RealName:   &params.RealName,
		BankName:   &params.BankName,
		BankBranch: &params.BankBranch,
		CardNo:     &params.CardNo,
		Phone:      &params.Phone,
	}
	res := models.DB.Debug().Create(&bankCard)
	if res.Error != nil {
		app.ResponseError(ctx, res.Error)
		return
	}

	app.ResponseSuccess(ctx, "绑定银行卡成功")
}
