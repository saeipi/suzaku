package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
	"reflect"
)

type AMQPServer struct {
	MessageBroker AMQPMessageBroker
	InfoConfig    *AMQPInfoConfig
	Consumers     []*AMQPConsumer
	Publishers    []*AMQPPublisher
	Group         *errgroup.Group
	Initialized   bool
	InitSignal    chan bool
}

type AMQPSession struct {
	ExchangeInfo   AMQPExchangeInfo
	QueueInfo      AMQPQueueInfo
	BindOptions    AMQPBindOptions
	ConsumeOptions AMQPConsumeOptions
	PublishOptions AMQPPublishOptions
}

var server *AMQPServer

func NewAMQPServer(messageBroker AMQPMessageBroker, infoConfig *AMQPInfoConfig) *AMQPServer {
	group, _ := errgroup.WithContext(context.Background())

	server = &AMQPServer{
		MessageBroker: messageBroker,
		InfoConfig:    infoConfig,
		Group:         group,
		InitSignal:    make(chan bool),
	}
	return server
}

func IsValidExchangeType(exchangeType string) bool {
	var (
		ExchangeTypes = map[string]bool{
			amqp.ExchangeFanout:  true,
			amqp.ExchangeDirect:  true,
			amqp.ExchangeTopic:   true,
			amqp.ExchangeHeaders: true,
		}
	)

	if _, ok := ExchangeTypes[exchangeType]; ok {
		return true
	}

	return false
}

func (s *AMQPServer) NewStandardPublisher(exchangeKey, queueKey string, po AMQPPublishOptions) *AMQPPublisher {
	config := s.InfoConfig
	exchangeInfoConfig := config.ExchangeInfoConfig
	queueInfoConfig := config.QueueInfoConfig

	var exchangeExist, queueExist bool
	var session AMQPSession
	if exchangeKey != "" {
		if e, ok := exchangeInfoConfig[exchangeKey]; ok {
			if !IsValidExchangeType(e.Type) {
				return nil
			}
			session.ExchangeInfo = e
			exchangeExist = true
		} else {
			return nil
		}
	}

	if queueKey != "" {
		if q, ok := queueInfoConfig[queueKey]; ok {
			session.QueueInfo = q
			queueExist = true
		} else {
			return nil
		}
	}

	if !exchangeExist && !queueExist {
		return nil
	} else if !exchangeExist {
		po.RoutingKey = session.QueueInfo.Name
	}
	session.PublishOptions = po

	publisher := &AMQPPublisher{
		Server:  s,
		Session: session,
	}

	s.Publishers = append(s.Publishers, publisher)
	return publisher
}

func (s *AMQPServer) NewAMQPPublisher(w string, po AMQPPublishOptions) *AMQPPublisher {
	config := s.InfoConfig

	exchangeInfoConfig := config.ExchangeInfoConfig
	queueInfoConfig := config.QueueInfoConfig
	bindingInfoConfig := config.BindingInfoConfig

	var session AMQPSession
	if bi, ok := bindingInfoConfig[w]; ok {
		session.BindOptions = bi.BindOptions

		if q, ok := queueInfoConfig[bi.QueueKey]; ok {
			session.QueueInfo = q
		} else {
			return nil
		}

		if e, ok := exchangeInfoConfig[bi.ExchangeKey]; ok {
			if !IsValidExchangeType(e.Type) {
				return nil
			}
			session.ExchangeInfo = e
		} else {
			po.RoutingKey = session.QueueInfo.Name
		}

		if session.BindOptions.RoutingKey == "" {
			if session.ExchangeInfo.Type == amqp.ExchangeDirect {
				session.BindOptions.RoutingKey = w
				po.RoutingKey = w
			}
		}

		session.PublishOptions = po
	} else {
		return nil
	}

	publisher := &AMQPPublisher{
		Server:  s,
		Session: session,
	}

	s.Publishers = append(s.Publishers, publisher)
	return publisher
}

func (s *AMQPServer) registerAllConsumers(consumerHandlers map[string]ConsumerHandler) {
	config := s.InfoConfig

	exchangeInfoConfig := config.ExchangeInfoConfig
	queueInfoConfig := config.QueueInfoConfig
	bindingInfoConfig := config.BindingInfoConfig
	channelInfoConfig := config.ChannelInfoConfig

	for w, bi := range bindingInfoConfig {
		if _, ok := consumerHandlers[w]; !ok {
			continue
		}

		var session AMQPSession
		session.BindOptions = bi.BindOptions

		if e, ok := exchangeInfoConfig[bi.ExchangeKey]; ok {
			if !IsValidExchangeType(e.Type) {
				continue
			}
			session.ExchangeInfo = e
		}

		if q, ok := queueInfoConfig[bi.QueueKey]; ok {
			session.QueueInfo = q
		} else {
			continue
		}

		if session.BindOptions.RoutingKey == "" {
			if session.ExchangeInfo.Type == amqp.ExchangeDirect {
				session.BindOptions.RoutingKey = session.QueueInfo.Name
			} else if session.ExchangeInfo.Type == amqp.ExchangeTopic {
				continue
			}
		}

		var prefetch int
		if c, ok := channelInfoConfig[w]; ok {
			prefetch = c.Prefetch
		}

		if prefetch <= 0 {
			prefetch = 1
		}

		s.Consumers = append(s.Consumers, &AMQPConsumer{
			Name:        w,
			Server:      s,
			StartButton: make(chan bool),
			CleanStart:  bi.CleanStart,
			Prefetch:    prefetch,
			Handler:     consumerHandlers[w],
			Session:     session,
		})
	}
}

type PublisherRegister interface {
	Register()
}

func (s *AMQPServer) RegisterAndRun(consumerHandlers map[string]ConsumerHandler, publisherRegisters ...PublisherRegister) {
	for _, r := range publisherRegisters {
		if r != nil && !reflect.ValueOf(r).IsNil() {
			r.Register()
		}
	}
	s.registerAllConsumers(consumerHandlers)

	go func() {
		s.Run()
	}()

	s.WaitInitialized()
}

func (s *AMQPServer) Run() error {
	go func() {
		for {
			select {
			case connected := <-s.MessageBroker.ConnectionStatusChannel():
				if !connected {
					if err := s.MessageBroker.NewConnection(); err == nil {
						for _, consumer := range s.Consumers {
							consumer.StartButton <- true
						}

						for _, publisher := range s.Publishers {
							publisher.Connect()
						}

						go func() {
							s.InitSignal <- true
						}()
					}
				}
			}
		}
	}()

	for _, consumer := range s.Consumers {
		consumer := consumer

		s.Group.Go(func() error {
			for {
				select {
				case permitted := <-consumer.StartButton:
					if permitted {
						if err := consumer.Connect(); err != nil {
							panic(err)
						}

						if err := consumer.Handler.Init(); err != nil {
							panic(err)
						}

						for i := 0; i < consumer.Prefetch; i++ {
							go func() {
								for delivery := range consumer.Deliveries {
									consumer.Handler.Process(delivery)
									delivery.Ack(false)
								}
							}()
						}
					}
				}
			}

			return nil
		})
	}

	return s.Group.Wait()
}

func (s *AMQPServer) WaitInitialized() {
	if s.Initialized {
		return
	}
	s.Initialized = <-server.InitSignal
}
