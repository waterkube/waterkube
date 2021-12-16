package models

import "github.com/gomodule/redigo/redis"

// RedisGridRepository type.
type RedisGridRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}
