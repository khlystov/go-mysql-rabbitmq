package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionIndex  = "index"
)

// ClientConfig is the configuration for the client.
type ClientConfig struct {
	AmqpURL string
}

// NewClient creates the Cient with configuration.
func NewChannel(conf *ClientConfig) *amqp.Channel {
	conn, err := amqp.Dial(conf.AmqpURL)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	errorChannel := make(chan *amqp.Error)
	conn.NotifyClose(errorChannel)

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	go HowIsLife(errorChannel)

	return ch
}

func HowIsLife(errorChannel chan *amqp.Error) {
	for {
		err := <-errorChannel
		log.Fatalln(err)
	}
}

type BulkRequest struct {
	Action string
	Index  string
	Type   string
	ID     string
	Parent string

	Data map[string]interface{}
}
