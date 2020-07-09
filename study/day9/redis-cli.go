package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main() {
	r := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "192.168.0.200:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		DB:                 0,
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
	sub := r.Subscribe("test")
	for {
		message,err := sub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(message.Payload)
	}
}
