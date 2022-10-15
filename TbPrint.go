package hangman

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func TbPrint(row, col int, FontColor, BackGroundColor string, text string) {
	for _, letter := range text {
		termbox.SetCell(row, col, letter, colorPicker(FontColor), colorPicker(BackGroundColor))
		row += runewidth.RuneWidth(letter)
	}
}
