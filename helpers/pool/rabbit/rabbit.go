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
	var err error
	// 先关闭管道,再关闭链接
	err = channel.Close()
	if err != nil {
		fmt.Println("MQ管道关闭失败")
	}
	err = connection.Close()
	if err != nil {
		fmt.Println("MQ链接关闭失败")
	}
}

// 生产队列消息
func Publish(exchangeName, queueName, data string) error {
	// 创建交换机
	var err error
	channel, err = connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	// exchangeType = "direct", "Exchange type - direct|fanout|topic|x-custom")
	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("MQ注册交换机失败:%s \n", err)
		return err
	}

	// 创建队列
	var queue amqp.Queue
	queue, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
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
	// 创建交换机
	var err error
	channel, err = connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	// exchangeType = "direct", "Exchange type - direct|fanout|topic|x-custom")
	err = channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("MQ注册交换机失败:%s \n", err)
		return err
	}

	// 创建队列
	var queue amqp.Queue
	queue, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// 队列绑定
	err = channel.QueueBind(queueName, queue.Name, exchangeName, false, nil)
	if err != nil {
		fmt.Printf("MQ绑定队列失败:%s \n", err)
		return err
	}

	// 获取消费通道 确保rabbitmq会一个一个发消息
	channel.Qos(1, 0, true)
	var message <-chan amqp.Delivery
	message, err = channel.Consume(
		queueName, // queue
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error())
	}

	// 使用callback消费数据
	for msg := range message {
		// 当接收者消息处理失败的时候，
		// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
		// 通过重试可以成功的操作，那么这个时候是需要重试的
		// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
		//for !receiver.OnReceive(msg.Body) {
		//	log.Warnf("receiver 数据处理失败，将要重试")
		//	time.Sleep(1 * time.Second)
		//}
		fmt.Println(string(msg.Body))
		// 确认收到本条消息, multiple必须为false
		msg.Ack(false)
		break
	}

	return nil
}
