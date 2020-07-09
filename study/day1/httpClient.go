package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	servers,err := etcdReg.GetService("api.service.123")
	fmt.Println(servers)

	if err != nil {
		fmt.Println(err)
		return
	}

	var services []*registry.Service
	for _,value := range servers {
		fmt.Println(value.Name, ":", value.Version)
		fmt.Println(value.Nodes)
		for _,v := range value.Nodes {
			fmt.Println(v.Address)
		}
		services = append(services, value)
	}

	fmt.Println(services)
}
