package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/bapturp/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

// Application struct to hold the application wide dependencies of the web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
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

	// Use log.New() to create a logger for writing information messages. This
	// takes three parameters: the destination to write the logs to (os.Stdout),
	// a string prefix for message (INFO followed by a tab), and flags to
	// indicate what additional information to include (local date and time).
	// Note that the flags are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use
	// stderr as the detination and use the log.Lshortfile flag to include
	// the relevant file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new instance of our application struct, containing the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// Create the web server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
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
