package main

import (
	"html/template"
	"net/http"

	"github.com/aqche/goterest/pkg/models"
	"github.com/gorilla/mux"
)

func (g *goterest) home(w http.ResponseWriter, r *http.Request) {
	pins := []models.Pin{
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/200/300",
			Username: "Abby",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/300/200",
			Username: "Brad",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/200/400",
			Username: "Cloe",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/400/200",
			Username: "Dan",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/300/300",
			Username: "Erica",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/300/400",
			Username: "Fiona",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/400/300",
			Username: "George",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/150/150",
			Username: "Herald",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/500/300",
			Username: "Irene",
		},
		models.Pin{
			ImageURL: "https://picsum.photos/seed/picsum/300/500",
			Username: "Joe",
		},
	}

	files := []string{"./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = tmpl.Execute(w, pins)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func (g *goterest) createForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createForm"))
}

func (g *goterest) create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func (g *goterest) delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func (g *goterest) loginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("loginForm"))
}

func (g *goterest) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func (g *goterest) logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}

func (g *goterest) registerForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("registerForm"))
}

func (g *goterest) register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register"))
}

func (g *goterest) user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(username))
}
