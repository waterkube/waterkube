package web

import (
	"log"

	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/mix"
	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
)

type app struct {
	debug                 bool
	url                   string
	errorLog              *log.Logger
	infoLog               *log.Logger
	mixManager            *mix.Mix
	inertiaManager        *inertia.Inertia
	gameManager           *game.Game
	explorationRepository models.ExplorationRepository
	gridRepository        models.GridRepository
	playerRepository      models.PlayerRepository
}
