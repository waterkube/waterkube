package cmd

import (
	"os"

	"github.com/petaki/support-go/cli"
	"github.com/waterkube/waterkube/internal/web"
)

// WebServe command.
func WebServe(group *cli.Group, command *cli.Command, arguments []string) int {
	debug := command.FlagSet().Bool("debug", false, "Application Debug Mode")
	addr := command.FlagSet().String("addr", os.Getenv("APP_ADDR"), "Application Address")
	url := command.FlagSet().String("url", os.Getenv("APP_URL"), "Application URL")

	redisURL, redisKeyPrefix := createRedisFlags(command)

	_, err := command.Parse(arguments)
	if err != nil {
		return command.PrintHelp(group)
	}

	redisPool := newRedisPool(*redisURL)
	defer redisPool.Close()

	web.Serve(*debug, *addr, *url, *redisKeyPrefix, redisPool)

	return cli.Success
}
