package web

import (
	"log"

	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/vite"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
)

type app struct {
	debug                 bool
	url                   string
	errorLog              *log.Logger
	infoLog               *log.Logger
	viteManager           *vite.Vite
	inertiaManager        *inertia.Inertia
	gameManager           *game.Game
	explorationRepository models.ExplorationRepository
	gridRepository        models.GridRepository
	playerRepository      models.PlayerRepository
}
