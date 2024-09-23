package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Application struct to hold the application wide dependencies of the web application.
type application struct {
	logger *slog.Logger
}

func main() {
	// Using `localhost:8080` instead of just `:8080` to prevents MacOS to prompt
	// each time the app starts to allow the connection.
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Initialize a new instance of our application struct, containing the dependencies.
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
