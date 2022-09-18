package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func RetrieveHackers() ([]redis.Z, error) {
	return GetDB().ZRangeWithScores(context.Background(), "hackers", 0, -1).Result()
}
