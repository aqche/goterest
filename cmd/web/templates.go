package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/aqche/goterest/pkg/forms"
	"github.com/aqche/goterest/pkg/models"
	"github.com/gorilla/csrf"
)

type templateData struct {
	Flashes   []interface{}
	User      interface{}
	Title     string
	Pins      []*models.Pin
	Form      *forms.Form
	CSRFField template.HTML
}

func (g *goterest) loadTemplates() error {
	templates := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/*.page.tmpl")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.ParseFiles(page, "./ui/html/base.layout.tmpl")
		if err != nil {
			return err
		}

		templates[name] = tmpl
	}

	g.templates = templates

	return nil
}

func (g *goterest) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	tmpl, ok := g.templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("invalid template %q", name), http.StatusInternalServerError)
		return
	}

	session, err := g.store.Get(r, "goterest")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	td.User = session.Values["user"]
	td.Flashes = session.Flashes()
	td.CSRFField = csrf.TemplateField(r)

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer := &bytes.Buffer{}

	err = tmpl.Execute(buffer, td)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buffer.WriteTo(w)
}
