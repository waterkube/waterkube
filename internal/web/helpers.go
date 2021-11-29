package web

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
)

func (a *app) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errorLog.Output(2, trace)

	if a.debug {
		http.Error(w, trace, http.StatusInternalServerError)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *app) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}

func (a *app) methodNotAllowed(w http.ResponseWriter, allow []string) {
	w.Header().Set("Allow", strings.Join(allow, ", "))
	a.clientError(w, http.StatusMethodNotAllowed)
}
