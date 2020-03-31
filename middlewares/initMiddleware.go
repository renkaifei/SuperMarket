package middlewares

import (
	"net/http"
)

type InitMiddleware struct {
	NextTask func(http.ResponseWriter, *http.Request)
}

func (a *InitMiddleware) HandleFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if a.NextTask != nil {
		a.NextTask(w, r)
	}
}
