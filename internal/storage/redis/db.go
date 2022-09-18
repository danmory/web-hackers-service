package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func GetDB() *redis.Client {
	if rdb == nil {
		address := fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		rdb = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	}
	return rdb
}

func CloseDB() error {
	return rdb.Close()
}