package files

import (
	"fmt"
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
		fmt.Println(genres[t].Name())
	}
}
