package main

import (
	"encoding/json"
	"fmt"
)

func commandInspect(c *config, name string) error {

	getPokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + name

	data, found := c.cache.Get(getPokemonUrl)

	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil

}
