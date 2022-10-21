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
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	var Files []string
	var Title string
	var Heigth int

	if what == "dictionary" {
		Files = hangman.ListFilesInFolder(hangman.GameData.DictionaryPath)
		Title = "Select your dictionary"
		Heigth = len(hangman.ListFilesInFolder(hangman.GameData.DictionaryPath))

	} else if what == "saves" {
		Files = hangman.ListFilesInFolder(hangman.GameData.SavesPath)
		Title = "Select a save / New game"
		Heigth = len(hangman.ListFilesInFolder(hangman.GameData.SavesPath))

	} else if what == "ascii" {
		Files = []string{"Oui", "Non"}
		Title = "Press \"ESC\" to start a new game"
		Heigth = 2

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
			hangman.CreateBox(Heigth+3, 94, 0, 0, "white", "black", Title, "white", Files, "white", 4)

			if what == "saves" {
				hangman.TbPrint(Heigth+2, 1, "cyan", "black", "Create a new game")
			}

			hangman.TbPrint(2, Selectindex+1, "white", "black", ">>")
			termbox.Flush()

			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
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
						if Selectindex != 0 {
							hangman.GameData.CurrentSavesPath = hangman.GetPathFromIndex(hangman.GameData.SavesPath, Selectindex-1)
						}
					} else if what == "ascii" {
						if Selectindex == 1 {
							hangman.GameData.UseAscii = true
						} else {
							hangman.GameData.UseAscii = false
						}
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
			Selector("dictionary")
		} else {
			hangman.LoadSave(hangman.GameData.CurrentSavesPath)
			hangman.GameMain()
		}
	} else if what == "dictionary" {
		if hangman.GameData.CurrentDictionaryPath == "" {
			termbox.Close()
		} else {
			hangman.GameData.WordToFind = hangman.GetRandomWord(hangman.GameData.CurrentDictionaryPath)
			hangman.WordBegining(hangman.GameData.WordToFind)
			Selector("ascii")
		}
	} else if what == "ascii" {
		hangman.GameMain()
	}
}
