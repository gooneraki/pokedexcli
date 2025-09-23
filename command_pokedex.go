package main

import "fmt"

func commandPokedex(c *config, name string) error {

	catched := c.caughtPokemon

	if len(catched) == 0 {
		fmt.Println("Nothing caught")
	}

	for k := range catched {
		fmt.Printf("- %s\n", k)
	}
	return nil

}
