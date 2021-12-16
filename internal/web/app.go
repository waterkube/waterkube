package web

import (
	"github.com/waterkube/waterkube/internal/models"
	"log"

	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/mix"
)

type app struct {
	debug                 bool
	url                   string
	errorLog              *log.Logger
	infoLog               *log.Logger
	mixManager            *mix.Mix
	inertiaManager        *inertia.Inertia
	explorationRepository models.ExplorationRepository
	gridRepository        models.GridRepository
	playerRepository      models.PlayerRepository
}
