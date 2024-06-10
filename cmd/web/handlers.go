package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	//family of funcs for safely parsing and rendering HTML templates
)

// Change the signature of the home handler so it is defined as a method against
// /*application.

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we use ... to pass the contents
	// of the files slice as variadic arguments.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler is now a method against the application
		// struct it can access its fields, including the structured logger. We'll
		// use this to create a log entry at Error level containing the error
		// message, also including the request method and URI as attributes to
		// assist with debugging.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// And we also need to update the code here to use the structured logger
		// too.
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Change the signature of the codeView handler so it is defined as a method
// against *application.
func (app *application) codeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific code with ID %d...", id)
}

// Change the signature of the codeCreate handler so it is defined as a method
// against *application.
func (app *application) codeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying a form for creating new code snippet..."))
}

// Change the signature of the codeCreatePost handler so it is defined as a method
// against *application.
func (app *application) codeCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new code snippet..."))
}
