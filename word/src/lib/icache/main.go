package icache

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var rdb *redis.Client

func Go() (err error) {
	rdb, err = setupRedis()
	if err != nil {
		fmt.Println("cache done")
	}

	return err
}

func Connect() *redis.Client {
	if rdb == nil {
		log.Fatalln("專案架構級別錯誤")
	}
	return rdb
}

func Close() {
	if rdb == nil {
		log.Fatalln("專案架構級別錯誤")
	}
	err := rdb.Close()
	if err != nil {
		log.Fatalln("專案架構級別錯誤")
	}
}
