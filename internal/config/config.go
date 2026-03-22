package config

import (
	"os"

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
	debug := command.FlagSet().Bool("debug", false, "Application Debug Mode")
	addr := command.FlagSet().String("addr", os.Getenv("APP_ADDR"), "Application Address")
	url := command.FlagSet().String("url", os.Getenv("APP_URL"), "Application URL")
	redisURL := command.FlagSet().String("redis-url", os.Getenv("REDIS_URL"), "Redis URL")
	redisKeyPrefix := command.FlagSet().String("redis-key-prefix", os.Getenv("REDIS_KEY_PREFIX"), "Redis Key Prefix")

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
