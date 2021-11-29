package web

import "net/http"

func (a *app) gameIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)

		return
	}

	if r.Method != "GET" {
		a.methodNotAllowed(w, []string{"GET"})

		return
	}

	err := a.inertiaManager.Render(w, r, "game/Index", nil)
	if err != nil {
		a.serverError(w, err)
	}
}
