package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *Goterest) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", g.home)
	return r
}
