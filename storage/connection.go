package storage

import (
	"github.com/redis/go-redis/v9"
)

func ConnectRedisDB() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	return client
}
