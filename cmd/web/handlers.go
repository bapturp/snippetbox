package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

/*
Handlers, responsible for executing the application logic and writing HTTP
response headers and bodies.
*/

// The home handler is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Demonstrate the use of specific header by adding `Server: Go`` to the
	// response header
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	// Read and parse the template
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		app.serverError(w, r, err)
		return
	}

	// Write the parsed template as response
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		app.serverError(w, r, err)
	}
}

// favicon handler is a method of *application
func (app *application) favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/img/favicon.ico")
}

// snippetView handler, method of *application
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Use r.PathValue() to go the value of {id} extracted from the request path
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use w.WriteHeader() method to send a 201 Created status code

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
