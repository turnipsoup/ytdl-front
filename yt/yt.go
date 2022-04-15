package yt

import (
	"fmt"
	"os/exec"
	"strings"
)

// Take a ytId and return a full YouTube Video URL
func CreateYTUrl(ytId string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", ytId)
}

// Download the audio of the passed YouTube URL
func DownloadVideoAudio(dbLoc string, ytId string, rootDirectory string, genre string) {

	url := CreateYTUrl(ytId)

	// If the user passes an entire YouTube URL, accept that instead
	if strings.Contains(ytId, "v=") {
		url = ytId
	}

	directoryTemplate := "%(title)s.%(ext)s"
	genreDirectory := fmt.Sprintf("%s/%s", rootDirectory, genre)
	directoryString := fmt.Sprintf("%s/%s", genreDirectory, directoryTemplate)

	cmd := exec.Command("yt-dlp", "-x", "-o", directoryString, url)

	MarkDownloadDone(dbLoc, ytId)

	cmd.Run()
}
