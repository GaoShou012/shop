package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	url := fmt.Sprintf("amqp://admin:admin@192.168.0.200:5672/")
	// 连接
	connection,err := amqp.Dial(url)
	if err != nil {
		fmt.Printf("MQ连接失败! %s\n",err)
		return
	}
	fmt.Println("MQ连接成功")

	// 打开通道
	channel,err := connection.Channel()
	if err != nil {
		fmt.Printf("MQ打开Channel失败! %s\n",err)
		return
	}
	fmt.Println("MQ打开Channel成功!")

	exchangeName := "test"
	queueName := "testQueue1"

	// 创建交换机
	err = channel.ExchangeDeclare(exchangeName,amqp.ExchangeFanout,true,false,false,false,nil)
	if err != nil {
		fmt.Printf("创建交换机失败! %s\n",exchangeName)
		return
	}else{
		fmt.Printf("创建交换机成功! %s\n",exchangeName)
	}

	// 创建消息队列
	_,err = channel.QueueDeclare(queueName,true,false,false,false,nil)
	if err != nil {
		fmt.Println("创建消息队列失败!")
		return
	}

	err = channel.QueueBind(queueName,queueName,exchangeName,false,nil)
	if err != nil {
		fmt.Println("消息队列绑定失败!")
		return
	}

	//message := amqp.Publishing{
	//	Headers:         nil,
	//	ContentType:     "",
	//	ContentEncoding: "",
	//	DeliveryMode:    0,
	//	Priority:        0,
	//	CorrelationId:   "",
	//	ReplyTo:         "",
	//	Expiration:      "",
	//	MessageId:       "",
	//	Timestamp:       time.Time{},
	//	Type:            "",
	//	UserId:          "",
	//	AppId:           "",
	//	Body:            []byte("1111"),
	//}
	//err = channel.Publish(exchangeName,queueName,false,false,message)
	//if err != nil {
	//	fmt.Println("发送消息失败!")
	//	fmt.Println(err)
	//	return
	//}

	// 拉消息
	msg,ok,err := channel.Get(queueName,false)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !ok {
		fmt.Println("n ok")
		return
	}
	fmt.Println(msg.MessageId)

	fmt.Println(string(msg.Body))
	//msg.Nack(true,true)


	// 拉消息
	msg,ok,err = channel.Get(queueName,false)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !ok {
		fmt.Println("n ok")
		return
	}

	fmt.Println(string(msg.Body))
	fmt.Println(msg.MessageId)

	fmt.Println("OK")
}
