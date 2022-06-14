package imsgqueue

import "github.com/streadway/amqp"

func setup() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@i18n_iqueue_1:5672/")
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
