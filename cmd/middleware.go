package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GET(path string, next httprouter.Handle) {
	router.GET(path, corsMiddleware(next))
}

func corsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}
