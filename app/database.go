package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func Gorm() (*gorm.DB,error) {
	// 连接数据库
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		Configs.DB.User,
		Configs.DB.Password,
		Configs.DB.Host,
		Configs.DB.Port,
		Configs.DB.DefaultDatabase,
	)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil,err
	}
	return db,nil
}