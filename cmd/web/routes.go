package main

import "net/http"

// Request multiplexer containing the routes of our web server.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir func is relation to the project.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the  mux.Handler func to register the file server as the handler for
	// URL paths that start with "/static/" before the request reaches the file
	// server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/favicon.ico", app.favicon)

	return mux
}
