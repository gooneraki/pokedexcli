package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	userInput := ""
	scanner := bufio.NewScanner(os.Stdin)
	firstWord := ""

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput = scanner.Text()
		}
		cleaned := cleanInput(userInput)
		userInput = ""

		if len(cleaned) > 0 {
			firstWord = cleaned[0]
		} else {
			firstWord = ""
		}

		fmt.Printf("Your command was: %s\n", firstWord)

	}
}
