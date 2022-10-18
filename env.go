package hangman

type HangmanData struct {
	CurrentPage    int
	Attempts       int
	WordToFind     string
	Word           string
	CurrentLetter  string
	PlayedLetters  string
	WordFinded     bool
	SavesPath      string
	DictionaryPath string
	PaternsPath    string
}

var GameData HangmanData
