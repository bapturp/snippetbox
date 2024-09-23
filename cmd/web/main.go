package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Application struct to hold the application wide dependencies of the web application.
type application struct {
	logger *slog.Logger
}

func main() {
	// Using `localhost:8080` instead of just `:8080` to prevents MacOS to prompt
	// each time the app starts to allow the connection.
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	dsn := flag.String("dsn", "web:hello world@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Initialize a new instance of our application struct, containing the dependencies.
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
