package hangman

var hang0 = []string{
	"",
	"",
	"",
	"",
	"",
	"",
	"==========",
}

var hang1 = []string{
	"       +",
	"       |",
	"       |",
	"       |",
	"       |",
	"       |",
	"==========",
}

var hang2 = []string{
	"   +---+",
	"       |",
	"       |",
	"       |",
	"       |",
	"       |",
	"==========",
}

var hang3 = []string{
	"   +---+",
	"   |   |",
	"       |",
	"       |",
	"       |",
	"       |",
	"==========",
}
var hang4 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"       |",
	"       |",
	"       |",
	"==========",
}
var hang5 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"   |   |",
	"       |",
	"       |",
	"==========",
}

var hang6 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"  /|   |",
	"       |",
	"       |",
	"==========",
}

var hang7 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"  /|\\  |",
	"       |",
	"       |",
	"==========",
}
var hang8 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"  /|\\  |",
	"  /    |",
	"       |",
	"==========",
}
var hang9 = []string{
	"   +---+",
	"   |   |",
	"   O   |",
	"  /|\\  |",
	"  / \\  |",
	"       |",
	"==========",
}

func PageGame(Attempts int, word string) {
	CreateBox(3, 70, 4, 0, "white", "black", "Game", "white", []string{`Have fun !`}, "white", 28)

	AttemptsBox(Attempts)
	AttemptsHang(Attempts)

	DisplayWord(word)
}
func AttemptsHang(Attempts int) {
	switch Attempts {
	case 0:
		HangBox([]string{})
	case 1:
		HangBox(hang0)
	case 2:
		HangBox(hang1)
	case 3:
		HangBox(hang2)
	case 4:
		HangBox(hang3)
	case 5:
		HangBox(hang4)
	case 6:
		HangBox(hang5)
	case 7:
		HangBox(hang6)
	case 8:
		HangBox(hang7)
	case 9:
		HangBox(hang8)
	case 10:
		HangBox(hang9)
	}
}
func HangBox(hangman []string) {
	CreateBox(9, 19, 8, 74, "white", "black", "HangMan", "white", hangman, "white", 4)
}

func AttemptsBox(attempts int) {
	var attempt string
	switch attempts {
	case 0:
		attempt = "0"
	case 1:
		attempt = "1"
	case 2:
		attempt = "2"
	case 3:
		attempt = "3"
	case 4:
		attempt = "4"
	case 5:
		attempt = "5"
	case 6:
		attempt = "6"
	case 7:
		attempt = "7"
	case 8:
		attempt = "8"
	case 9:
		attempt = "9"
	}
	CreateBox(3, 19, 4, 74, "white", "black", "Attempts", "white", []string{attempt}, "white", 8)
}

func DisplayWord(word string) {
	CreateBox(5, 70, 8, 0, "white", "black", "Word", "white", []string{"", word}, "white", 28)
}
