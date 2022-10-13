package hangman

import "github.com/nsf/termbox-go"

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
