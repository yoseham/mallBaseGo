package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)
func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	for i :=0; i < 100; i++ {
		imoocOne.PublishRouting("Hello imooc_one" + strconv.Itoa(i))
		imoocTwo.PublishRouting("Hello imooc_two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("发送成功")
}
