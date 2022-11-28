package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup server and dependencies.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := &Database{}

	server := &Server{
		infoLog:  infoLog,
		errorLog: errorLog,
		users:    db,
		notes:    db,
	}

	// 🔥 tomorrow ...
	//   - html templates
	//   - error logger ✔️
	//   - auto reloading?
	//   - database?
	//   - git repo? ✔️

	// Start server.
	addr := ":4000"
	infoLog.Printf("Starting server on %s\n", addr)

	err := http.ListenAndServe(addr, server.Handler())
	errorLog.Fatal(err)
}
