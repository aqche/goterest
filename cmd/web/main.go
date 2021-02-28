package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aqche/goterest/pkg/models/postgres"
	"github.com/gorilla/sessions"

	_ "github.com/lib/pq"
)

type goterest struct {
	store     *sessions.CookieStore
	templates map[string]*template.Template
	users     postgres.UserModel
	pins      postgres.PinModel
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := &goterest{
		store: sessions.NewCookieStore([]byte(config.SessionKey)),
		users: postgres.UserModel{DB: db},
		pins:  postgres.PinModel{DB: db},
	}

	err = app.loadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: app.routes(config.CSRFKey, config.Secure),
	}

	fmt.Printf("Starting server...")
	err = server.ListenAndServe()
	log.Fatal(err)
}
