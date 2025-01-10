package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/facades"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Rabbitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	ctx       context.Context
	queueName string // 队列的名称
	key       string // 路由key
	exchange  string // 交换机
}

func NewRabbitmq() (*Rabbitmq, error) {
	host := facades.Config().GetString("rabbitmq.host")
	port := facades.Config().GetInt("rabbitmq.port")
	name := facades.Config().GetString("rabbitmq.username")
	password := facades.Config().GetString("rabbitmq.password")
	vhost := facades.Config().GetString("rabbitmq.vhost")
	sdn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", name, password, host, port, vhost)
	queueName := facades.Config().GetString("rabbitmq.queue")

	var mq = &Rabbitmq{}
	var err error
	mq.conn, err = amqp.DialConfig(sdn, amqp.Config{Heartbeat: 10})
	if err != nil {
		return nil, err
	}
	mq.queueName = queueName
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return mq, nil
}

func (c *Rabbitmq) Close() {
	if err := c.channel.Close(); err != nil {
		log.Printf("rabbitmq channel close error: %v", err)
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("rabbitmq conn close error: %v", err)
	}
}

// 统一发送消息
func (r *Rabbitmq) seed(data any) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.channel.PublishWithContext(r.ctx, "", r.queueName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         marshal,
	})
}

func (r *Rabbitmq) Msg(data any) error {
	_, err := r.channel.QueueDeclare(r.queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := r.channel.Qos(1, 0, false); err != nil {
		return err
	}
	return r.seed(data)
}
func (r *Rabbitmq) Publish(exchangeName, exchangeType string, data interface{}) error {
	// 定义交换机
	err := r.channel.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.channel.PublishWithContext(r.ctx, exchangeName, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         marshal,
	})
	return nil
}
func (r *Rabbitmq) Routing(exchangeName, key string, data interface{}) error {
	if err := r.channel.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil); err != nil {
		return err
	}
	marshal, err2 := json.Marshal(data)
	if err2 != nil {
		return err2
	}
	return r.channel.PublishWithContext(r.ctx, exchangeName, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        marshal,
	})
}
func (r *Rabbitmq) Topic(exchangeName, key string, data interface{}) error {
	// 1、定义交换机
	if err := r.channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil); err != nil {
		return err
	}

	marshal, err2 := json.Marshal(data)
	if err2 != nil {
		return err2
	}
	return r.channel.PublishWithContext(r.ctx, exchangeName, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        marshal,
	})
}
func (r *Rabbitmq) ReceiverTopic(exchangeName, key string) (<-chan amqp.Delivery, error) {
	// 1、定义交换机
	err := r.channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	// 2、定义消息队列
	q, err := r.channel.QueueDeclare(r.queueName, true, false, true, false, nil)
	if err != nil {
		return nil, err
	}

	// 3、绑定:队列名称 key 交换机
	if err := r.channel.QueueBind(r.queueName, key, exchangeName, false, nil); err != nil {
		return nil, err
	}

	return r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
}

func (r *Rabbitmq) Copsume() {
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("consume error: %s", err.Error())
		return
	}
	var forever chan struct{}
	go func() {
		for data := range msgs {
			log.Printf("接收到的消息是：%s", string(data.Body))
			err := data.Ack(false)
			if err != nil {
				log.Printf("ack error: %s", err.Error())
			}
		}
	}()
	<-forever
}

// 消费消息
func (r *Rabbitmq) copsume() (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		r.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}
