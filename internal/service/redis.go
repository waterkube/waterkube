package service

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/waterkube/waterkube/internal/config"
)

// RedisPool function.
func RedisPool(appConfig *config.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(appConfig.RedisURL)
		},
	}
}
