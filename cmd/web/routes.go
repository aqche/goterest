package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *Goterest) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", g.home).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

	return r
}
