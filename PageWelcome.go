package hangman

func PageWelcome() {
	body := []string{
		"",
		"Bienvenue dans HangMan version Termbox.",
		"Bon courage à vous !",
		"/!\\ ATTENTION : chaque lettre entrée est définitive !",
		"Appuyez sur \"Entrer\" pour confirmer votre choix.",
		"By ANNEG Noemie & YAKOUBEN Mazigh"}

	CreateBox(9, 70, 5, 0, "white", "black", "Welcome", "white", body, "white", 5)
}
