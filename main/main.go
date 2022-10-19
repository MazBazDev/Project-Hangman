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

		defer NextStep(what)

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

func NextStep(what string) {
	if what == "saves" {
		if hangman.GameData.CurrentSavesPath == "" {
			termbox.Close()
			Selector("dictionary")
		} else {
			termbox.Close()
			hangman.LoadSave(hangman.GameData.CurrentSavesPath)
			StartGame()
		}
	} else if what == "dictionary" {
		if hangman.GameData.CurrentDictionaryPath == "" {
			termbox.Close()
		} else {
			termbox.Close()
			hangman.GameData.WordToFind = hangman.GetRandomWord(hangman.GameData.CurrentDictionaryPath)
			WordBegining(hangman.GameData.WordToFind)
			StartGame()

		}
	}
}
func StartGame() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer EndGame()

	NavigateTo(0)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				if hangman.AskSaveGame() {
					hangman.CreateSave()
					break mainloop
				} else {
					break mainloop
				}
			case termbox.KeyArrowRight:
				switch hangman.GameData.CurrentPage {
				case 0:
					NavigateTo(1)
				case 1:
					NavigateTo(2)
				case 2:
					NavigateTo(0)
				}
			case termbox.KeyArrowLeft:
				switch hangman.GameData.CurrentPage {
				case 0:
					NavigateTo(2)
				case 1:
					NavigateTo(0)
				case 2:
					NavigateTo(1)
				}
			case termbox.KeyEnter:
				if len(hangman.GameData.CurrentLetter) >= 1 {
					Play()
				}
			case termbox.KeyBackspace2, termbox.KeyDelete:
				hangman.GameData.CurrentLetter = ""
				NavigateTo(hangman.GameData.CurrentPage)
			default:
				if hangman.GameData.Attempts == 0 || hangman.GameData.WordFinded {
					break mainloop
				}

				for i := 'A'; i <= 'Z'; i++ {
					if i == ev.Ch {
						hangman.GameData.CurrentLetter += strings.ToLower(string(ev.Ch))
						NavigateTo(hangman.GameData.CurrentPage)
					}
				}
				for i := 'a'; i <= 'z'; i++ {
					if i == ev.Ch {
						hangman.GameData.CurrentLetter += string(ev.Ch)
						NavigateTo(hangman.GameData.CurrentPage)
					}
				}
			}
		}
	}
}

func NavBar() {
	var selectedIndex int
	switch hangman.GameData.CurrentPage {
	case 0:
		selectedIndex = 20
	case 1:
		selectedIndex = 42
	case 2:
		selectedIndex = 65
	}

	body := []string{"Welcome               Game                   Help"}
	hangman.CreateBox(3, 94, 0, 0, "white", "black", "Welcome", "white", body, "white", 21)
	if !(hangman.GameData.CurrentPage > 2) {
		hangman.TbPrint(selectedIndex-1, 1, "white", "black", ">>")
	}
}

func NavigateTo(page int) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	switch page {
	case 0:
		hangman.GameData.CurrentPage = page
		hangman.PageWelcome()
	case 1:
		hangman.GameData.CurrentPage = page
		hangman.PageGame(hangman.GameData.Attempts, hangman.GameData.Word, hangman.GameData.PlayedLetters, hangman.GameData.CurrentLetter, hangman.GameData.PaternsPath)
	case 2:
		hangman.GameData.CurrentPage = page
		hangman.PageHelp()
	case 3:
		hangman.GameData.CurrentPage = page
		hangman.PageFinal(hangman.GameData.WordFinded, hangman.GameData.Attempts, hangman.GameData.Word, hangman.GameData.WordToFind, hangman.GetHangPatern(hangman.GameData.PaternsPath, hangman.GameData.Attempts))
	}
	NavBar()
	termbox.Flush()
}

func Play() {
	if len(hangman.GameData.CurrentLetter) == 1 {
		if !strings.Contains(hangman.GameData.PlayedLetters, hangman.GameData.CurrentLetter) {
			if !strings.Contains(hangman.GameData.WordToFind, hangman.GameData.CurrentLetter) {
				if hangman.GameData.Attempts == 1 {
					NavigateTo(3)
				}
				hangman.GameData.PlayedLetters += hangman.GameData.CurrentLetter
				hangman.GameData.Attempts--
			}

			if hangman.GameData.WordToFind == AddLetter(hangman.GameData.CurrentLetter, hangman.GameData.WordToFind, hangman.GameData.Word) {
				hangman.GameData.WordFinded = true
				NavigateTo(3)
			}

			hangman.GameData.Word = AddLetter(hangman.GameData.CurrentLetter, hangman.GameData.WordToFind, hangman.GameData.Word)
			hangman.GameData.CurrentLetter = ""
		} else {
			hangman.GameData.CurrentLetter = ""
		}
	} else {
		if len(hangman.GameData.CurrentLetter) == len(hangman.GameData.WordToFind) && hangman.GameData.CurrentLetter == hangman.GameData.WordToFind {
			hangman.GameData.WordFinded = true
			NavigateTo(3)
		} else {
			if hangman.GameData.Attempts == 1 {
				hangman.GameData.Attempts--
			} else {
				hangman.GameData.Attempts -= 2
			}
			hangman.GameData.CurrentLetter = ""
			if hangman.GameData.Attempts == 0 {
				NavigateTo(3)
			}
		}
	}

	NavigateTo(hangman.GameData.CurrentPage)
}

func WordBegining(toFind string) string {
	for range toFind {
		hangman.GameData.Word += "_"
	}
	n := len(toFind)/2 - 1
	for i := 1; i <= n; i++ {
		randLetter := hangman.GetRandomLettersInWord(toFind)
		hangman.GameData.Word = AddLetter(randLetter, toFind, hangman.GameData.Word)

		if !strings.Contains(hangman.GameData.PlayedLetters, randLetter) {
			hangman.GameData.PlayedLetters += randLetter
		}
	}
	return hangman.GameData.Word
}

func AddLetter(letter string, toFind string, word string) string {
	tabToFind := []string{}
	for _, v := range toFind {
		tabToFind = append(tabToFind, string(v))
	}
	for i, v := range tabToFind {
		if v == letter {
			word = strings.Join([]string{word[:i], string(letter), word[i+1:]}, "")
		}
	}
	return word
}

func EndGame() {
	termbox.Close()
	hangman.DeleteSaveIfWinOrLoose()
}
