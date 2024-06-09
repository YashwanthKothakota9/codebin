package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}",home)
	mux.HandleFunc("GET /code/view/{id}",codeView)
	mux.HandleFunc("GET /code/create",codeCreate)
	mux.HandleFunc("POST /code/create",codeCreatePost)

	log.Println("Server running on PORT: 4000")

	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}