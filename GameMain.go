package hangman

import (
	"strings"

	"github.com/nsf/termbox-go"
)

func GameMain() {
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
				if AskSaveGame() {
					CreateSave()
					break mainloop
				} else {
					break mainloop
				}
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
				if len(GameData.CurrentLetter) >= 1 {
					Play()
				}
			case termbox.KeyBackspace2, termbox.KeyDelete:
				if len(GameData.CurrentLetter) > 0 {
					GameData.CurrentLetter = GameData.CurrentLetter[:len(GameData.CurrentLetter)-1]
					NavigateTo(GameData.CurrentPage)
				}
			default:
				if GameData.Attempts == 0 || GameData.WordFinded {
					break mainloop
				}

				for i := 'A'; i <= 'Z'; i++ {
					if i == ev.Ch {
						GameData.CurrentLetter += strings.ToLower(string(ev.Ch))
						NavigateTo(GameData.CurrentPage)
					}
				}
				for i := 'a'; i <= 'z'; i++ {
					if i == ev.Ch {
						GameData.CurrentLetter += string(ev.Ch)
						NavigateTo(GameData.CurrentPage)
					}
				}
			}
		}
	}
}
