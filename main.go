package main

//
// Handles the web routes and embeded files,
// everything else is in its own package
//

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

// Files embedded into the binary for easy deployment (static stuff)

//go:embed web/index.html
var webIndex string

//go:embed web/static
var staticFiles embed.FS

// Main

func main() {
	log.Print("Starting application")

	// Get and handle static files
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.FS(staticFiles))))

	// Getting the root of the application

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request recieved on /")
		log.Print(r.UserAgent())

		// Ensure that the browser knows it is HTML
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, webIndex)
	})

	// Listen on port
	log.Print("Listening on port 5050")

	http.ListenAndServe(":5050", nil)

}
