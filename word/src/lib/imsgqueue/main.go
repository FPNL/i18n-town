package imsgqueue

import (
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection
var chn *amqp.Channel
var chnq = make(map[string]*amqp.Queue)

func Go() (err error) {
	conn, err = setup()
	if err != nil {
		return err
	}
	chn, err = createCN("hello")
	if err != nil {
		return err
	}

	return nil
}

func ConnectChn() *amqp.Channel {
	if chn == nil {
		log.Fatalln("架構錯誤")
	}

	return chn
}

func Close() {
	if conn == nil {
		log.Fatalln("架構錯誤")
	}

	chn.Close()
	conn.Close()
}

func GetQueue() map[string]*amqp.Queue {
	return chnq
}
