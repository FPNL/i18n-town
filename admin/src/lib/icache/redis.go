package icache

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
)

func setupRedis() (*redis.Client, error) {
	dns := fmt.Sprintf("redis://%s:%s/%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_DEFAULT_DB"),
	)
	opt, err := redis.ParseURL(dns)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	err = client.Ping(context.Background()).Err()

	return client, err
}
