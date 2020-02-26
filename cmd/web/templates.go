package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/aqche/goterest/pkg/models"
)

type templateData struct {
	Title string
	Pins  []models.Pin
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

func (g *goterest) renderTemplate(w http.ResponseWriter, name string, td templateData) {
	tmpl, ok := g.templates[name]
	if !ok {
		w.Write([]byte("invalid template"))
		return
	}

	err := tmpl.Execute(w, td)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
