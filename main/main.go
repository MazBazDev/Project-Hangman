package main

import (
	"hangman"
	"strings"

	"github.com/nsf/termbox-go"
)

type HangmanData struct {
	CurrentPage   int
	Attempts      int
	WordToFind    string
	Word          string
	CurrentLetter string
	PlayedLetters string
	WordFinded    bool
}

var hangmanPaterns string = "./files/hangman.txt"
var hangmanWords string = "./files/words.txt"
var hangmanMaxAttempts int = 10

var GameData HangmanData

func main() {
	GameData.WordToFind = hangman.GetRandomWord(hangmanWords)
	GameData.Word = WordBegining(GameData.WordToFind)

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
				switch GameData.CurrentPage {
				case 0:
					NavigateTo(1)
				case 1:
					NavigateTo(2)
				case 2:
					NavigateTo(0)
				}
			case termbox.KeyArrowLeft:
				switch GameData.CurrentPage {
				case 0:
					NavigateTo(2)
				case 1:
					NavigateTo(0)
				case 2:
					NavigateTo(1)
				}
			case termbox.KeyEnter:
				if len(GameData.CurrentLetter) == 1 {
					Play()
				}
			case termbox.KeyBackspace2:
				if len(GameData.CurrentLetter) == 1 {
					GameData.CurrentLetter = ""
					Refresh()
				}
			case termbox.KeyDelete:
				if len(GameData.CurrentLetter) == 1 {
					GameData.CurrentLetter = ""
					Refresh()
				}
			default:
				if GameData.Attempts == hangmanMaxAttempts || GameData.WordFinded {
					break mainloop
				}

				for i := 'A'; i <= 'Z'; i++ {
					if i == ev.Ch {
						GameData.CurrentLetter = strings.ToLower(string(ev.Ch))
						Refresh()
					}
				}
				for i := 'a'; i <= 'z'; i++ {
					if i == ev.Ch {
						GameData.CurrentLetter = string(ev.Ch)
						Refresh()
					}
				}
			}
		}
	}
}

func Refresh() {
	NavigateTo(GameData.CurrentPage)
}

func NavBar() {
	var selectedIndex int
	switch GameData.CurrentPage {
	case 0:
		selectedIndex = 20
	case 1:
		selectedIndex = 42
	case 2:
		selectedIndex = 65
	}

	body := []string{"Welcome               Game                   Help"}
	hangman.CreateBox(3, 94, 0, 0, "white", "black", "Welcome", "white", body, "white", 21)
	if !(GameData.CurrentPage > 2) {
		hangman.TbPrint(selectedIndex, 1, "white", "black", ">>")
	}
}

func NavigateTo(page int) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	switch page {
	case 0:
		GameData.CurrentPage = page
		hangman.PageWelcome()
	case 1:
		GameData.CurrentPage = page
		hangman.PageGame(GameData.Attempts, GameData.Word, GameData.PlayedLetters, GameData.CurrentLetter, hangmanPaterns, hangmanMaxAttempts)
	case 2:
		GameData.CurrentPage = page
		hangman.PageHelp()
	case 3:
		GameData.CurrentPage = page
		hangman.PageFinal(GameData.WordFinded, GameData.Attempts, GameData.WordToFind, hangman.GetHangPatern(hangmanPaterns, GameData.Attempts))
	}
	NavBar()
	termbox.Flush()
}

func Play() {
	if !strings.Contains(GameData.PlayedLetters, GameData.CurrentLetter) {
		if !strings.Contains(GameData.WordToFind, GameData.CurrentLetter) {
			if GameData.Attempts == 9 {
				NavigateTo(3)
			}
			GameData.PlayedLetters += GameData.CurrentLetter
			GameData.Attempts++
		}

		if GameData.WordToFind == AddLetter(GameData.CurrentLetter, GameData.WordToFind, GameData.Word) {
			GameData.WordFinded = true
			NavigateTo(3)
		}

		GameData.Word = AddLetter(GameData.CurrentLetter, GameData.WordToFind, GameData.Word)
		GameData.CurrentLetter = ""
	} else {
		GameData.CurrentLetter = ""
	}
	Refresh()
}

func WordBegining(toFind string) string {
	for range toFind {
		GameData.Word += "_"
	}
	for i := 1; i <= len(toFind)/2-1; i++ {
		ranLetter := hangman.GetRandomLettersInWord(toFind)
		GameData.Word = AddLetter(ranLetter, toFind, GameData.Word)
		GameData.PlayedLetters += ranLetter
	}
	return GameData.Word
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
