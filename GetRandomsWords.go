package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetWords(path string) []string {
	var fileLines []string
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	return fileLines
}

func GetRandomWord(path string) string {
	words := GetWords(path)
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(len(words))
	return words[r]
}

func GetRandomLettersInWord(word string) string {
	var randomLetters string

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(word)/2-1; i++ {

		rand := rand.Intn(len(word))

		randomLetters += string(word[rand])
	}
	return randomLetters
}
