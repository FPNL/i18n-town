package imsgqueue

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func setupConn() (*amqp.Connection, error) {
	dns := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("IMSGQUEUE_USERNAME"),
		os.Getenv("IMSGQUEUE_PASSWORD"),
		os.Getenv("IMSGQUEUE_HOST"),
		os.Getenv("IMSGQUEUE_PORT"),
	)
	return amqp.Dial(dns)
}

func createCN(name string) (<-chan amqp.Delivery, error) {
	var err error
	chn, err = conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := chn.QueueDeclare(
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

	//serv.msgChn.Consume(
	msgs, err := chn.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
