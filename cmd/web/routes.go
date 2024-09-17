package main

import "net/http"

// Router or servemux in Go terminoligy
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir func is relation to the project.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the  mux.Handler func to register the file server as the handler for
	// URL paths that start with "/static/" before the request reaches the file
	// server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// `{$}` restricts to the route to match only `/` (a single slash) and nothing after`
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.HandleFunc("GET /favicon.ico", app.favicon)

	return mux
}
