package main

import (
	"html/template"
	"net/http"

	"github.com/aqche/goterest/pkg/models"
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
