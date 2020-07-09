package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func ParseIdFromPost(context *gin.Context,key string) (id uint,err error) {
	str := context.DefaultPostForm(key,"0")
	id64,err := strconv.ParseUint(str,10,32)
	if err != nil {
		return
	}
	id = uint(id64)
	return
}

func ParseTimeFromPost(context *gin.Context,key string) (t *time.Time,err error) {
	return
}

func ParseUintFromPost(context *gin.Context,key string) (num uint64,err error) {
	str := context.DefaultPostForm(key,"0")
	i64,err := strconv.ParseUint(str,10,32)
	if err != nil {
		err = fmt.Errorf("%s %s",key,err.Error())
		return
	}
	num = i64
	return
}

func ParsePriceFromPost(context *gin.Context,key string) (price float32,err error) {
	str := context.PostForm(key)
	f64,err := strconv.ParseFloat(str,32)
	if err != nil {
		return
	}
	price = float32(f64)
	return
}

func ParsePageFroPost(context *gin.Context) (page uint,err error) {
	str := context.DefaultPostForm("page","0")
	i64,err := strconv.ParseUint(str,10,32)
	if err != nil {
		return
	}
	page = uint(i64)
	if page > 0 {page = page -1}
	return
}

func ParseIdListFromPost(context *gin.Context,key string) (li []uint,err error) {
	str := context.PostForm(key)
	strList := strings.Split(str,",")
	if len(strList) < 1 || len(strList) > 64 {
		err = fmt.Errorf("ID数量不在范围内")
		return
	}

	for _,v := range strList {
		id,err := strconv.ParseUint(v,10,32)
		if err != nil {
			return nil,err
		}
		li = append(li,uint(id))
	}

	if len(strList) != len(li) {
		err = fmt.Errorf("解析ID列表，未知错误")
		return
	}

	return
}