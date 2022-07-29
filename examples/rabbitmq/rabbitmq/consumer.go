package rabbitmq

// 消费者

import (
	"github.com/streadway/amqp"
	"strings"
)

type AMQPConsumer struct {
	Name        string
	Server      *AMQPServer
	StartButton chan bool
	CleanStart  bool
	Channel     *amqp.Channel
	Prefetch    int
	Handler     ConsumerHandler
	Session     AMQPSession
	Deliveries  <-chan amqp.Delivery
}

type ConsumerHandler interface {
	Init() error
	Process(delivery amqp.Delivery) error
}

func (c *AMQPConsumer) Connect() error {
	var err error
	c.Channel, err = c.Server.MessageBroker.NewChannel()
	if err != nil {
		return err
	}

	err = c.Channel.Qos(c.Prefetch, 0, false)
	if err != nil {
		return err
	}

	q := c.Session.QueueInfo
	e := c.Session.ExchangeInfo
	bo := c.Session.BindOptions
	co := c.Session.ConsumeOptions

	var queue amqp.Queue
	if queue, err = c.Channel.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Exclusive,
		q.NoWait,
		q.Args,
	); err != nil {
		return err
	}

	if c.CleanStart {
		if _, err = c.Channel.QueuePurge(q.Name, false); err != nil {
			return err
		}
	}

	if e.Name != "" && e.Type != "" {
		if err = c.Channel.ExchangeDeclare(
			e.Name,
			e.Type,
			e.Durable,
			e.AutoDelete,
			e.Internal,
			e.NoWait,
			e.Args,
		); err != nil {
			return err
		}

		for _, routingKey := range strings.Split(bo.RoutingKey, ",") {
			if err = c.Channel.QueueBind(
				queue.Name,
				routingKey,
				e.Name,
				bo.NoWait,
				bo.Args,
			); err != nil {
				return err
			}
		}
	}

	deliveries, err := c.Channel.Consume(
		queue.Name,
		co.Tag,
		co.AutoAck,
		co.Exclusive,
		co.NoLocal,
		co.NoWait,
		co.Args,
	)
	if err != nil {
		return err
	}

	c.Deliveries = deliveries
	return nil
}
