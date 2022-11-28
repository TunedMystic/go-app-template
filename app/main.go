package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP listen address")
	debug := flag.Bool("debug", true, "Run application in debug mode")
	flag.Parse()

	// Setup server and dependencies.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := &Database{}

	server := &Server{
		debug:    *debug,
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
	infoLog.Printf("Starting server on %s\n", *addr)

	err := http.ListenAndServe(*addr, server.Handler())
	errorLog.Fatal(err)
}
