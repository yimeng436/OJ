package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"questionsubmitsvr/config"
	"questionsubmitsvr/log"
	"sync"
)

// rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

var (
	mqonce   sync.Once
	mqclient *RabbitMQ
)

// 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	var MQURL = fmt.Sprintf("amqp://%s:%s@%s:%s", config.GetGlobalConfig().RabbitMQConfig.Username,
		config.GetGlobalConfig().RabbitMQConfig.Password,
		config.GetGlobalConfig().RabbitMQConfig.Host,
		config.GetGlobalConfig().RabbitMQConfig.Port)
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	mqclient = rabbitmq
	fmt.Println("队列创建成功")
	return rabbitmq
}

// 路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//要改成direct
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an excha"+
		"nge")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//要设置
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 路由模式接受消息
func (r *RabbitMQ) RecieveRouting() {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+
		"ange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//需要绑定key
		r.Key,
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messges {
			log.Infof("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("退出请按 CTRL+C")
	<-forever
}

func GetMQ(exchangeName string, routingKey string) *RabbitMQ {
	mqonce.Do(func() {
		NewRabbitMQRouting(exchangeName, routingKey)
	})
	return mqclient
}
