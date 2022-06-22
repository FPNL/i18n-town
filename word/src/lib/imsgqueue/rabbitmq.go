package imsgqueue

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func setup() (*amqp.Connection, error) {
	dns := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("IMSGQUEUE_USERNAME"),
		os.Getenv("IMSGQUEUE_PASSWORD"),
		os.Getenv("IMSGQUEUE_HOST"),
		os.Getenv("IMSGQUEUE_PORT"),
	)
	return amqp.Dial(dns)
}

func createCN(name string) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	chnq[name] = &q

	return ch, nil
}
