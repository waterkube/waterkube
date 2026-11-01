package web

import (
	"net/http"

	"github.com/justinas/alice"
	supportweb "github.com/petaki/support-go/web"
	"github.com/waterkube/waterkube/static"
)

func (a *app) routes() http.Handler {
	baseMiddleware := alice.New(supportweb.RecoverPanic(func(w http.ResponseWriter, r *http.Request, err error) {
		a.serverError(w, err)
	}))
	webMiddleware := alice.New(
		a.inertiaManager.Middleware,
	)

	mux := http.NewServeMux()
	mux.Handle("/", webMiddleware.ThenFunc(a.gameIndex))

	fileServer := supportweb.Static(a.appConfig.Debug, "./static/", static.Files)

	mux.Handle("/build/", fileServer)
	mux.Handle("/favicon.ico", fileServer)

	return baseMiddleware.Then(mux)
}
