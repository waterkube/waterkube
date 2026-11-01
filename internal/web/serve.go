package web

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/cli"
	"github.com/petaki/support-go/vite"
	supportweb "github.com/petaki/support-go/web"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/resources/views"
	"github.com/waterkube/waterkube/static"
)

// Serve function.
func Serve(appConfig *config.Config, redisPool *redis.Pool) error {
	viteManager, inertiaManager, err := newViteAndInertiaManager(appConfig)
	if err != nil {
		return err
	}

	webApp := &app{
		appConfig:      appConfig,
		infoLog:        cli.InfoLog,
		errorLog:       cli.ErrorLog,
		viteManager:    viteManager,
		inertiaManager: inertiaManager,
		redisPool:      redisPool,
	}

	srv := supportweb.NewServer(appConfig.Addr, webApp.routes(), webApp.errorLog)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	return srv.RunAndShutdown(ctx, 5*time.Second)
}

func newViteAndInertiaManager(appConfig *config.Config) (*vite.Vite, *inertia.Inertia, error) {
	var viteManager *vite.Vite

	if appConfig.Debug {
		viteManager = vite.New("static", "build")
	} else {
		viteManager = vite.New("static", "build", static.Files)
	}

	version, err := viteManager.ManifestHash()
	if err != nil {
		return nil, nil, err
	}

	inertiaManager := inertia.New(appConfig.URL, "app.gohtml", version, views.Templates)
	inertiaManager.Share("title", "Waterkube")

	for name, helper := range viteManager.Helpers() {
		inertiaManager.ShareFunc(name, helper)
	}

	return viteManager, inertiaManager, nil
}
