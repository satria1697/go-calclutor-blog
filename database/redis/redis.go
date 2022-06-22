package redis

import (
	"dumpro/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis(config utils.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: "",
		DB:       0,
	})

	return rdb
}
