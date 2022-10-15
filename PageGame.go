package hangman

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func PageGame(Attempts int, word, PlayedLetters, CurrentLetter, hangmanPaternsPath string, maxAttempts int) {
	body := []string{
		"You have " + strconv.Itoa(maxAttempts) + " attempts to find this word",
		"       Good luck and Have Fun ;)",
	}

	CreateBox(4, 70, 4, 0, "white", "black", "Info", "white", body, "white", 14)

	AttemptsBox(Attempts)
	HangBox(GetHangPatern("./files/hangman.txt", Attempts))
	DisplayWord(word)
	DisplayPlayedLetters(PlayedLetters)
	DisplayCurrentLetter(CurrentLetter)
}

func GetHangPatern(path string, step int) []string {
	step = step - 1

	var fileLines []string

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	start := step*7 + step
	end := step*7 + 6 + step

	for fileScanner.Scan() {
		if i >= start && i <= end {
			fileLines = append(fileLines, fileScanner.Text())
		}
		i++
	}

	readFile.Close()
	return fileLines
}

func HangBox(hangman []string) {
	CreateBox(9, 19, 8, 75, "white", "black", "HangMan", "white", hangman, "white", 4)
}

func AttemptsBox(attempts int) {
	CreateBox(3, 19, 4, 75, "white", "black", "Attempts", "white", []string{strconv.Itoa(attempts)}, "white", 8)
}

func DisplayWord(word string) {
	CreateBox(5, 70, 9, 0, "white", "black", "Word", "white", []string{"", word}, "white", 28)
}

func DisplayPlayedLetters(PlayedLetters string) {
	CreateBox(5, 19, 18, 75, "white", "black", "Letters", "white", []string{"", PlayedLetters}, "white", 2)
}

func DisplayCurrentLetter(CurrentLetter string) {
	CreateBox(5, 70, 15, 0, "white", "black", "Press \"ENTER\" to try you'r letter", "white", []string{"", CurrentLetter}, "white", 28)
}

func WordState(letter, toFind, word string) string {
	tabToFind := []string{}
	for _, v := range toFind {
		tabToFind = append(tabToFind, string(v))
	}

	if strings.Contains(toFind, letter) {
		for i, v := range tabToFind {
			if v == letter {
				word = strings.Join([]string{word[:i], string(letter), word[i+1:]}, "")
			}
		}
	}
	return word
}
