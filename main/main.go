package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var currentPage int = 0

func main() {
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
				RefreshPage()
			}
		}
	}
}

func RefreshPage() {
	NavigateTo(currentPage)
}

func NavigateTo(page int) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	switch page {
	case 0:
		currentPage = page
		PageWelcome()
	case 1:
		currentPage = page
		PageGame()
	case 2:
		currentPage = page
		PageHelp()
	}
	NavBar()
	termbox.Flush()
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
	CreateBox(3, 70, 0, 0, "white", "black", "Welcome", "white", body, "white", 7)
	tbprint(selectedIndex, 1, termbox.ColorWhite, termbox.ColorDefault, ">>")
}

func PageWelcome() {
	body := []string{
		"",
		"Bienvenue dans HangMan version Termbox.",
		"Bon courage à vous !",
		"/!\\ ATTENTION : chaque lettre entrée est définitive !",
		"Appuyez sur \"Entrer\" pour confirmer votre choix.",
		"By ANNEG Noémie & YAKOUBEN Mazigh"}

	CreateBox(9, 70, 5, 0, "white", "black", "Welcome", "white", body, "white", 5)
}
func PageGame() {
	tbprint(2, 10, termbox.ColorRed, termbox.ColorDefault, "Page 1")
}
func PageHelp() {
	body := []string{
		"",
		"1. \"ESC\" to quit.",
		"2. \"ENTER\" to confirm your choice.",
		"3. \"BACKSPACE\" or \"DEL\" to delete the last letter.",
	}

	CreateBox(7, 70, 5, 0, "white", "black", "Help", "white", body, "white", 5)
}

func CreateBox(rows, cols, marginTop, marginLeft int, boxColor, boxBg, title, titleColor string, lines []string, lineColor string, linesMarginLeft int) {
	rows, cols = cols, rows
	for col := 0; col < cols; col++ {
		if col == 0 || col == (cols-1) {
			for row := 0; row < rows; row++ {
				if row == 0 || row == (rows-1) {
					termbox.SetCell(row+marginLeft, col+marginTop, '+', colorPicker(boxColor), colorPicker(boxBg))
				} else {
					termbox.SetCell(row+marginLeft, col+marginTop, '=', colorPicker(boxColor), colorPicker(boxBg))
				}
			}
			if len(title) > 0 {
				sentence := "[ " + title + " ]"
				for i := 0; i < len(sentence); i++ {
					termbox.SetCell(marginLeft+3+i, marginTop, rune(sentence[i]), colorPicker(titleColor), colorPicker(boxBg))
				}
			}
		} else {
			for row := 0; row < rows; row++ {
				if row == 0 || row == (rows-1) {
					termbox.SetCell(row+marginLeft, col+marginTop, '|', colorPicker(boxColor), colorPicker(boxBg))
				} else {
					termbox.SetCell(row+marginLeft, col+marginTop, ' ', colorPicker(boxColor), colorPicker(boxBg))
				}
			}
		}
	}

	for i, line := range lines {
		for e, letter := range line {
			termbox.SetCell(linesMarginLeft+marginLeft+1+e, marginTop+1+i, letter, colorPicker(lineColor), colorPicker(boxBg))
		}
	}
}

func colorPicker(color string) termbox.Attribute {
	switch color {
	case "black":
		return termbox.ColorBlack
	case "white":
		return termbox.ColorWhite
	case "red":
		return termbox.ColorRed
	default:
		return termbox.ColorBlack
	}
}

func tbprint(row, col int, fg, bg termbox.Attribute, text string) {
	for _, letter := range text {
		termbox.SetCell(row, col, letter, fg, bg)
		row += runewidth.RuneWidth(letter)
	}
}

func EndGame() {
	termbox.Close()
}
