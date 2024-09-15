package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// The home handler is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	// Read and parse the template
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	// Write the parsed template as response
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}

// favicon handler is a method of *application
func (app *application) favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/img/favicon.ico")
}

// snippetView handler, method of *application
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables to holding dummy data. We'll remove these later
	// on during the build.
	title := "O snail"
	content := "O snail\nCimb Mount Fuji,\nBut slowly!\n\nKobayashi Issa"
	expires := 7

	// pass the data to the SnipppetModel.Insert() method, receiving the ID of
	// the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
