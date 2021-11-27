package cmd

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/petaki/support-go/cli"
)

func createRedisFlags(command *cli.Command) (*string, *string) {
	redisURL := command.FlagSet().String("redis-url", os.Getenv("REDIS_URL"), "Redis URL")
	redisKeyPrefix := command.FlagSet().String("redis-key-prefix", os.Getenv("REDIS_KEY_PREFIX"), "Redis Key Prefix")

	return redisURL, redisKeyPrefix
}

func newRedisPool(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
	}
}
