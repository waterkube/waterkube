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

// Create function.
func (rpr *RedisPlayerRepository) Create(player *Player) error {
	conn := rpr.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do(
		"HMSET", redis.Args{}.Add(rpr.RedisKeyPrefix+playerKeyPrefix).AddFlat(player)...,
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

// Update function.
func (rpr *RedisPlayerRepository) Update(newPlayer *Player) error {
	conn := rpr.RedisPool.Get()
	defer conn.Close()

	player, err := rpr.Find()
	if err != nil {
		return err
	}

	player.Money = newPlayer.Money
	player.BoatCount = newPlayer.BoatCount
	player.DiverCount = newPlayer.DiverCount
	player.SubmarineCount = newPlayer.SubmarineCount

	_, err = conn.Do(
		"HMSET", redis.Args{}.Add(rpr.RedisKeyPrefix+playerKeyPrefix).AddFlat(player)...,
	)
	if err != nil {
		return err
	}

	return nil
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
