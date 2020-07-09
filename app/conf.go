package app

import (
	"fmt"
)

const (
	confFilePath = "E:\\work\\golang\\shopServer\\app\\conf.yml"
)

var Configs struct {
	Etcd struct {
		Host string
		Port string
	}
	DB struct {
		User            string
		Password        string
		Host            string
		Port            string
		DefaultDatabase string
		Dns             string
	}
	RabbitMq struct {
		User     string
		Password string
		Host     string
		Port     string
	}
	LogRedis struct {
		Host     string
		Port     int
		Password string
		DB       int
		Key      string
	}
}

func RabbitMqUrl() string {
	return fmt.Sprintf("amqp://admin:admin@%s:%s/", Configs.RabbitMq.Host, Configs.RabbitMq.Port)
}
