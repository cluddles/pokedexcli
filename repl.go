package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokedexcli/pokeapi"
	"strings"
)

func replLoop() {
	cfg := &config{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
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
			err := def.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Unrecognised command: %s\n", command)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	locationAreasNextUrl *string
	locationAreasPrevUrl *string
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
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMapCommon(url *string, cfg *config) error {
	res, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	for _, loc := range res.Locations {
		fmt.Println(loc.Name)
	}
	cfg.locationAreasNextUrl = &res.Next
	cfg.locationAreasPrevUrl = &res.Previous
	return nil
}

func commandMap(cfg *config) error {
	return commandMapCommon(cfg.locationAreasNextUrl, cfg)
}

func commandMapb(cfg *config) error {
	if cfg.locationAreasPrevUrl == nil {
		return fmt.Errorf("you're on the first page")
	}
	return commandMapCommon(cfg.locationAreasPrevUrl, cfg)
}
