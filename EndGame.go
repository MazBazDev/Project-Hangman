package hangman

import "github.com/nsf/termbox-go"

func EndGame() {
	termbox.Close()
	DeleteSaveIfWinOrLoose()
}
