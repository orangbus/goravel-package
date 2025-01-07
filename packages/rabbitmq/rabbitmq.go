package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/facades"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Rabbitmq struct {
	queueName string
	conn      *amqp.Connection
	chanel    *amqp.Channel
}

func NewClient() (*Rabbitmq, error) {
	host := facades.Config().GetString("rabbitmq.host")
	port := facades.Config().GetInt("rabbitmq.port")
	name := facades.Config().GetString("rabbitmq.username")
	password := facades.Config().GetString("rabbitmq.password")
	sdn := fmt.Sprintf("amqp://%s:%s@%s:%d", name, password, host, port)

	queueName := facades.Config().GetString("rabbitmq.queue")
	var mq = &Rabbitmq{}
	var err error
	u := fmt.Sprintf("%s/%s", sdn, queueName)
	mq.conn, err = amqp.DialConfig(u, amqp.Config{Heartbeat: 10})
	if err != nil {
		return nil, err
	}
	mq.queueName = queueName
	mq.chanel, err = mq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return mq, nil
}

func (c *Rabbitmq) Close() {
	if err := c.chanel.Close(); err != nil {
		log.Printf("rabbitmq channel close error: %v", err)
	}
	if err := c.conn.Close(); err != nil {
		log.Printf("rabbitmq conn close error: %v", err)
	}
}

func (c *Rabbitmq) QueueDeclare(data interface{}) error {
	_, err := c.chanel.QueueDeclare(c.queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()
	body, _ := json.Marshal(data)
	err = c.chanel.PublishWithContext(ctx, "", c.queueName, false, false, amqp.Publishing{
		ContentType: "text/plan",
		Body:        body,
	})
	return err
}
