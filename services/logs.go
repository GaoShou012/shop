package main

import (
	"fmt"
	"github.com/app"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/v2"
	"time"
)

func main() {
	addr := fmt.Sprintf("%s:%d", app.Configs.LogRedis.Host, app.Configs.LogRedis.Port)
	db := app.Configs.LogRedis.DB
	pass := app.Configs.LogRedis.Password
	key := app.Configs.LogRedis.Key

	// 连接redis
	r := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           pass,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	handler := func(message string) {
		fmt.Println(message)
	}

	go func() {
		for {
			cmd := r.LPop(key)
			res, err := cmd.Result()
			if err != nil {
				if err.Error() == "redis: nil" {
					time.Sleep(time.Second * 1)
					continue
				} else {
					panic(err)
				}
			}
			handler(res)
		}
	}()

	service := micro.NewService(
		micro.Name("micro.service.logger"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
	)

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}
