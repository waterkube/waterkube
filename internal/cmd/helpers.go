package cmd

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
)

func createRedisFlags(command *cli.Command) (*string, *string) {
	redisURL := command.FlagSet().String("redis-url", os.Getenv("REDIS_URL"), "Redis URL")
	redisKeyPrefix := command.FlagSet().String("redis-key-prefix", os.Getenv("REDIS_KEY_PREFIX"), "Redis Key Prefix")

	return redisURL, redisKeyPrefix
}

func newGameManager(redisKeyPrefix string, redisPool *redis.Pool) *game.Game {
	explorationRepository := &models.RedisExplorationRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: redisKeyPrefix,
	}

	gridRepository := &models.RedisGridRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: redisKeyPrefix,
	}

	playerRepository := &models.RedisPlayerRepository{
		RedisPool:      redisPool,
		RedisKeyPrefix: redisKeyPrefix,
	}

	return game.New(
		explorationRepository,
		gridRepository,
		playerRepository,
	)
}

func newRedisPool(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
	}
}
