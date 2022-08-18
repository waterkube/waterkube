package models

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

const (
	gridKeyPrefix = "grid:"
)

// RedisGridRepository type.
type RedisGridRepository struct {
	RedisPool      *redis.Pool
	RedisKeyPrefix string
}

// CreateOrUpdate function.
func (rgr *RedisGridRepository) CreateOrUpdate(grid *Grid) error {
	conn := rgr.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do(
		"HSET", redis.Args{}.Add(rgr.RedisKeyPrefix+gridKeyPrefix+strings.ToLower(grid.Name)).AddFlat(grid)...,
	)
	if err != nil {
		return err
	}

	return nil
}

// Find function.
func (rgr *RedisGridRepository) Find(name string) (*Grid, error) {
	conn := rgr.RedisPool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", rgr.RedisKeyPrefix+gridKeyPrefix+strings.ToLower(name)))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, ErrNoRecord
	}

	var grid Grid

	err = redis.ScanStruct(values, &grid)
	if err != nil {
		return nil, err
	}

	return &grid, nil
}

// Delete function.
func (rgr *RedisGridRepository) Delete(name string) error {
	conn := rgr.RedisPool.Get()
	defer conn.Close()

	err := conn.Send("MULTI")
	if err != nil {
		return err
	}

	err = conn.Send("DEL", rgr.RedisKeyPrefix+explorationKeyPrefix+strings.ToLower(name))
	if err != nil {
		return err
	}

	err = conn.Send("DEL", rgr.RedisKeyPrefix+gridKeyPrefix+strings.ToLower(name))
	if err != nil {
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}

	return nil
}
