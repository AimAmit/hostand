package cache

import (
	"github.com/go-redis/redis"
)

var (
	Cache redis.Client
)

func GetRedisConn() {
	Cache = *redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
