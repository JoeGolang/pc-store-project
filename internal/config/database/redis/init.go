package redisConn

import "github.com/go-redis/redis/v8"

func InitRedisClient() *redis.Client {

	opts := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	}

	client := redis.NewClient(opts)

	return client
}
