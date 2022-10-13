package hangman

func PageHelp() {
	body := []string{
		"",
		"1. \"ESC\" to quit.",
		"2. \"ENTER\" to confirm your choice.",
		"3. \"BACKSPACE\" or \"DEL\" to delete the last letter.",
	}
	CreateBox(7, 70, 5, 0, "white", "black", "Help", "white", body, "white", 5)
}
