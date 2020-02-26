package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type goterest struct {
	templates map[string]*template.Template
}

func main() {
	app := &goterest{}

	err := app.loadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	fmt.Printf("Starting server...")
	err = server.ListenAndServe()
	log.Fatal(err)
}
