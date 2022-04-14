package files

import (
	"io/ioutil"
	"log"
)

// Reading the files present in the storage area
func GetCurrentlyDownloading(rootDirectory string) {

	genres, err := ioutil.ReadDir(rootDirectory)

	if err != nil {
		log.Println("Could not read root directory")
		log.Println(err)
	}

	for t := range genres {
		t = t
	}
}

// Get all of the genres we can save to
func GetAllGenres(rootDirectory string) []string {
	genreDirs, err := ioutil.ReadDir(rootDirectory)
	var genres []string

	if err != nil {
		log.Println(err)
	}

	for i := range genreDirs {
		genres = append(genres, genreDirs[i].Name())
	}

	return genres

}
