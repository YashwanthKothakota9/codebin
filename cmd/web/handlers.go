package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//family of funcs for safely parsing and rendering HTML templates
)

func home(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Server","Go")

	// Use the template.ParseFiles() function to read the template file into a
    // template set. If there's an error, we log the detailed error message, use
    // the http.Error() function to send an Internal Server Error response to the
    // user, and then return from the handler so no subsequent code is executed.
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err!=nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server error",http.StatusInternalServerError)
		return
	}

	// Then we use the Execute() method on the template set to write the
    // template content as the response body. The last parameter to Execute()
    // represents any dynamic data that we want to pass in, which for now we'll
    // leave as nil.

	err = ts.Execute(w,nil)
	if err!=nil{
		log.Print(err.Error())
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
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