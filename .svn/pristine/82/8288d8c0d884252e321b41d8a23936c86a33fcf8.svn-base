package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	d,err := time.ParseDuration("-24h")
	if err != nil {
		panic(err)
	}
	//fmt.Println(d)
	fmt.Println(now.Add(d*1).Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 ") + "00:00:00")
	fmt.Println(now.Format("2006-01-02 ") + "23:59:59")
}
