package main

import (
	"fmt"
	"github.com/app"
)

func main() {
	services,err := app.EtcdRegistry().GetService("api.service.hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(services)

	for _,v := range services {
		fmt.Println(v.Name,":",v.Version)
		fmt.Println(v.Nodes)
		for _,node := range v.Nodes {
			fmt.Println(node.Address)
		}
	}

}
