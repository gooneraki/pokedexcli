package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func fetchFromUrl(locationAreaUrl string, c *config) error {
	res, err := http.Get(locationAreaUrl)
	if err != nil {
		return fmt.Errorf("error fetching %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
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

func commandMap(c *config) error {

	locationAreaUrl := "https://pokeapi.co/api/v2/location-area"

	if len(c.Next) > 0 {
		locationAreaUrl = c.Next
	}

	return fetchFromUrl(locationAreaUrl, c)

}

func commandMapb(c *config) error {
	locationAreaUrl := "https://pokeapi.co/api/v2/location-area"

	if len(c.Previous) > 0 {
		locationAreaUrl = c.Previous
	}

	return fetchFromUrl(locationAreaUrl, c)

}
