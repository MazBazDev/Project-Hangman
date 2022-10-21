package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
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

func GetRandomLettersInWord(toFind string) string {
	var randomLetter string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(toFind); i++ {
		rand := rand.Intn(len(toFind))
		randomLetter = string(toFind[rand])
	}
	return randomLetter
}

// --
// Parameters | Type
// path | string
// --
// This function returns an array containing all the filenames of the directory passed as parameter
// --
func ListFilesInFolder(path string) []string {
	var files []string

	f, err := os.Open(path)

	if err != nil {
		log.Fatalln(err, "path: ", path)
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

// --
// This function creates/updates a json file with current game settings
// --
func CreateSave() {
	count := len(ListFilesInFolder(GameData.SavesPath))

	var Path string
	if GameData.CurrentSavesPath != "" {
		Path = GameData.CurrentSavesPath
	} else {
		Path = GameData.SavesPath + "Save-" + strconv.Itoa(count) + ".json"
		GameData.CurrentSavesPath = Path
	}

	str, err := json.Marshal(GameData)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(Path, str, os.ModePerm)
}

// --
// Parameters | Type
// path | string
// --
// This function loads a save (json file), and sets the game parameters from the contents of the file
// --
func LoadSave(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &GameData)
	if err != nil {
		log.Fatal(err)
	}
}

// --
// Parameters | Type
// path | string
// index | int
// --
// This function returns a file path deduced from an index found by the ListFilesInFolder() function
// --
func GetPathFromIndex(path string, index int) string {
	var temp string
	for i, v := range ListFilesInFolder(path) {
		if i == index {
			temp = v
			break
		}
	}
	return path + temp
}

// --
// This function destroys the current save file if the game is over
// --
func DeleteSaveIfWinOrLoose() {
	if (GameData.WordFinded || GameData.Attempts == 0) && GameData.CurrentSavesPath != "" {
		os.Remove(GameData.CurrentSavesPath)
	}
}

// --
// This function opens a modal to ask if the player wants to save his game and returns true if yes is selected
// --
func AskSaveGame() bool {
	if GameData.Attempts == 0 {
		return false
	}
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	var SelectIndex int
	if !GameData.WordFinded {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}

	mainloop:
		for {
			CreateBox(4, 94, 0, 0, "white", "black", "Do you want to save the game?", "white", []string{"Yes", "No"}, "white", 4)
			TbPrint(2, SelectIndex+1, "white", "black", ">>")
			termbox.Flush()

			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowDown:
					if SelectIndex < 1 {
						SelectIndex++
					}
				case termbox.KeyArrowUp:
					if SelectIndex > 0 {
						SelectIndex--
					}
				case termbox.KeyEnter:
					break mainloop
				}
			}
		}
	}

	if SelectIndex == 0 {
		return true
	} else {
		return false
	}
}
