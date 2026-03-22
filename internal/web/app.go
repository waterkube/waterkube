package web

import (
	"log"

	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/vite"
	"github.com/waterkube/waterkube/internal/config"
	"github.com/waterkube/waterkube/internal/game"
)

type app struct {
	appConfig      *config.Config
	errorLog       *log.Logger
	infoLog        *log.Logger
	viteManager    *vite.Vite
	inertiaManager *inertia.Inertia
	gameManager    *game.Game
}
