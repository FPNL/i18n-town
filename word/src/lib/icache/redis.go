package icache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"os"
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
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
}

//
//var ctx = context.Background()
//
//func ExampleClient() {
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
