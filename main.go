package main

//
// Handles the web routes and embeded files,
// everything else is in its own package
//

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"letseatlabs/ytdl-front/files"
	"letseatlabs/ytdl-front/yt"
)

// Read config.json

type Config struct {
	RootDirectory string `json:"storage_root"`
	DBLocation    string `json:"db_location"`
}

// Let's first read the `config.json` file
func getConfiguration() Config {
	content, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal("Error when opening configuration file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var config Config

	err = json.Unmarshal(content, &config)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return config
}

// Files embedded into the binary for easy deployment (static stuff)

//go:embed web/index.html
var webIndex string

//go:embed web/static
var staticFiles embed.FS

// Main

func main() {
	log.Print("Starting application")

	config := getConfiguration()
	// Initialize the DB "connection"
	yt.OpenDatabaseInit(config.DBLocation)
	yt.GetAllDownloads(config.DBLocation)

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

	// Submitting POST requests to download

	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		// parse request payload
		// Content-Type: application/x-www-form-urlencoded
		log.Print("Submission recieved on /push")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		ytId := yt.ParseYouTubeURL(r.Form["yt-id"][0])
		genre := r.Form["genre"][0]

		// Insert our record
		yt.InsertYTDLRecord(config.DBLocation, ytId, 0, 0, yt.CreateYTUrl(ytId), genre, "Active")

		log.Print(fmt.Sprintf("Processing request for YT-ID %s", ytId))
		go yt.DownloadVideoAudio(config.DBLocation, ytId, config.RootDirectory, genre)

		http.Redirect(w, r, "/", 302)
	})

	// Returns a list of currently downloading files
	http.HandleFunc("/current", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Fetching current history")

		var currentDownloads []string

		rows := yt.GetAllDownloads(config.DBLocation)

		for i := range rows {
			if rows[i].Status != "Done" {
				currentDownloads = append(currentDownloads, yt.RowToJSON(rows[i]))
			}

		}

		innerJSON := strings.Join(currentDownloads, ",")

		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, fmt.Sprintf("[%s]", innerJSON))

	})

	// Returns a list of all files in the db
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Fetching current history")

		var currentDownloads []string

		rows := yt.GetAllDownloads(config.DBLocation)

		for i := range rows {
			currentDownloads = append(currentDownloads, yt.RowToJSON(rows[i]))

		}

		innerJSON := strings.Join(currentDownloads, ",")

		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, fmt.Sprintf("[%s]", innerJSON))

	})

	// Returns a list of available genres
	http.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		genres := files.GetAllGenres(config.RootDirectory)
		log.Printf("Fetched list of genres: %s", genres)

		log.Println(genres)

		// Convert the []string to a JSON array
		genreJson := fmt.Sprintf("[\"%s\"]", strings.Join(genres, "\",\""))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, genreJson)

	})

	// Listen on port
	log.Print("Listening on port 5050")

	err := http.ListenAndServe(":5050", nil)

	if err != nil {
		log.Println("Could not start application")
		fmt.Println(err)
	}
}
