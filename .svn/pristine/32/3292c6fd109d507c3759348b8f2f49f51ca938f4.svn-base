package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

/*
	etcd client v3
	set
	get
	watch
*/


func main () {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          time.Second*5,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("connect to etcd success")

	// 延时关闭
	defer cli.Close()

	// put
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	_,err = cli.Put(ctx,"micro/bonjour","test1")
	cancel()
	if err != nil {
		log.Printf("put to etcd failed,err:%v\n",err)
		return
	}

	// get
	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	res,err := cli.Get(ctx,"micro/bonjour")
	cancel()
	if err != nil {
		log.Printf("get from etcd failed,err:%v\n",err)
		return
	}
	for _,el := range res.Kvs {
		fmt.Printf("%s:%s\n",el.Key,el.Value)
	}

	timeoutWatch(cli)
	loopWatch(cli)
}

func loopWatch(cli *clientv3.Client) {
	fmt.Println("loopWatch Entry")
	defer fmt.Println("loopWatch Exit")

	watchChan := cli.Watch(context.Background(),"testabc")
	for watchResponse := range watchChan {
		for _,ev := range watchResponse.Events {
			fmt.Printf("Type: %s key:%s value:%s\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
		}
	}
}

func timeoutWatch(cli *clientv3.Client) {
	fmt.Println("timeoutWatch Entry")
	defer fmt.Println("timeoutWatch Exit")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	watchChan := cli.Watch(ctx,"testabc")
	for watchResponse := range watchChan {
		for _,ev := range watchResponse.Events {
			fmt.Printf("Type: %s key:%s value:%s\n",ev.Type,ev.Kv.Key,ev.Kv.Value)
		}
	}
	cancel()
}
