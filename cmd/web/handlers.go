package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aqche/goterest/pkg/forms"
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

	g.renderTemplate(w, r, "pins.page.tmpl", &templateData{
		Title: "Pins",
		Pins:  pins,
	})
}

func (g *goterest) createForm(w http.ResponseWriter, r *http.Request) {
	g.renderTemplate(w, r, "create.page.tmpl", &templateData{
		Title: "Create Pin",
		Form:  forms.NewForm(nil),
	})
}

func (g *goterest) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.NewForm(r.PostForm)
	form.ValidateRequired("image-url")
	form.ValidateURL("image-url")

	if form.ContainsErrors() {
		g.renderTemplate(w, r, "create.page.tmpl", &templateData{
			Title: "Create Pin",
			Form:  form,
		})
		return
	}

	fmt.Fprintf(w, "image: %s", form.Values.Get("image-url"))
}

func (g *goterest) delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func (g *goterest) loginForm(w http.ResponseWriter, r *http.Request) {
	g.renderTemplate(w, r, "login.page.tmpl", &templateData{
		Title: "Log In",
		Form:  forms.NewForm(nil),
	})
}

func (g *goterest) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.NewForm(r.PostForm)
	form.ValidateRequired("username")
	form.ValidateRequired("password")

	if form.ContainsErrors() {
		g.renderTemplate(w, r, "login.page.tmpl", &templateData{
			Title: "Log In",
			Form:  form,
		})
		return
	}

	fmt.Fprintf(w, "username: %s, password: %s", form.Values.Get("username"), form.Values.Get("password"))
}

func (g *goterest) logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}

func (g *goterest) registerForm(w http.ResponseWriter, r *http.Request) {
	g.renderTemplate(w, r, "register.page.tmpl", &templateData{
		Title: "Register",
		Form:  forms.NewForm(nil),
	})
}

func (g *goterest) register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.NewForm(r.PostForm)
	form.ValidateRequired("username")
	form.ValidateRequired("password")
	form.ValidateRequired("confirm-password")
	form.ValidateMinLength("password", 6)
	form.ValidateMatch("confirm-password", "password")

	if form.ContainsErrors() {
		g.renderTemplate(w, r, "register.page.tmpl", &templateData{
			Title: "Register",
			Form:  form,
		})
		return
	}

	err = g.users.Create(strings.ToLower(form.Values.Get("username")), form.Values.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrUserExists) {
			form.Errors["username"] = append(form.Errors["username"], "The username is already taken.")
			g.renderTemplate(w, r, "register.page.tmpl", &templateData{
				Title: "Register",
				Form:  form,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	session, err := g.store.Get(r, "goterest")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash("Successfully registered.")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (g *goterest) user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	g.renderTemplate(w, r, "pins.page.tmpl", &templateData{
		Title: fmt.Sprintf("%s's Pins", username),
		Pins:  []models.Pin{},
	})
}
