package main

import "fmt"

type User struct {
	Username string
	Password string
	Age int
}

func main() {
	var p *User
	fmt.Println(p)  // 输出是 nil

	p1 := User{}
	fmt.Println(p1) // 输出{  0} 表示已经实例化

	p2 := &User{}
	fmt.Println(p2) // 输出 &{  0} 表示已经实例化
}
