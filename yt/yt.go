package yt

import (
	"fmt"
	"os/exec"
)

// Take a ytId and return a full YouTube Video URL
func CreateYTUrl(ytId string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", ytId)
}

// Download the audio of the passed YouTube URL
func DownloadVideoAudio(ytId string, rootDirectory string) {
	url := CreateYTUrl(ytId)

	directoryTemplate := "%(title)s.%(ext)s"
	directoryString := fmt.Sprintf("%s/%s", rootDirectory, directoryTemplate)

	cmd := exec.Command("yt-dlp", "-x", "-o", directoryString, url)

	cmd.Run()
}
