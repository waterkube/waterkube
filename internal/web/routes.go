package web

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/petaki/waterkube/static"
)

func (a *app) routes() http.Handler {
	baseMiddleware := alice.New(a.recoverPanic)
	webMiddleware := alice.New(
		a.inertiaManager.Middleware,
	)

	mux := http.NewServeMux()
	mux.Handle("/", webMiddleware.ThenFunc(a.gameIndex))

	var fileServer http.Handler

	if a.debug {
		fileServer = http.FileServer(http.Dir("./static/"))
	} else {
		staticFS := http.FS(static.Files)
		fileServer = http.FileServer(staticFS)
	}

	mux.Handle("/css/", fileServer)
	mux.Handle("/images/", fileServer)
	mux.Handle("/js/", fileServer)
	mux.Handle("/favicon.ico", fileServer)

	return baseMiddleware.Then(mux)
}
