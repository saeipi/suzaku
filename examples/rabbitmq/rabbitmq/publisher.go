package rabbitmq

import (
	"github.com/streadway/amqp"
)

type AMQPPublisher struct {
	Server  *AMQPServer
	Channel *amqp.Channel
	Session AMQPSession
}

func (p *AMQPPublisher) Connect() error {
	var err error
	p.Channel, err = p.Server.MessageBroker.NewChannel()
	if err != nil {
		return err
	}

	q := p.Session.QueueInfo
	e := p.Session.ExchangeInfo

	if q.Name != "" {
		if _, err = p.Channel.QueueDeclare(
			q.Name,
			q.Durable,
			q.AutoDelete,
			q.Exclusive,
			q.NoWait,
			q.Args,
		); err != nil {
			return err
		}
	}

	if e.Name != "" && e.Type != "" {
		if err = p.Channel.ExchangeDeclare(
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
	}

	return nil
}

func (p *AMQPPublisher) Publish(body []byte) error {
	var err error

	cp := AMQPPublisher{
		Server:  p.Server,
		Channel: p.Channel,
		Session: p.Session,
	}

	e := cp.Session.ExchangeInfo
	po := cp.Session.PublishOptions
	po.Publishing.Body = body

	if cp.Channel == nil {
		// unexpected to be here
		return ErrorAMQPServerNotReadyForPublisherChannel
		//if err = cp.Connect(); err != nil {
		//    return err
		//}
	}

	err = cp.Channel.Publish(
		e.Name,
		po.RoutingKey,
		po.Mandatory,
		po.Immediate,
		po.Publishing,
	)

	return err
}
