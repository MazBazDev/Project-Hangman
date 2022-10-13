package main

import (
	"fmt"
	"hangman"

	"github.com/nsf/termbox-go"
)

var currentPage int = 0
var Attempts int = 0
var Word string = "Bonj___"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer hangman.EndGame()
	NavigateTo(0)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowRight:
				switch currentPage {
				case 0:
					NavigateTo(1)
				case 1:
					NavigateTo(2)
				case 2:
					NavigateTo(0)
				}
			case termbox.KeyArrowLeft:
				switch currentPage {
				case 0:
					NavigateTo(2)
				case 1:
					NavigateTo(0)
				case 2:
					NavigateTo(1)
				}
			default:
				for i := 'A'; i <= 'Z'; i++ {
					if i == ev.Ch {
						fmt.Println(ev.Ch)
					}
				}

			}
		}
	}
}
func Refresh() {
	NavigateTo(currentPage)
}

func NavBar() {
	var selectedIndex int

	switch currentPage {
	case 0:
		selectedIndex = 6
	case 1:
		selectedIndex = 28
	case 2:
		selectedIndex = 51
	}

	body := []string{"Welcome               Game                   Help"}
	hangman.CreateBox(3, 70, 0, 0, "white", "black", "Welcome", "white", body, "white", 7)
	hangman.TbPrint(selectedIndex, 1, termbox.ColorWhite, termbox.ColorDefault, ">>")
}

func NavigateTo(page int) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	switch page {
	case 0:
		currentPage = page
		hangman.PageWelcome()
	case 1:
		currentPage = page
		hangman.PageGame(Attempts, Word)
	case 2:
		currentPage = page
		hangman.PageHelp()
	}
	NavBar()
	termbox.Flush()
}
