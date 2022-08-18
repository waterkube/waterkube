package models

import (
	"github.com/gomodule/redigo/redis"
)

const (
	playerKeyPrefix = "player"
)

// RedisPlayerRepository type.
type RedisPlayerRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}

// CreateOrUpdate function.
func (rpr *RedisPlayerRepository) CreateOrUpdate(player *Player) error {
	conn := rpr.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do(
		"HSET", redis.Args{}.Add(rpr.RedisKeyPrefix+playerKeyPrefix).AddFlat(player)...,
	)
	if err != nil {
		return err
	}

	return nil
}

// Find function.
func (rpr *RedisPlayerRepository) Find() (*Player, error) {
	conn := rpr.RedisPool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", rpr.RedisKeyPrefix+playerKeyPrefix))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, ErrNoRecord
	}

	var player Player

	err = redis.ScanStruct(values, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

// Delete function.
func (rpr *RedisPlayerRepository) Delete() error {
	conn := rpr.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", rpr.RedisKeyPrefix+playerKeyPrefix)
	if err != nil {
		return err
	}

	return nil
}
