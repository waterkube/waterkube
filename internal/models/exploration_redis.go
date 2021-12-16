package models

import "github.com/gomodule/redigo/redis"

// RedisExplorationRepository type.
type RedisExplorationRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}
