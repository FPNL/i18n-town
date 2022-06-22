package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"math/rand"
)

type IPingService interface {
	Pong() (string, error)
}
type pingService struct {
	msgchn *amqp.Channel
	msgq   map[string]*amqp.Queue
}

var singlePing = pingService{}

func Ping(msgchn *amqp.Channel, msgq map[string]*amqp.Queue) IPingService {
	singlePing.msgchn = msgchn
	singlePing.msgq = msgq
	return &singlePing
}

func (service pingService) Pong() (string, error) {
	type Twilight struct {
		Hint   string
		Person string
	}

	var x Twilight
	if r := rand.Intn(100); r > 50 {
		x = Twilight{"Strix", "Anya"}
	} else {
		x = Twilight{"State Security Service", "Desmond"}
	}

	y, err := json.Marshal(x)
	if err != nil {
		return "", err
	}

	q, ok := service.msgq["hello"]
	if ok {
		err := service.msgchn.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/json",
				Body:        y,
			})
		if err != nil {
			return "", err
		}
	}

	return "pong", nil
}
