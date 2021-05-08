package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("imoocSimple")
	for i :=0; i < 100; i++ {
		rabbitmq.PublishPub("Hello imooc" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("发送成功")
}
