package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *goterest) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", g.home).Methods("GET")
	r.HandleFunc("/create", g.createForm).Methods("GET")
	r.HandleFunc("/create", g.create).Methods("POST")
	r.HandleFunc("/delete", g.delete).Methods("POST")
	r.HandleFunc("/login", g.loginForm).Methods("GET")
	r.HandleFunc("/login", g.login).Methods("POST")
	r.HandleFunc("/logout", g.logout).Methods("POST")
	r.HandleFunc("/register", g.registerForm).Methods("GET")
	r.HandleFunc("/register", g.register).Methods("POST")
	r.HandleFunc("/user/{username}", g.user).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

	return r
}
