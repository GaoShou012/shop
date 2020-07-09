package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/broker"
)


func main() {
	if err := broker.Connect(); err != nil{
		panic(err)
	}

	sub,err := broker.Subscribe("topic", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		return nil
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	for{}
}

