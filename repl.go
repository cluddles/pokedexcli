package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokedexcli/pokeapi"
	"strings"
)

type config struct {
	client               pokeapi.Client
	locationAreasNextUrl *string
	locationAreasPrevUrl *string
	caught               map[string]pokeapi.Pokemon
}

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
		"catch": {
			name:        "catch",
			description: "Try to catch the named Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View stats of a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View list of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func replLoop() {
	cfg := &config{
		client: pokeapi.NewClient(),
		caught: map[string]pokeapi.Pokemon{},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		runCommand(text, cfg)
	}
}

func runCommand(text string, cfg *config) {
	fields := cleanInput(text)

	if len(fields) > 0 {
		command := fields[0]
		def, exists := commandRegistry[command]
		if exists {
			err := def.callback(fields[1:], cfg)
			if err != nil {
				fmt.Printf("Command error: %v\n", err)
			}
		} else {
			fmt.Printf("Unrecognised command: %s\n", command)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
