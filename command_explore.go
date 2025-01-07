package main

import (
	"fmt"
	"pokedexcli/pokeapi"
)

func commandExplore(args []string, cfg *config) error {
	if len(args) < 1 {
		return fmt.Errorf("expected location name")
	}
	loc := args[0]
	res, err := pokeapi.GetLocationArea(cfg.client, loc)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", loc)
	fmt.Println("Found Pokemon:")
	for _, p := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}
