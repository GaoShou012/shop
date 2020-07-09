package GORMDemo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:19841226Joe@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&User{})
}

func GORMDemo() {
	control := new(UserController)

	//fields := map[string]interface{}{"status":0}
	//
	//count, err := control.IsExists(fields)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("count = ", count)

	//fields := map[string]interface{}{"email":"wenxing.cn@gmail.com"}
	//user, err := control.Create(fields)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("user = ", user)


	//fields := map[string]interface{}{"email":"alwen"}
	//user, err := control.RetrieveOne(fields)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("user = ", user)
	//
	//temp := map[string]interface{}{"email": "wenxing.cn@gmail.com"}
	//user, err = control.Update(user, temp)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("user = ", user)

	//fields := map[string]interface{}{"email": "wenxing.cn@gmail.com"}
	//users, err := control.RetrieveAll(fields)
	//if err != nil {
	//	panic(err)
	//}
	//for k, v := range users {
	//	fmt.Println("user", k, " = ", v)
	//}

	fields := map[string]interface{}{"email": "wenxing.cn@gmail.com"}
	err := control.DeleteAll(fields)
	if err != nil {
		panic(err)
	}
}