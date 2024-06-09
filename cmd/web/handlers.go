package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Server","Go")
	w.Write([]byte("Hello from CODEBIN!"))
}

func codeView(w http.ResponseWriter,r *http.Request){
	id,err := strconv.Atoi(r.PathValue("id"))
	if err!=nil || id<1{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Display a specific code with ID %d...",id)
}

func codeCreate(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Displaying a form for creating new code snippet..."))
}

func codeCreatePost(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new code snippet..."))
}