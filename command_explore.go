package main

import (
	"encoding/json"
	"fmt"
)

type ExploredArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, name string) error {

	locationAreaUrl := "https://pokeapi.co/api/v2/location-area/" + name

	fmt.Printf("Exploring %s...\n", name)

	body, err := fetchFromUrl(locationAreaUrl, c)
	if err != nil {
		return err
	}

	var exArea ExploredArea
	if err := json.Unmarshal(body, &exArea); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if len(exArea.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")

		for _, v := range exArea.PokemonEncounters {
			fmt.Printf("- %s\n", v.Pokemon.Name)
		}
	}

	return nil

}
