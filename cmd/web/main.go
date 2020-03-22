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
	db, err := sql.Open("postgres", "user=goterest password=pass dbname=goterest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := &goterest{
		store: sessions.NewCookieStore([]byte("test-session-key")),
		users: postgres.UserModel{DB: db},
		pins:  postgres.PinModel{DB: db},
	}

	err = app.loadTemplates()
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
