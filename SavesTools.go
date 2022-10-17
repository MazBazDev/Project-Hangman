package hangman

import (
	"log"
	"os"
)

func ListSaves(path string) []string {
	var files []string

	f, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}
