package yt

import (
	"fmt"
	"os/exec"
	"strings"
)

// Take a ytId and return a full YouTube Video URL
func CreateYTUrl(ytId string) string {

	if strings.Contains(ytId, "youtu") {
		return ytId
	}

	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", ytId)
}

// Parse YouTube URL and return just the ID.
// If passed just an ID, just return that same ID.
func ParseYouTubeURL(inp string) string {
	if strings.Contains(inp, "/") {

		if strings.Contains(inp, "v=") {
			splitStr := strings.Split(inp, "v=")
			return splitStr[len(splitStr)-1]
		} else {
			splitStr := strings.Split(inp, "/")
			return splitStr[len(splitStr)-1]
		}

	}

	// If nothing else matches, it must be a plain ID
	return inp
}

// Download the audio of the passed YouTube URL
func DownloadVideoAudio(dbLoc string, ytId string, rootDirectory string, genre string) {

	url := CreateYTUrl(ytId)

	// If the user passes an entire YouTube URL, accept that instead
	if strings.Contains(ytId, "youtu") {
		url = ytId
	}

	directoryTemplate := "%(title)s.%(ext)s"
	genreDirectory := fmt.Sprintf("%s/%s", rootDirectory, genre)
	directoryString := fmt.Sprintf("%s/%s", genreDirectory, directoryTemplate)

	cmd := exec.Command("yt-dlp", "-x", "-o", directoryString, url)

	cmd.Run()

	MarkDownloadDone(dbLoc, ytId)
}
