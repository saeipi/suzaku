package rabbitmq

import (
	"errors"
)

var (
	ErrorAMQPServerNotReadyForConnection       = errors.New("amqp server not ready for connection")
	ErrorAMQPServerNotReadyForPublisherChannel = errors.New("amqp server not ready for publisher channel")
)
