package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
)

// GameManager function.
func GameManager(appConfig *config.Config, redisPool *redis.Pool) *game.Game {
	explorationRepository := &models.RedisExplorationRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: appConfig.RedisKeyPrefix,
	}

	gridRepository := &models.RedisGridRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: appConfig.RedisKeyPrefix,
	}

	playerRepository := &models.RedisPlayerRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: appConfig.RedisKeyPrefix,
	}

	return game.New(
		explorationRepository,
		gridRepository,
		playerRepository,
	)
}
