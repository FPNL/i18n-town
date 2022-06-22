package imsgqueue

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

var conn *amqp.Connection
var chn *amqp.Channel
var chnq = make(map[string]*amqp.Queue)

func Go() (err error) {
	for retry := 1; retry <= 3; retry++ {
		time.Sleep(5 * time.Second)
		conn, err = setupConn()
		if err == nil {
			break
		} else {
			log.Printf("錯誤 %d 次\n", retry)
		}
	}

	if err != nil {
		return err
	}

	fmt.Println("ianime 開始工作",
		os.Getenv("IMSGQUEUE_HOST"),
		os.Getenv("IMSGQUEUE_PORT"),
	)
	return setupWorker()
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
