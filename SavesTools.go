package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/nsf/termbox-go"
)

func ListFilesInFolder(path string) []string {
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

func DeleteSaveIfWinOrLoose() {
	if (GameData.WordFinded || GameData.Attempts == 0) && GameData.CurrentSavesPath != "" {
		os.Remove(GameData.CurrentSavesPath)
	}
}

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
