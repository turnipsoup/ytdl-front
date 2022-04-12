package main

//
// Handles the web routes and embeded files,
// everything else is in its own package
//

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Files embedded into the binary for easy deployment (static stuff)

//go:embed web/index.html
var webIndex string

//go:embed web/static
var staticFiles embed.FS

// Reading the files present in the storage area
func getCurrentlyDownloading() {
	rootDirectory := "/data/shared/music/YouTube/"
	fmt.Println(rootDirectory)

	genres, err := ioutil.ReadDir(rootDirectory)

	if err != nil {
		log.Println("Could not read root directory")
		log.Println(err)
	}

	for t := range genres {
		fmt.Println(genres[t].Name())
	}
}

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

	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		// parse request payload
		// Content-Type: application/x-www-form-urlencoded
		log.Print("Submission recieved on /push")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		ytId := r.Form["yt-id"][0]

		log.Print(fmt.Sprintf("Processing request for YT-ID %s", ytId))

		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/current", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Fetching current history")
		getCurrentlyDownloading()

	})

	// Listen on port
	log.Print("Listening on port 5050")

	err := http.ListenAndServe(":5050", nil)

	if err != nil {
		log.Println("Could not start application")
		fmt.Println(err)
	}
}
