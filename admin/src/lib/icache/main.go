package icache

import "github.com/go-redis/redis/v9"

var cache *redis.Client

func Go() (err error) {
	cache, err = setupRedis()
	return
}

func Connect() *redis.Client {
	if cache == nil {
		panic("專案架構層級錯誤")
	}
	return cache
}

func Close() {
	if cache == nil {
		panic("專案架構層級錯誤")
	}

	err := cache.Close()
	if err != nil {
		panic(err)
	}
}
