package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, name string) error {

	getPokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + name

	body, err := fetchFromUrl(getPokemonUrl, c)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	var oddDecimal float32
	if pokemon.BaseExperience > 0 {
		oddDecimal = 1.0 / float32(pokemon.BaseExperience)
		if oddDecimal < 0.2 {
			oddDecimal = 0.2
		}
	} else {
		oddDecimal = 1
	}

	var result string
	if rand.Float32() < oddDecimal {
		result = "was caught"
	} else {
		result = "escaped"
	}

	fmt.Printf("%s %s!\n", pokemon.Name, result)

	return nil

}
