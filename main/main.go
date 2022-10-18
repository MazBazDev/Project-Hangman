package main

import (
	"hangman"
	"strings"

	"github.com/nsf/termbox-go"
)

func InitGame() {
	hangman.GameData.PaternsPath = "./files/hangman.txt"
	hangman.GameData.DictionaryPath = "./files/words.txt"
	hangman.GameData.Attempts = 10
	hangman.GameData.WordToFind = hangman.GetRandomWord(hangman.GameData.DictionaryPath)
}
func main() {
	InitGame()

	WordBegining(hangman.GameData.WordToFind)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	NavigateTo(0)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
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
			case termbox.KeyBackspace2:
				hangman.GameData.CurrentLetter = ""
				Refresh()
			case termbox.KeyDelete:
				hangman.GameData.CurrentLetter = ""
				Refresh()
			default:
				if hangman.GameData.Attempts == 0 || hangman.GameData.WordFinded {
					break mainloop
				}

				for i := 'A'; i <= 'Z'; i++ {
					if i == ev.Ch {
						hangman.GameData.CurrentLetter += strings.ToLower(string(ev.Ch))
						Refresh()
					}
				}
				for i := 'a'; i <= 'z'; i++ {
					if i == ev.Ch {
						hangman.GameData.CurrentLetter += string(ev.Ch)
						Refresh()
					}
				}
			}
		}
	}
}

func Refresh() {
	NavigateTo(hangman.GameData.CurrentPage)
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
			hangman.GameData.Attempts--
			hangman.GameData.CurrentLetter = ""
		}
	}

	Refresh()
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
