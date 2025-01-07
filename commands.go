package main

import (
	"fmt"
	"os"
	"pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string, *config) error
}

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays prev 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List pokemon in the named location",
			callback:    commandExplore,
		},
	}
}

func commandExit(args []string, cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args []string, cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMapCommon(url *string, cfg *config) error {
	res, err := pokeapi.GetLocationAreas(cfg.client, url)
	if err != nil {
		return err
	}
	for _, loc := range res.Locations {
		fmt.Println(loc.Name)
	}
	cfg.locationAreasNextUrl = res.Next
	cfg.locationAreasPrevUrl = res.Previous
	return nil
}

func commandMap(args []string, cfg *config) error {
	return commandMapCommon(cfg.locationAreasNextUrl, cfg)
}

func commandMapb(args []string, cfg *config) error {
	if cfg.locationAreasPrevUrl == nil {
		return fmt.Errorf("you're on the first page")
	}
	return commandMapCommon(cfg.locationAreasPrevUrl, cfg)
}

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
