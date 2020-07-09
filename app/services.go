package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	microconfig "github.com/micro/go-micro/v2/config"
	microetcd "github.com/micro/go-micro/v2/config/source/etcd"

	"github.com/micro/go-micro/v2/client"
)

type RedisService struct {
	Host string
	User string
	Pass string
}

type Mysql struct {
	User string
	Password string
	Host string
	DefaultDatabase string
}



func ServiceClient() client.Client {
	service := micro.NewService(
		micro.Registry(EtcdRegistry()),
		)
	service.Init()
	return service.Client()
}

func RabbitMqFromEtcd() {
	// 连接配置中心
	conf,err := microconfig.NewConfig()
	if err != nil {
		panic(err)
	}

	etcdsrc := microetcd.NewSource(microetcd.WithAddress(EtcdAddress()))
	err = conf.Load(etcdsrc)
	if err != nil {
		return
	}
}

func (m *Mysql) GormFromEtcd() (db *gorm.DB,err error) {
	// 连接配置中心
	conf,err := microconfig.NewConfig()
	if err != nil {
		return
	}
	etcdsrc := microetcd.NewSource(microetcd.WithAddress(EtcdAddress()))
	err = conf.Load(etcdsrc)
	if err != nil {
		return
	}

	// 读取配置
	val := conf.Get("micro","config","mysql")
	err = json.Unmarshal(val.Bytes(),m)
	if err != nil {
		return
	}

	// 打印配置
	fmt.Printf("数据库配置%v\n",m)

	// 连接数据库
	db,err = gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		m.User,
		m.Password,
		m.Host,
		m.DefaultDatabase,
	))
	if err != nil {
		return
	}

	return
}

