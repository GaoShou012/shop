package main

import (
	"fmt"
	"github.com/app"
	_ "github.com/app"
	"github.com/streadway/amqp"
	"time"
)

func pub() {
	// 连接
	connection, err := amqp.Dial(app.RabbitMqUrl())
	if err != nil {
		panic(fmt.Sprintf("MQ连接失败:%v\n", err))
	}

	// 打开通道
	channel, err := connection.Channel()
	if err != nil {
		panic(fmt.Sprintf("MQ打开Channel失败:%v\n", err))
	}

	count := 0
	for {
		count++
		content := fmt.Sprintf("%d",count)
		message := amqp.Publishing{
			Headers:         nil,
			ContentType:     "",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(content),
		}
		err = channel.Publish("orders", "", false, false, message)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	pub()
}
