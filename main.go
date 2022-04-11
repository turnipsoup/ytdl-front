package main

//
// Handles the web routes and embeded files, everything else is in its own package
//

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

// Files embedded into the binary for easy deployment (static stuff)

//go:embed web/index.html
var webIndex string

// Main

func main() {
	log.Print("Starting application")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request recieved on /")
		fmt.Fprintln(w, webIndex)
	})

	// listen to port
	log.Print("Listening on port 5050")

	http.ListenAndServe(":5050", nil)

}
