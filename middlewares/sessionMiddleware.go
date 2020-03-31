package middlewares

import (
	"net/http"
	"superMarket/repo"
)

type SessionMiddleware struct {
	NextTask func(http.ResponseWriter, *http.Request)
}

func (a *SessionMiddleware) HandleFunc(w http.ResponseWriter, r *http.Request) {
	goSessionId := r.Header.Get("goSessionId")
	if goSessionId == "" {
		http.Error(w, "sessionId is not exists", 500)
		return
	}
	key, err := repo.GetExpireKey(goSessionId)
	if err != nil || key == "" {
		http.Error(w, err.Error(), 500)
		return
	}
	if a.NextTask != nil {
		a.NextTask(w, r)
	}
}
