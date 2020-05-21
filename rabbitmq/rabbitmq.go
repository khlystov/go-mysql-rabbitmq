package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	ActionCreate = "create"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionIndex  = "index"
)

const (
	connPattern = "amqp://%s:%s@%s/"
)

// ClientConfig is the configuration for the client.
type ClientConfig struct {
	Host     string
	User     string
	Password string
}

// NewClient creates the Cient with configuration.
func NewChannel(conf *ClientConfig) *amqp.Channel {
	conn, err := amqp.Dial(
		fmt.Sprintf(connPattern, conf.User, conf.Password, conf.Host),
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	return ch
}

type BulkRequest struct {
	Action string
	Index  string
	Type   string
	ID     string
	Parent string

	Data map[string]interface{}
}
