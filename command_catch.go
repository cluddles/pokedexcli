package main

import (
	"fmt"
	"math/rand"
	"pokedexcli/pokeapi"
)

func commandCatch(args []string, cfg *config) error {
	if len(args) < 1 {
		return fmt.Errorf("expected pokemon name")
	}
	pokemon := args[0]
	res, err := pokeapi.GetPokemon(cfg.client, pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	difficulty := res.BaseExperience
	roll := rand.Int() % 500
	if roll > difficulty {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.caught[pokemon] = *res
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}
