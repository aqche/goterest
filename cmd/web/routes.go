package main

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func (g *goterest) routes(csrfKey string, secure bool) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", g.home).Methods("GET")

	r.Handle("/create", g.requireAuthenticated(http.HandlerFunc(g.createForm))).Methods("GET")
	r.Handle("/create", g.requireAuthenticated(http.HandlerFunc(g.create))).Methods("POST")

	r.Handle("/delete/{id}", g.requireAuthenticated(http.HandlerFunc(g.delete))).Methods("POST")

	r.Handle("/login", g.requireUnauthenticated(http.HandlerFunc(g.loginForm))).Methods("GET")
	r.Handle("/login", g.requireUnauthenticated(http.HandlerFunc(g.login))).Methods("POST")

	r.Handle("/logout", g.requireAuthenticated(http.HandlerFunc(g.logout))).Methods("POST")

	r.Handle("/register", g.requireUnauthenticated(http.HandlerFunc(g.registerForm))).Methods("GET")
	r.Handle("/register", g.requireUnauthenticated(http.HandlerFunc(g.register))).Methods("POST")

	r.HandleFunc("/user/{username}", g.user).Methods("GET")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

	csrfMiddleware := csrf.Protect([]byte(csrfKey), csrf.Secure(secure))

	return csrfMiddleware(r)
}
