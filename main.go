package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating snippet"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
