package queue

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/wanglilind/qqq/pkg/config"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(config *config.Config) (*RabbitMQ, error) {
	// 建立连接
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
	))
	if err != nil {
		return nil, err
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

func (r *RabbitMQ) PublishMessage(exchange, routingKey string, message []byte) error {
	return r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:       message,
		},
	)
} 
