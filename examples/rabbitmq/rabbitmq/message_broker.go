package rabbitmq

import (
	"github.com/streadway/amqp"
	"time"
)

type AMQPMessageBroker interface {
	NewConnection() error
	NewChannel() (*amqp.Channel, error)
	SignalConnectionStatus(status bool)
	ConnectionStatusChannel() chan bool
}

type MessageBroker struct {
	ConnectionConfig     *AMQPConnectionConfig
	Connection           *amqp.Connection
	ConnectionStatusChan chan bool
}

func NewAMQPMessageBroker(connectionConfig *AMQPConnectionConfig) AMQPMessageBroker {
	messageBroker := new(MessageBroker)

	messageBroker.ConnectionConfig = connectionConfig
	messageBroker.ConnectionStatusChan = make(chan bool)
	messageBroker.SignalConnectionStatus(false)

	return messageBroker
}

func (b *MessageBroker) NewConnection() error {
	var err error

	for {
		if err = b.Dial(); err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		go func() {
			err := <-b.Connection.NotifyClose(make(chan *amqp.Error))
			if err != nil {
				b.SignalConnectionStatus(false)
			}
		}()

		b.SignalConnectionStatus(true)
		break
	}

	return err
}

func (b *MessageBroker) Dial() (err error) {
	conf := amqp.URI{
		Scheme:   "amqp",
		Host:     b.ConnectionConfig.Host,
		Port:     b.ConnectionConfig.Port,
		Username: b.ConnectionConfig.User,
		Password: b.ConnectionConfig.Password,
	}.String()

	b.Connection, err = amqp.Dial(conf)
	return
}

func (b *MessageBroker) NewChannel() (*amqp.Channel, error) {
	if b.Connection == nil {
		// unexpected to be here
		return nil, ErrorAMQPServerNotReadyForConnection
		//if err := b.Dial(); err != nil {
		//    return nil, err
		//}
	}
	return b.Connection.Channel()
}

func (b *MessageBroker) SignalConnectionStatus(status bool) {
	go func() {
		b.ConnectionStatusChan <- status
	}()
}

func (b *MessageBroker) ConnectionStatusChannel() chan bool {
	return b.ConnectionStatusChan
}
