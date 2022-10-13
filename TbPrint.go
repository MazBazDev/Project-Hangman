package hangman

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func TbPrint(row, col int, fg, bg termbox.Attribute, text string) {
	for _, letter := range text {
		termbox.SetCell(row, col, letter, fg, bg)
		row += runewidth.RuneWidth(letter)
	}
}
