package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/bapturp/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// Application struct to hold the application wide dependencies of the web application.
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// Using `localhost:8080` instead of just `:8080` to prevents MacOS to prompt
	// each time the app starts to allow the connection.
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	// Data Source Name is the configuration parameters for communicating with the database
	// Cleartext password, please don't do this on production environment.
	dsn := flag.String("dsn", "web:hello world@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	loggerErr := slog.NewLogLogger(slog.NewTextHandler(os.Stderr, nil), slog.LevelError)

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()

	// Initialize a new instance of our application struct, containing the dependencies.
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	// Create the web server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: loggerErr,
		Handler:  app.routes(),
	}

	logger.Info("starting server", "addr", *addr)
	err = server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
