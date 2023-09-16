package main

import (
	"net/http"

	lg "github.com/walkersumida/go-api-server/internal/pkg/log"
)

type middleware func(http.Handler) http.Handler

func middlewareHTTPLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := lg.NewLogger()
		log.Infow("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func setupHTTPHandler(h http.Handler, ms []middleware) http.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i](h)
	}

	return h
}
