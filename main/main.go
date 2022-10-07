package main

import (
	"fmt"
	"hangman"
)

func main() {
	fmt.Println(hangman.GetRandomWord("./words/words.txt"))
}
