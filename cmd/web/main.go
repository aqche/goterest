package main

import (
	"fmt"
	"log"
	"net/http"
)

type Goterest struct{}

func main() {
	goterest := &Goterest{}

	server := &http.Server{
		Addr:    ":4000",
		Handler: goterest.routes(),
	}

	fmt.Printf("Starting server...")
	err := server.ListenAndServe()
	log.Fatal(err)
}
