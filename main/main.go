package main

import (
	"hangman"
	"strings"

	"github.com/nsf/termbox-go"
)

func main() {
	hangman.GameData.PaternsPath = "./files/hangman.txt"
	hangman.GameData.SavesPath = "./files/saves/"
	hangman.GameData.DictionaryPath = "./files/dictionary/"
	hangman.GameData.AsciiPath = "./files/ascii/"
	hangman.GameData.Attempts = 10

	Selector("saves")
}

func Selector(what string) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	var Title string
	var Heigth int
	var Rows []string

	if what == "dictionary" {
		Files := hangman.ListFilesInFolder(hangman.GameData.DictionaryPath)
		for _, v := range Files {
			Rows = append(Rows, strings.Replace(v, ".txt", "", -1))
		}
		Title = "Select your dictionary"
		Heigth = len(Files)

	} else if what == "saves" {
		Rows = []string{""}

		Files := hangman.ListFilesInFolder(hangman.GameData.SavesPath)
		for _, v := range Files {
			Rows = append(Rows, strings.Replace(v, ".json", "", -1))
		}
		Title = "Select a save / New game"
		Heigth = len(Files) + 1

	} else if what == "ascii" {
		Rows = []string{"Oui", "Non"}
		Title = "Use Ascii art design ?"
		Heigth = 2
	} else if what == "ascii2" {
		Files := hangman.ListFilesInFolder(hangman.GameData.AsciiPath)
		for _, v := range Files {
			Rows = append(Rows, strings.Replace(v, ".txt", "", -1))
		}
		Title = "Select your Ascii theme"
		Heigth = len(Files)

	}

	if len(Rows) == 0 {
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
			hangman.CreateBox(Heigth+2, 94, 0, 0, "white", "black", Title, "white", Rows, "white", 4)

			if what == "saves" {
				hangman.TbPrint(5, 1, "cyan", "black", "Start a new game")
			}

			hangman.TbPrint(2, Selectindex+1, "white", "black", ">>")
			termbox.Flush()

			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowDown:
					if Selectindex < len(Rows)-1 {
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
						if Selectindex == 0 {
							hangman.GameData.UseAscii = true
						} else {
							hangman.GameData.UseAscii = false
						}
					} else if what == "ascii2" {
						hangman.GameData.CurrentAsciiPath = hangman.GetPathFromIndex(hangman.GameData.AsciiPath, Selectindex)
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
		hangman.GameData.WordToFind = hangman.GetRandomWord(hangman.GameData.CurrentDictionaryPath)
		hangman.WordBegining(hangman.GameData.WordToFind)
		Selector("ascii")

	} else if what == "ascii" {
		if hangman.GameData.UseAscii {
			Selector("ascii2")
		} else {
			hangman.GameMain()
		}
	} else if what == "ascii2" {
		hangman.GameMain()
	}
}
