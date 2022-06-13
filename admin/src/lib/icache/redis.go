package icache

import (
	"context"
	"github.com/go-redis/redis/v9"
)

const (
	E_REDIS_URI = "redis://localhost:6379/0"
)

func setupRedis() (*redis.Client, error) {
	opt, err := redis.ParseURL(E_REDIS_URI)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	err = client.Ping(context.Background()).Err()

	return client, err
}
