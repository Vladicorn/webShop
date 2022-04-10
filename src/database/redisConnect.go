package database

import (
	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "192.168.0.57:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
