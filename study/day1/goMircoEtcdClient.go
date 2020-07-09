package main

import (
	"fmt"
	microconfig "github.com/micro/go-micro/v2/config"
	microetcd "github.com/micro/go-micro/v2/config/source/etcd"
)

type TestingStruct struct {
	Id int
	Name string
}

func main() {
	conf, err := microconfig.NewConfig()
	if err != nil {
		panic(err)
	}

	src := microetcd.NewSource(
		microetcd.WithAddress("127.0.0.1:2379"))

	conf.Load(src)

	confMap := conf.Map()
	fmt.Printf("conf map=:%v\n",confMap)

	fmt.Println("watch begin")
	defer fmt.Println("watch exit")
	for {

		watcher, err := conf.Watch("micro","config","bonjour")
		if err != nil {
			panic(err)
		}

		watcher.Next()

		err = conf.Load(src)
		if err != nil {
			panic(err)
		}

		val := conf.Get("micro","config","bonjour")
		fmt.Printf("the val is %s\n",val.Bytes())

		watcher.Stop()
	}
}