package web

import (
	"errors"
	"net/http"

	"github.com/waterkube/waterkube/internal/game"
	"github.com/waterkube/waterkube/internal/models"
	"github.com/waterkube/waterkube/internal/service"
)

func (a *app) gameIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)

		return
	}

	if r.Method != http.MethodGet {
		a.methodNotAllowed(w, []string{http.MethodGet})

		return
	}

	gameManager := service.GameManager(a.appConfig, a.redisPool)

	err := gameManager.MapLoad()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)

			return
		}

		a.serverError(w, err)

		return
	}

	freeBoat, freeDiver, freeSubmarine := gameManager.FreeUnits()

	err = a.inertiaManager.Render(w, r, "game/Index", map[string]any{
		"player":              gameManager.Player,
		"freeBoat":            freeBoat,
		"freeDiver":           freeDiver,
		"freeSubmarine":       freeSubmarine,
		"progress":            gameManager.MapProgress(),
		"cols":                game.Cols,
		"rows":                game.Rows,
		"grids":               gameManager.Grids,
		"explorations":        gameManager.Explorations,
		"discoveredArtifacts": gameManager.DiscoveredArtifacts(),
		"donatedArtifacts":    gameManager.DonatedArtifacts(),
	})
	if err != nil {
		a.serverError(w, err)
	}
}
