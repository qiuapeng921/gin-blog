package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

var (
	connection *amqp.Connection
	//channel    *amqp.Channel
)

func SetupRabbitMq() {
	var err error
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s", "admin", "123456", "127.0.0.1:5672")
	// 创建连接
	connection, err = amqp.Dial(RabbitUrl)
	if err != nil {
		panic("MQ打开链接失败")
	}
	//// 创建信道
	//channel, err = connection.Channel()
	//if err != nil {
	//	panic("MQ打开管道失败")
	//}
	fmt.Println("rabbitMq连接成功")
}

// 关闭RabbitMQ连接
func Close() {
	// 先关闭管道,再关闭链接
	//err := channel.Close()
	//if err != nil {
	//	fmt.Println("MQ管道关闭失败")
	//}
	err := connection.Close()
	if err != nil {
		fmt.Println("MQ链接关闭失败")
	}
}

// 生产队列消息
func Publish(exchangeName, queueName, data string) error {
	// 创建交换机
	channel, err := connection.Channel()
	if err != nil {
		return err
	}
	// exchangeType = "direct", "Exchange type - direct|fanout|topic|x-custom")
	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, true, nil)
	if err != nil {
		fmt.Printf("MQ注册交换机失败:%s \n", err)
		return err
	}

	// 创建队列
	var queue amqp.Queue
	queue, err = channel.QueueDeclare(queueName, true, false, false, true, nil)
	if err != nil {
		return err
	}

	// 队列绑定
	err = channel.QueueBind(queueName, queue.Name, exchangeName, true, nil)
	if err != nil {
		fmt.Printf("MQ绑定队列失败:%s \n", err)
		return err
	}

	// 发送任务消息
	err = channel.Publish(exchangeName, queue.Name, false, false, amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Body:            []byte(data),
		DeliveryMode:    amqp.Transient,
		Priority:        0,
	})
	if err != nil {
		fmt.Printf("MQ任务发送失败:%s \n", err)
		return err
	}
	return nil
}

//消费队列消息
func Consumer(exchangeName, queueName string) error {
	return nil
}
