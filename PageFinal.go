package hangman

import "strconv"

func PageFinal(Status bool, Attempts int, Word, WordToFind string, HangMan []string) {
	if Status {
		body := []string{
			" __     __                                 _ ",
			" \\ \\   / /                                | |",
			"  \\ \\_/ /__  _   _  __      _____  _ __   | |",
			"   \\   / _ \\| | | | \\ \\ /\\ / / _ \\| '_ \\  | |",
			"    | | (_) | |_| |  \\ V  V / (_) | | | | |_|",
			"    |_|\\___/ \\__,_|   \\_/\\_/ \\___/|_| |_| (_)",
			"",
			"         You find \"" + WordToFind + "\" in " + strconv.Itoa(Attempts) + " attempts",
			"",
		}

		for _, v := range HangMan {
			body = append(body, "                   "+v)
		}

		end := []string{
			"",
			"         Press any key on the keyboard",
			"                to close the game",
		}

		for _, v := range end {
			body = append(body, ""+v)
		}
		CreateBox(22, 94, 4, 0, "white", "black", "You won !", "white", body, "white", 22)

	} else {
		body := []string{
			"__     __           _           _",
			"\\ \\   / /          | |         | |  ",
			" \\ \\_/ /__  _   _  | | ___  ___| |_ ",
			"  \\   / _ \\| | | | | |/ _ \\/ __| __|",
			"   | | (_) | |_| | | | (_) \\__ \\ |_ ",
			"   |_|\\___/ \\__,_| |_|\\___/|___/\\__|",
			"",
			"",
		}

		for _, v := range HangMan {
			body = append(body, "             "+v)
		}

		end := []string{
			"",
			"      The word to find was:",
			"      \"" + WordToFind + "\"",
			"",
			"      You finded:",
			"      \"" + Word + "\"",
			"",
			"    Press any key on the keyboard",
			"         to close the game",
		}

		for _, v := range end {
			body = append(body, ""+v)
		}
		CreateBox(26, 94, 4, 0, "white", "black", "You lost !", "white", body, "white", 28)
	}
}
