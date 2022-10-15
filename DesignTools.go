package hangman

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func CreateBox(rows, cols, marginTop, marginLeft int, boxColor, boxBg, title, titleColor string, lines []string, lineColor string, linesMarginLeft int) {
	rows, cols = cols, rows
	for col := 0; col < cols; col++ {
		if col == 0 || col == (cols-1) {
			for row := 0; row < rows; row++ {
				if row == 0 {
					if col == 0 {
						termbox.SetCell(row+marginLeft, col+marginTop, '╔', colorPicker(boxColor), colorPicker(boxBg))
					} else if col == cols-1 {
						termbox.SetCell(row+marginLeft, col+marginTop, '╚', colorPicker(boxColor), colorPicker(boxBg))
					}
				} else if row == (rows - 1) {
					if col == 0 {
						termbox.SetCell(row+marginLeft, col+marginTop, '╗', colorPicker(boxColor), colorPicker(boxBg))
					} else if col == cols-1 {
						termbox.SetCell(row+marginLeft, col+marginTop, '╝', colorPicker(boxColor), colorPicker(boxBg))
					}
				} else {
					termbox.SetCell(row+marginLeft, col+marginTop, '═', colorPicker(boxColor), colorPicker(boxBg))
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
					termbox.SetCell(row+marginLeft, col+marginTop, '║', colorPicker(boxColor), colorPicker(boxBg))
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

func TbPrint(row, col int, FontColor, BackGroundColor string, text string) {
	for _, letter := range text {
		termbox.SetCell(row, col, letter, colorPicker(FontColor), colorPicker(BackGroundColor))
		row += runewidth.RuneWidth(letter)
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
	case "bleu":
		return termbox.ColorBlue
	case "cyan":
		return termbox.ColorCyan
	case "green":
		return termbox.ColorGreen
	default:
		return termbox.ColorDefault
	}
}
