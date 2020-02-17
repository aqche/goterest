package main

import (
	"net/http"
)

func (g *Goterest) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
