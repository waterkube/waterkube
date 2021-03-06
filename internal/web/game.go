package web

import (
	"net/http"

	"github.com/waterkube/waterkube/internal/artifact"
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

	var artifacts []interface{}

	for name, ar := range artifact.ShallowUnique {
		artifacts = append(artifacts, map[string]interface{}{
			"name":  name,
			"price": ar.Price,
		})
	}

	err = a.inertiaManager.Render(w, r, "game/Index", map[string]interface{}{
		"cols":      game.Cols,
		"rows":      game.Rows,
		"grids":     a.gameManager.Grids,
		"artifacts": artifacts,
	})
	if err != nil {
		a.serverError(w, err)
	}
}
