package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

func SetupRabbitMq() {
	var err error
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s", "admin", "123456", "127.0.0.1:5672")
	connection, err = amqp.Dial(RabbitUrl)
	if err != nil {
		fmt.Printf("MQ打开链接失败:%s \n", err)
		return
	}
	channel, err = connection.Channel()
	if err != nil {
		fmt.Printf("MQ打开管道失败:%s \n", err)
		return
	}
	fmt.Println("rabbitMq连接成功")
}

// 关闭RabbitMQ连接
func Close() {
	// 先关闭管道,再关闭链接
	err := channel.Close()
	if err != nil {
		fmt.Printf("MQ管道关闭失败:%s \n", err)
	}
	err = connection.Close()
	if err != nil {
		fmt.Printf("MQ链接关闭失败:%s \n", err)
	}
}

func Publish() {

}

func Consumer() {

}
