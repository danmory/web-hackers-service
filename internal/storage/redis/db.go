package redis

import (
	"os"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func GetDB() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	}
	return rdb
}
