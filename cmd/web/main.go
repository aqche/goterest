package main

import (
	"fmt"
	"log"
	"net/http"
)

type goterest struct{}

func main() {
	app := &goterest{}

	server := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	fmt.Printf("Starting server...")
	err := server.ListenAndServe()
	log.Fatal(err)
}
