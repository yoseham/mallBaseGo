package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式 amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@localhost:5672/imooc"
type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	Queuename string
	Exchange string
	Key string
	Mqurl string
}

func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{Queuename: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err.Error())
		panic(fmt.Sprintf("%s:%s", message, err.Error()))
	}
}
//1.创建Simple模式实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "" )
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取Channel失败")

	return rabbitmq
}
//2.简单模式生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列，不存在创建，否则跳过
	_, err := r.channel.QueueDeclare(
		r.Queuename,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否排他性
		false,
		//是否阻塞
		false,
		nil,
		)
	if err != nil {
		fmt.Println(err)
	}
	//发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.Queuename,
		//true:根据exchange类型和routkey规则，如果找不到合适队列会返回发送者
		false,
		//true:exchange发送到队列后发现队列没有绑定消费者，则会返回发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	//申请队列，不存在创建，否则跳过
	_, err := r.channel.QueueDeclare(
		r.Queuename,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否排他性
		false,
		//是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	mes, err := r.channel.Consume(
		r.Queuename,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//排他性
		false,
		//true:不能将同一个connection的消息传递给connection的消费者
		false,
		//阻塞
		false,
		nil)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程创建消息
	go func() {
		for d := range mes {
			//实现处理函数
			log.Printf("收到消息:%s", d.Body)
		}
	}()
	log.Println("等待接收消息")
	<-forever
}

func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error

	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取Channel失败")

	return rabbitmq
}

func (r *RabbitMQ) PublishPub(message string) {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型，订阅模式
		"fanout",
		true,
		false,
		//true:交换机不能推送消息，只能用于exchange之间的绑定
		false,
		false,
		nil)
	r.failOnErr(err, "无法创建交换机")

	//发送消息
	err = r.channel.Publish(
			r.Exchange,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:[]byte(message),
			})
	r.failOnErr(err, "无法发送成功")
}

func (r *RabbitMQ) RecieveSub() {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型，订阅模式
		"fanout",
		true,
		false,
		//true:交换机不能推送消息，只能用于exchange之间的绑定
		false,
		false,
		nil)
	r.failOnErr(err, "无法创建交换机")

	//创建队列，队列名称不能写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	r.failOnErr(err, "无法创建队列")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"",//必须为空
		r.Exchange,
		false,
		nil)

	//接收消息
	mes, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "接收失败")

	forever := make(chan bool)
	//启用协程创建消息
	go func() {
		for d := range mes {
			//实现处理函数
			log.Printf("收到消息:%s", d.Body)
		}
	}()
	log.Println("等待接收消息")
	<-forever
}

func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error

	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取Channel失败")

	return rabbitmq
}

func (r *RabbitMQ) PublishRouting(message string) {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//路由模式
		"direct",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "无法创建交换机")

	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,//
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		})
	r.failOnErr(err, "无法发送成功")
}

func (r *RabbitMQ) RecieveRouting() {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "无法创建交换机")

	//创建队列，队列名称不能写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	r.failOnErr(err, "无法创建队列")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,//
		r.Exchange,
		false,
		nil)

	//接收消息
	mes, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "接收失败")

	forever := make(chan bool)
	//启用协程创建消息
	go func() {
		for d := range mes {
			//实现处理函数
			log.Printf("收到消息:%s", d.Body)
		}
	}()
	log.Println("等待接收消息")
	<-forever
}
