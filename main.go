package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {

	result := strings.Fields(text)
	for i := 0; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}

	return result
}

func main() {
	fmt.Println("Hello, World!")
	text := "hey I JUST       met you"
	fmt.Println(cleanInput(text))
}
