package main
import "../RabbitMQ"
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("imoocSimple")
	rabbitmq.RecieveSub()
}
