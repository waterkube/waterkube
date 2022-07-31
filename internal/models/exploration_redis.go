package models

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

const (
	explorationKeyPrefix = "exploration:"
)

// RedisExplorationRepository type.
type RedisExplorationRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}

// Create function.
func (rer *RedisExplorationRepository) Create(exploration *Exploration) error {
	conn := rer.RedisPool.Get()
	defer conn.Close()

	err := conn.Send("MULTI")
	if err != nil {
		return err
	}

	key := rer.RedisKeyPrefix + explorationKeyPrefix + strings.ToLower(exploration.Grid)

	err = conn.Send(
		"HMSET", redis.Args{}.Add(key).AddFlat(exploration)...,
	)
	if err != nil {
		return err
	}

	err = conn.Send("EXPIRE", key, ExplorationTime.Seconds())
	if err != nil {
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}

	return nil
}

// Find function.
func (rer *RedisExplorationRepository) Find(grid *Grid) (*Exploration, error) {
	conn := rer.RedisPool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", rer.RedisKeyPrefix+explorationKeyPrefix+strings.ToLower(grid.Name)))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, ErrNoRecord
	}

	var exploration Exploration

	err = redis.ScanStruct(values, &exploration)
	if err != nil {
		return nil, err
	}

	return &exploration, nil
}
