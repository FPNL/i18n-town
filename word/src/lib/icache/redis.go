package icache

import (
	"context"
	"github.com/go-redis/redis/v9"
)

const (
	E_REDIS_URI = "redis://i18n_icache_1:6379/0"
)

func setupRedis() (*redis.Client, error) {
	opt, err := redis.ParseURL(E_REDIS_URI)
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
