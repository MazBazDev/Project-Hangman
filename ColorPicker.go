package hangman

import "github.com/nsf/termbox-go"

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
