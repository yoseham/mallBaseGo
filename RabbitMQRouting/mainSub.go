package main
import "../RabbitMQ"
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	rabbitmq.RecieveRouting()
}