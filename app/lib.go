package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetPostUint(context *gin.Context) {

}

type PageParams struct {
	Page 			uint64			`json:"page"`
	PageSize		uint64			`json:"pageSize"`
}

func (p *PageParams) CheckPageParams() error {
	// 前端显示页数最少值为1，这里可以进行减少1，配合前端操作
	if p.Page > 0 { p.Page -= 1 }
	// 页数大小默认为20
	if p.PageSize == 0 { p.PageSize = 20 }

	if p.Page > 1024 {
		return fmt.Errorf("搜索的页数超过限制")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return fmt.Errorf("页大小，超过限制")
	}

	return nil
}

func (p *PageParams) Offset() uint64 {
	return p.Page * p.PageSize
}
