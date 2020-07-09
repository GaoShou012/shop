package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"reflect"
	"time"
)

var DB *gorm.DB

//type Error struct{
//	Code int
//	Message string
//}

//func (e *Error) Error() string {
//	return fmt.Printf("Error %i:%s",e.Code,e.Message)
//}

type Model struct {
	ID        *uint64    `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"default:NULL" sql:"index"`
}

func RecoverRecord(model interface{},val interface{},values ...interface{}) error {
	obj := reflect.ValueOf(model)
	v := obj.MethodByName("Table")
	if !v.IsValid() {
		panic("TableName Method is not exists!")
	}

	table := v.Call(nil)[0].String()

	switch val.(type) {
	case string:
		res := DB.Table(table).Unscoped().Where(val,values...).Update("deleted_at = NULL")
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected <= 0 {
			return errors.New("恢复记录失败")
		}
		break
	case int:
	case int32:
	case int64:
	case uint:
	case uint32:
	case uint64:
		res := DB.Table(table).Unscoped().Where(val,values...).Update("deleted_at = NULL")
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected <= 0 {
			return errors.New("恢复记录失败")
		}
		break
	default:
		return errors.New("IsExists 未知数据类型")
	}

	return nil
}

func IsExistsV1(DB *gorm.DB,model interface{}, unscoped bool, val interface{}, values ...interface{}) (exists bool, err error) {
	obj := reflect.ValueOf(model)
	v := obj.MethodByName("Table")
	if !v.IsValid() {
		panic("Table Method is not exists!")
	}

	table := v.Call(nil)[0].String()
	count := 0

	switch val.(type) {
	case string:
		db := DB.Table(table)
		if unscoped {
			db = db.Unscoped()
		}
		res := db.Where(val, values...).Count(&count)
		if res.Error != nil {
			err = res.Error
			return
		}
		break
	case int:
	case int32:
	case int64:
	case uint:
	case uint32:
	case uint64:
		db := DB.Table(table)
		if unscoped {
			db = db.Unscoped()
		}
		res := db.Where("`id` = ?", val).Count(&count)
		if res.Error != nil {
			err = res.Error
			return
		}
		break
	default:
		err = errors.New("IsExists 未知数据类型")
		return
	}

	if count > 0 {
		exists = true
	}

	return
}

/*
	unscoped = true
	表示可以查询软删除的记录

	val 如果是数字，直接查询 id = ?
	val 如果是字符串，充当query来拼接 values 充当 查询的值
*/
func IsExists(model interface{}, unscoped bool, val interface{}, values ...interface{}) (exists bool, err error) {
	obj := reflect.ValueOf(model)
	v := obj.MethodByName("Table")
	if !v.IsValid() {
		panic("TableName Method is not exists!")
	}

	table := v.Call(nil)[0].String()
	count := 0

	switch val.(type) {
	case string:
		db := DB.Table(table)
		if unscoped {
			db = db.Unscoped()
		}
		res := db.Where(val, values...).Count(&count)
		if res.Error != nil {
			err = res.Error
			return
		}
		break
	case int:
	case int32:
	case int64:
	case uint:
	case uint32:
	case uint64:
		db := DB.Table(table)
		if unscoped {
			db = db.Unscoped()
		}
		res := db.Where("`id` = ?", val).Count(&count)
		if res.Error != nil {
			err = res.Error
			return
		}
		break
	default:
		err = errors.New("IsExists 未知数据类型")
		return
	}

	fmt.Printf("count=%d\n", count)
	if count > 0 {
		exists = true
	}

	return
}



//type Model struct {
//	Stt interface{}
//	TableName string
//}
//
//func (m *Model) DeleteById(id int) error {
//	var stt interface{}
//	res := DB.Table(m.TableName).Where("id=?",id).Delete(&stt)
//	return res.Error
//}
//
//func (m *Model) UpdateById(id int,values map[string]interface{}) error {
//	values["update_at"] = time.Now()
//	res := DB.Table(m.TableName).Where("id=?",id).Update(values)
//	return res.Error
//}
//
//func (m *Model) SelectById(id int,v interface{}) error {
//	res := DB.Table(m.TableName).Where("id=?",id).Find(v)
//	if res.RecordNotFound() {
//		return fmt.Errorf("查询的ID(%d)不存在",id)
//	}
//	return nil
//}
//
//func GetErrorCode(err error) (int,error) {
//	code,err1 := strconv.Atoi(err.Error()[6:10])
//	if err1 != nil {
//		return 0,fmt.Errorf("获取错误编码失败:%s",err.Error())
//	}
//	return code,nil
//}

func init() {
	var err error
	dbType := "mysql"
	user := "root"
	password := "Forever634312."
	host := "192.168.0.200"
	dbName := "server"

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user,
		password,
		host,
		dbName,
	))

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	DB = db
}
