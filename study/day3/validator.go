package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string
	Password string			`validate:"required,min=6,max=18" vmsg:"用户密码必须是6-18位"`
}


func main() {

	u := &User{
		Username: "213",
		Password: "333",
	}

	valid := validator.New()
	err := valid.Struct(u)
	if err != nil {
		fmt.Println(err)

		// 处理错误信息
		if errs,ok := err.(validator.ValidationErrors); ok {
			for _,e := range errs {
				fmt.Println(e.Value())
				fmt.Println(e.Field())
				fmt.Println(e.Tag())
			}
		}

		return
	}

	fmt.Println("finished")
}
