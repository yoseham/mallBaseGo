package main
import "../RabbitMQ"
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	rabbitmq.RecieveRouting()
}
