package main

import (
	"encoding/json"
	"fmt"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func fetchAndPrint(locationAreaUrl string, c *config) error {

	body, err := fetchFromUrl(locationAreaUrl, c)
	if err != nil {
		return nil
	}

	var la LocationArea
	if err := json.Unmarshal(body, &la); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	c.Next = la.Next
	c.Previous = la.Previous

	for _, v := range la.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMap(c *config, name string) error {

	locationAreaUrl := "https://pokeapi.co/api/v2/location-area"

	if len(c.Next) > 0 {
		locationAreaUrl = c.Next
	}

	return fetchAndPrint(locationAreaUrl, c)

}

func commandMapb(c *config, name string) error {
	locationAreaUrl := "https://pokeapi.co/api/v2/location-area"

	if len(c.Previous) > 0 {
		locationAreaUrl = c.Previous
	}

	return fetchAndPrint(locationAreaUrl, c)

}
