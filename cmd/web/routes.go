package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (g *goterest) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", g.home).Methods("GET")
	r.Handle("/create", g.loginRequired(http.HandlerFunc(g.createForm))).Methods("GET")
	r.Handle("/create", g.loginRequired(http.HandlerFunc(g.create))).Methods("POST")
	r.Handle("/delete", g.loginRequired(http.HandlerFunc(g.delete))).Methods("POST")
	r.HandleFunc("/login", g.loginForm).Methods("GET")
	r.HandleFunc("/login", g.login).Methods("POST")
	r.Handle("/logout", g.loginRequired(http.HandlerFunc(g.logout))).Methods("POST")
	r.HandleFunc("/register", g.registerForm).Methods("GET")
	r.HandleFunc("/register", g.register).Methods("POST")
	r.HandleFunc("/user/{username}", g.user).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

	return r
}
