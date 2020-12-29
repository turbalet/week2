package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome"))
}

func snippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Displaying snippet"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippet)
	err:= http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
