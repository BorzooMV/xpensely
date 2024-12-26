package services

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
	})
}
