package main

import (
	"fmt"
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

	td := templateData{
		Title: "Pins",
		Pins:  pins,
	}

	g.renderTemplate(w, "pins.page.tmpl", td)
}

func (g *goterest) createForm(w http.ResponseWriter, r *http.Request) {
	td := templateData{
		Title: "Create Pin",
	}

	g.renderTemplate(w, "create.page.tmpl", td)
}

func (g *goterest) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error parsing form"))
	}

	image := r.PostForm.Get("image")

	fmt.Fprintf(w, "image: %s", image)
}

func (g *goterest) delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func (g *goterest) loginForm(w http.ResponseWriter, r *http.Request) {
	td := templateData{
		Title: "Log In",
	}

	g.renderTemplate(w, "login.page.tmpl", td)
}

func (g *goterest) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error parsing form"))
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	fmt.Fprintf(w, "username: %s, password: %s", username, password)
}

func (g *goterest) logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}

func (g *goterest) registerForm(w http.ResponseWriter, r *http.Request) {
	td := templateData{
		Title: "Register",
	}

	g.renderTemplate(w, "register.page.tmpl", td)
}

func (g *goterest) register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error parsing form"))
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	confirmPassword := r.PostForm.Get("confirm-password")

	fmt.Fprintf(w, "username: %s, password: %s, confirmPassword: %s", username, password, confirmPassword)
}

func (g *goterest) user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	td := templateData{
		Title: fmt.Sprintf("%s's Pins", username),
	}

	g.renderTemplate(w, "pins.page.tmpl", td)
}
