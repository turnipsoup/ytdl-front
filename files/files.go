package files

import (
	"io/ioutil"
	"log"
	"strings"
)

// Get all of the genres we can save to
func GetAllGenres(rootDirectory string) []string {
	genreDirs, err := ioutil.ReadDir(rootDirectory)
	var genres []string

	if err != nil {
		log.Println(err)
	}

	for i := range genreDirs {
		if !strings.Contains(genreDirs[i].Name(), ".") {
			genres = append(genres, genreDirs[i].Name())
		}

	}

	return genres

}
