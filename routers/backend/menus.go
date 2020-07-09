package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiMenus struct {
}

func (api *ApiMenus) All(ctx *gin.Context) {
	type Item struct {
		Title string `json:"title"`
		Index string `json:"index"`
		Icon  string `json:"icon"`
	}

	type sub struct {
		Title string `json:"title"`
		Index string `json:"index"`
		Route string `json:"route"`
		Meta  string `json:"meta"`
		Icon  string `json:"icon"`
	}

	type menu struct {
		Item
		Subs []sub `json:"subs"`
	}

	var list []menu

	{
		var row menu
		var subs []sub
		row = menu{
			Item: Item{Title: "商品", Index: "goods", Icon: "el-icon-lx-warn"},
			Subs: nil,
		}
		subs = append(subs, sub{
			Title: "商品列表",
			Index: "goodsGoods",
			Icon:  "",
		})
		subs = append(subs, sub{
			Title: "商品种类",
			Index: "goodsGoodsTypes",
			Icon:  "",
		})

		row.Subs = subs
		list = append(list, row)
	}
	{
		var row menu
		var subs []sub
		row = menu{
			Item: Item{Title: "商店", Index: "shops", Icon: "el-icon-lx-warn"},
			Subs: nil,
		}
		subs = append(subs,sub{
			Title: "商店管理",
			Index: "shopsShops",
			Route: "",
			Meta:  "",
			Icon:  "",
		})
		subs = append(subs,sub{
			Title: "种类栏目",
			Index: "shopsGoodsTypes",
			Route: "",
			Meta:  "",
			Icon:  "",
		})
		subs = append(subs,sub{
			Title: "栏目列表",
			Index: "shopsGoods",
			Route: "",
			Meta:  "",
			Icon:  "",
		})
		subs = append(subs,sub{
			Title: "商品分配",
			Index: "shopsGoodsAssign",
			Route: "",
			Meta:  "",
			Icon:  "",
		})

		row.Subs = subs
		list = append(list, row)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"list":    list,
	})
}
