package rabbitmq

import (
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, err
	}
	return conn, nil
}