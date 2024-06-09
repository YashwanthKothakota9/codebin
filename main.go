package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func codeView(w http.ResponseWriter,r *http.Request){

	id, err := strconv.Atoi(r.PathValue("id"))
	if err!=nil || id<1 {
		http.NotFound(w,r)
		return
	}

	msg := fmt.Sprintf("Display a specific code snippet with ID %d...",id)
	w.Write([]byte(msg))
}

func codeCreate(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Display a form for creating new code"))
}

// Add a snippetCreatePost handler function.
func codeCreatePost(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Save a new code snippet..."))
}


func home(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello from Codebin!"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern.
    mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}",home)
	mux.HandleFunc("GET /code/view/{id}",codeView)//Add the {id} wild card segment.
	mux.HandleFunc("GET /code/create",codeCreate)

	//restricted to POST requests only
	mux.HandleFunc("POST /code/create",codeCreatePost)

	log.Println("Starting server on: 4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
    // two parameters: the TCP network address to listen on (in this case ":4000")
    // and the servemux we just created. If http.ListenAndServe() returns an error
    // we use the log.Fatal() function to log the error message and exit. Note
    // that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}