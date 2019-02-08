package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/start", Start)
	http.HandleFunc("/move", Move)
	http.HandleFunc("/end", End)
	http.HandleFunc("/ping", Ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, lib.LoggingHandler(http.DefaultServeMux))
}
