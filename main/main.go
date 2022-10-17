package main

import (
	"fmt"
	"hangman"
)

func main() {
	fmt.Println(hangman.OSReadDir("./files/saves"))
}
