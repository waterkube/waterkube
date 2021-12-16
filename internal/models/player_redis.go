package models

import "github.com/gomodule/redigo/redis"

// RedisPlayerRepository type.
type RedisPlayerRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}
