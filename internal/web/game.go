package web

import (
	"net/http"

	"github.com/waterkube/waterkube/internal/game"
)

func (a *app) gameIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)

		return
	}

	if r.Method != "GET" {
		a.methodNotAllowed(w, []string{"GET"})

		return
	}

	err := a.gameManager.MapLoad()
	if err != nil {
		a.notFound(w)

		return
	}

	freeBoat, freeDiver, freeSubmarine := a.gameManager.FreeUnits()

	err = a.inertiaManager.Render(w, r, "game/Index", map[string]interface{}{
		"player":              a.gameManager.Player,
		"freeBoat":            freeBoat,
		"freeDiver":           freeDiver,
		"freeSubmarine":       freeSubmarine,
		"progress":            a.gameManager.MapProgress(),
		"cols":                game.Cols,
		"rows":                game.Rows,
		"grids":               a.gameManager.Grids,
		"explorations":        a.gameManager.Explorations(),
		"discoveredArtifacts": a.gameManager.DiscoveredArtifacts(),
		"donatedArtifacts":    a.gameManager.DonatedArtifacts(),
	})
	if err != nil {
		a.serverError(w, err)
	}
}
