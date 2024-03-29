package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aqche/goterest/pkg/forms"
	"github.com/aqche/goterest/pkg/models"
	"github.com/gorilla/mux"
)

func (g *goterest) home(w http.ResponseWriter, r *http.Request) {
	pins, err := g.pins.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	form.ValidateImgExtension("image-url")

	if form.ContainsErrors() {
		g.renderTemplate(w, r, "create.page.tmpl", &templateData{
			Title: "Create Pin",
			Form:  form,
		})
		return
	}

	session, err := g.store.Get(r, "goterest")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := session.Values["user"]

	err = g.pins.Create(form.Values.Get("image-url"), username.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash("Successfully created pin.")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (g *goterest) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pin, err := g.pins.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrPinNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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

	if session.Values["user"] != pin.Username {
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = g.pins.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash("Successfully deleted pin.")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	err = g.users.ValidatePassword(strings.ToLower(form.Values.Get("username")), form.Values.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) || errors.Is(err, models.ErrInvalidPassword) {
			form.Errors["login"] = append(form.Errors["login"], "Invalid username or password.")
			g.renderTemplate(w, r, "login.page.tmpl", &templateData{
				Title: "Log In",
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

	session.Values["user"] = strings.ToLower(form.Values.Get("username"))
	session.AddFlash("Successfully logged in.")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (g *goterest) logout(w http.ResponseWriter, r *http.Request) {
	session, err := g.store.Get(r, "goterest")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = nil
	session.AddFlash("Successfully logged out.")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	username := strings.ToLower(vars["username"])

	_, err := g.users.GetByUsername(username)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	pins, err := g.pins.GetAllByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.renderTemplate(w, r, "pins.page.tmpl", &templateData{
		Title: fmt.Sprintf("%s's Pins", strings.Title(username)),
		Pins:  pins,
	})
}
