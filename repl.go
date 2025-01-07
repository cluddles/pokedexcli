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
	cfg := &config{
		client: pokeapi.NewClient(),
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

type config struct {
	client               pokeapi.Client
	locationAreasNextUrl *string
	locationAreasPrevUrl *string
}
