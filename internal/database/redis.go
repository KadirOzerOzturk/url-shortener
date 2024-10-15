package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func GetRedisClient() *redis.Client {
	return redisClient
}
