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
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/service"
	"github.com/waterkube/waterkube/resources/views"
	"github.com/waterkube/waterkube/static"
)

// Serve function.
func Serve(appConfig *config.Config, redisPool *redis.Pool) {
	viteManager, inertiaManager, err := newViteAndInertiaManager(appConfig)
	if err != nil {
		cli.ErrorLog.Fatal(err)
	}

	webApp := &app{
		appConfig:      appConfig,
		infoLog:        cli.InfoLog,
		errorLog:       cli.ErrorLog,
		viteManager:    viteManager,
		inertiaManager: inertiaManager,
		gameManager:    service.GameManager(appConfig, redisPool),
	}

	srv := &http.Server{
		Addr:         appConfig.Addr,
		ErrorLog:     webApp.errorLog,
		Handler:      webApp.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	webApp.infoLog.Printf("Starting server on "+cli.Green("%s"), appConfig.Addr)

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
	inertiaManager.ShareFunc("isRunningHot", viteManager.IsRunningHot)
	inertiaManager.ShareFunc("asset", viteManager.Asset)
	inertiaManager.ShareFunc("css", viteManager.CSS)

	return viteManager, inertiaManager, nil
}
