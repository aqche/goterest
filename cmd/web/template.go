package main

import "github.com/aqche/goterest/pkg/models"

type templateData struct {
	Title string
	Pins  []models.Pin
}
