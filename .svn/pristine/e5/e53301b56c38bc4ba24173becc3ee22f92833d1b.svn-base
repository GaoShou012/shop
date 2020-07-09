package main

import "fmt"

type A struct {
	Id uint64
}

type B struct {
	Id uint64
	Id1 uint64
}



func main() {
	var c struct{
		A
		B
	}
	c.A.Id = 123
	c.B.Id = 456
	fmt.Println(c)
}
