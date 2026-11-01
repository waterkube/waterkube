package config

import (
	"github.com/petaki/support-go/cli"
)

// Config type.
type Config struct {
	Debug          bool
	Addr           string
	URL            string
	RedisURL       string
	RedisKeyPrefix string
}

// NewConfig function.
func NewConfig(command *cli.Command, arguments []string) (*Config, error) {
	debug := command.Bool("debug", "APP_DEBUG", false, "Application Debug Mode")
	addr := command.String("addr", "APP_ADDR", ":4000", "Application Address")
	url := command.String("url", "APP_URL", "", "Application URL")
	redisURL := command.Secret("redis-url", "REDIS_URL", "Redis URL")
	redisKeyPrefix := command.String("redis-key-prefix", "REDIS_KEY_PREFIX", "", "Redis Key Prefix")

	_, err := command.Parse(arguments)
	if err != nil {
		return nil, err
	}

	return &Config{
		Debug:          *debug,
		Addr:           *addr,
		URL:            *url,
		RedisURL:       *redisURL,
		RedisKeyPrefix: *redisKeyPrefix,
	}, nil
}
