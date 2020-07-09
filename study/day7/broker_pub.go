package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/broker"
	"time"
)

func main() {
	if err := broker.Connect(); err != nil{
		panic(err)
	}
	for i :=0;i< 10; i++ {
		broker.Publish("topic",&broker.Message{
			Header: nil,
			Body:   []byte(fmt.Sprintf("id%v",i)),
		})
		fmt.Printf("i am sending:%v\n",i)
		time.Sleep(time.Second*1)
	}
	for{}
}

