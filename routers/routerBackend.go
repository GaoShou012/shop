package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/middleware/cors"
	"github.com/routers/api"
	"github.com/routers/backend"
	"io/ioutil"
	"net/http"
)

func InitRouterForBackend() *gin.Engine {
	r := gin.New()
	r.Use(cors.CorsHandler())
	apiAuth := backend.ApiAuth{}

	r.GET("/weapp/login", func(ctx *gin.Context) {
		type Token struct {
			OpenId     string `json:"openid"`
			SessionKey string `json:"session_key"`
		}

		code := ctx.Request.URL.Query().Get("code")
		appId := "wxec46555d91948223"
		appSecret := "32fc5595749d3b65c62c7162fbb7c0c5"

		fmt.Println("get authSession")

		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, code)
		rsp, _ := http.Get(url)
		defer rsp.Body.Close()
		body, _ := ioutil.ReadAll(rsp.Body)
		t := Token{}
		json.Unmarshal(body, &t)
		ctx.JSON(http.StatusOK, gin.H{
			"code":        0,
			"message":     "ok",
			"openid":      t.OpenId,
			"session_key": t.SessionKey,
		})
	})

	apiMenu := backend.ApiMenus{}
	r.POST("/menu/all",apiMenu.All)

	apiAdmin := backend.ApiAdmins{}
	r.POST("/login", apiAdmin.Login)
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "testing",
		})
	})

	auth := r.Group("/auth")
	auth.Use(func(ctx *gin.Context) {
		if apiAuth.Verify(ctx) == false {
			return
		}
	})
	{
		auth.POST("/ping", apiAuth.Ping)
	}

	goods := r.Group("/goods")
	goods.Use(func(context *gin.Context) {
		// JWT 校验
		// 得到用户数据
		// 下面的接口就已经可以访问
		// 不然就 context.Abort()
	})
	{
		apiGoods := backend.Goods{}
		goods.POST("/insert", apiGoods.Insert)
		goods.POST("/delete", apiGoods.Delete)
		goods.POST("/update", apiGoods.Update)
		goods.POST("/search", apiGoods.Search)

		apiGoodsType := backend.ApiGoodsType{}
		goods.POST("/types/insert", apiGoodsType.Insert)
		goods.POST("/types/delete", apiGoodsType.Delete)
		goods.POST("/types/update", apiGoodsType.Update)
		goods.POST("/types/search", apiGoodsType.Search)
	}

	shops := r.Group("/shops")
	shops.Use(func(context *gin.Context) {
		// 校验用户是否已经验证
		//context.JSON(http.StatusOK,gin.H{})
		//context.Abort()
	})
	{
		apiShop := backend.ApiShop{}
		shops.POST("/insert", apiShop.Insert)
		shops.POST("/update", apiShop.Update)
		shops.POST("/search", apiShop.Search)

		// 商店的商品种类
		apiShopsGoodsTypes := backend.ApiShopsGoodsTypes{}
		shops.POST("/goods/types/insert", apiShopsGoodsTypes.Insert)
		shops.POST("/goods/types/update", apiShopsGoodsTypes.Update)
		shops.POST("/goods/types/search", apiShopsGoodsTypes.Search)

		// 商店的商品
		apiShopsGoods := backend.ApiShopsGoods{}
		shops.POST("/goods/copy", apiShopsGoods.Copy)
		shops.POST("/goods/update", apiShopsGoods.Update)
		shops.POST("/goods/stock/update", apiShopsGoods.UpdateStock)
		shops.POST("/goods/search", apiShopsGoods.Search)

		// 商店内，商品的分配
		apiShopsGoodsAssign := backend.ApiShopsGoodsAssign{}
		shops.POST("/goods/assign/insert", apiShopsGoodsAssign.Insert)
		shops.POST("/goods/assign/delete", apiShopsGoodsAssign.Delete)
		shops.POST("/goods/assign/update", apiShopsGoodsAssign.Update)
		shops.POST("/goods/assign/search", apiShopsGoodsAssign.Search)
	}

	franchisees := r.Group("/franchisees")
	franchisees.Use(func(ctx *gin.Context) {
		//if !apiAuth.Franchisees(ctx) {
		//	return
		//}
	})
	{
		apiFranchisees := backend.ApiFranchisees{}
		franchisees.POST("/insert", apiFranchisees.Insert)
		franchisees.POST("/update", apiFranchisees.Update)
		franchisees.POST("/search", apiFranchisees.Search)
		franchisees.POST("/bind/bankcard", apiFranchisees.BindBankCard)

		apiFranchiseesShops := backend.ApiFranchiseesShops{}
		franchisees.POST("/shop/bind", apiFranchiseesShops.Bind)
		franchisees.POST("/shop/unbind", apiFranchiseesShops.Unbind)
		franchisees.POST("/shop/update", apiFranchiseesShops.Update)
		franchisees.POST("/shop/search", apiFranchiseesShops.Search)
	}

	apiUsers := api.Users{}
	users := r.Group("/users")
	{
		users.POST("/order/token", apiUsers.OrderToken)
		users.POST("/order/order", apiUsers.Order)
		users.POST("/order/search", apiUsers.OrderSearch)
		users.POST("/order/details", apiUsers.OrderDetails)
		users.POST("/info", apiUsers.Info)
	}

	return r
}
