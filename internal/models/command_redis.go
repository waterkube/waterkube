package models

import "github.com/gomodule/redigo/redis"

// RedisCommandRepository type.
type RedisCommandRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}
