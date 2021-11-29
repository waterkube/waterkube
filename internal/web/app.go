package web

import (
	"log"

	"github.com/petaki/inertia-go"
	"github.com/petaki/support-go/mix"
)

type app struct {
	debug          bool
	url            string
	errorLog       *log.Logger
	infoLog        *log.Logger
	mixManager     *mix.Mix
	inertiaManager *inertia.Inertia
}
