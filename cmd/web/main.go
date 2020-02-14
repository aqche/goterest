package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
