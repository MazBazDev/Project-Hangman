package main

import (
	"hangman"

	"github.com/nsf/termbox-go"
)

func main() {
	hangman.GameData.PaternsPath = "./files/hangman.txt"
	hangman.GameData.SavesPath = "./files/saves/"
	hangman.GameData.DictionaryPath = "./files/dictionary/"
	hangman.GameData.Attempts = 10

	Selector("saves")
}

func Selector(what string) {
	var Files []string
	var Title string
	var Heigth int

	if what == "dictionary" {
		Files = hangman.ListFilesInFolder(hangman.GameData.DictionaryPath)
		Title = "Select your dictionary"
		Heigth = len(hangman.ListFilesInFolder(hangman.GameData.DictionaryPath))
	} else if what == "saves" {
		Files = hangman.ListFilesInFolder(hangman.GameData.SavesPath)
		Title = "Press \"ESC\" to start a new game"
		Heigth = len(hangman.ListFilesInFolder(hangman.GameData.SavesPath))

	}

	if len(Files) == 0 {
		Selector("dictionary")
	} else {
		var Selectindex int
		err := termbox.Init()
		if err != nil {
			panic(err)
		}

		defer NextSelector(what)

	mainloop:
		for {
			hangman.CreateBox(Heigth+2, 94, 0, 0, "white", "black", Title, "cyan", Files, "white", 4)
			hangman.TbPrint(2, Selectindex+1, "white", "black", ">>")

			if what == "saves" {
				hangman.TbPrint(2, 10, "white", "black", "")
			}

			termbox.Flush()

			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					break mainloop
				case termbox.KeyArrowDown:
					if Selectindex < len(Files)-1 {
						Selectindex++
					}
				case termbox.KeyArrowUp:
					if Selectindex > 0 {
						Selectindex--
					}
				case termbox.KeyEnter:
					if what == "dictionary" {
						hangman.GameData.CurrentDictionaryPath = hangman.GetPathFromIndex(hangman.GameData.DictionaryPath, Selectindex)

					} else if what == "saves" {
						hangman.GameData.CurrentSavesPath = hangman.GetPathFromIndex(hangman.GameData.SavesPath, Selectindex)
					}
					break mainloop
				}
			}
		}
	}
}

func NextSelector(what string) {
	if what == "saves" {
		if hangman.GameData.CurrentSavesPath == "" {
			termbox.Close()
			Selector("dictionary")
		} else {
			termbox.Close()
			hangman.LoadSave(hangman.GameData.CurrentSavesPath)
			hangman.GameMain()
		}
	} else if what == "dictionary" {
		if hangman.GameData.CurrentDictionaryPath == "" {
			termbox.Close()
		} else {
			termbox.Close()
			hangman.GameData.WordToFind = hangman.GetRandomWord(hangman.GameData.CurrentDictionaryPath)
			hangman.WordBegining(hangman.GameData.WordToFind)
			hangman.GameMain()

		}
	}
}
