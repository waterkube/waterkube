package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/cli"
	"github.com/petaki/support-go/vite"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
	"github.com/waterkube/waterkube/resources/views"
	"github.com/waterkube/waterkube/static"
)

// Serve function.
func Serve(debug bool, addr, url, redisKeyPrefix string, redisPool *redis.Pool) {
	viteManager, inertiaManager, err := newViteAndInertiaManager(debug, url)
	if err != nil {
		cli.ErrorLog.Fatal(err)
	}

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

	webApp := &app{
		debug:          debug,
		url:            url,
		infoLog:        cli.InfoLog,
		errorLog:       cli.ErrorLog,
		viteManager:    viteManager,
		inertiaManager: inertiaManager,
		gameManager: game.New(
			explorationRepository,
			gridRepository,
			playerRepository,
		),
		explorationRepository: explorationRepository,
		gridRepository:        gridRepository,
		playerRepository:      playerRepository,
	}

	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     webApp.errorLog,
		Handler:      webApp.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	webApp.infoLog.Printf("Starting server on "+cli.Green("%s"), addr)

	go func() {
		err = srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			webApp.errorLog.Fatal(err)
		}
	}()

	<-done
	webApp.infoLog.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		webApp.errorLog.Fatal(err)
	}

	webApp.infoLog.Print("Server exited properly")
}

func newViteAndInertiaManager(debug bool, url string) (*vite.Vite, *inertia.Inertia, error) {
	var viteManager *vite.Vite
	var version string
	var err error

	if debug {
		viteManager = vite.New("static", "build")
	} else {
		viteManager = vite.NewWithFS("static", "build", static.Files)
	}

	version, err = viteManager.ManifestHash()
	if err != nil {
		return nil, nil, err
	}

	inertiaManager := inertia.NewWithFS(url, "app.gohtml", version, views.Templates)
	inertiaManager.Share("title", "Waterkube")
	inertiaManager.ShareFunc("isRunningHot", viteManager.IsRunningHot)
	inertiaManager.ShareFunc("asset", viteManager.Asset)
	inertiaManager.ShareFunc("css", viteManager.CSS)

	return viteManager, inertiaManager, nil
}
