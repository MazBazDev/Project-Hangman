package hangman

func PageWelcome() {
	body := []string{
		"",
		"Welcome to HangMan Termbox version.",
		"Good luck to you !",
		"/!\\ PLEASE NOTE: each letter entered is final!",
		"Press \"Enter\" to confirm your choice.",
		"By ANNEG Noemie & YAKOUBEN Mazigh"}

	CreateBox(9, 94, 4, 0, "white", "black", "Welcome", "white", body, "white", 5)
}
