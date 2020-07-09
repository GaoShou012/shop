package app

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/micro/go-micro/v2"
	microconfig "github.com/micro/go-micro/v2/config"
	microetcd "github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

)

func init() {
	// 加载配置文件
	err := configor.Load(&Configs,confFilePath)
	if err != nil {
		panic(err)
	}

	// 日志配置
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat:"2006-01-02 15:04:05"})
}


var etcdRegistry registry.Registry
var etcdRegistryInit sync.Once
func EtcdRegistry() registry.Registry {
	etcdRegistryInit.Do(func() {
		r := etcd.NewRegistry(
			registry.Addrs(fmt.Sprintf("%s:%s",Configs.Etcd.Host,Configs.Etcd.Port)),
			)
		etcdRegistry = r
	})

	return etcdRegistry
}

func EtcdAddress() string {
	return "127.0.0.1:2379"
}

func NewService(name string) micro.Service {
	// etcd registry
	r := EtcdRegistry()

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(name),
		micro.Registry(r),
		micro.RegisterTTL(time.Second*10),
		)

	// Init will parse the command line flags.
	service.Init()

	return service
}

var serviceInit sync.Once
var service micro.Service
func Service() micro.Service {
	serviceInit.Do(func() {
		service = micro.NewService(micro.Registry(EtcdRegistry()))
		service.Init()
	})
	return service
}

func EtcdSource() (conf microconfig.Config,err error) {
	conf,err = microconfig.NewConfig()
	if err != nil {
		return
	}
	src := microetcd.NewSource(microetcd.WithAddress(EtcdAddress()))
	err = conf.Load(src)
	if err != nil {
		return
	}
	return
}