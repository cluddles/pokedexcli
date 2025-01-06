package main

import (
	"fmt"
	"os"
	"strings"
)

func runCommand(text string) {
	fields := cleanInput(text)

	if len(fields) > 0 {
		command := fields[0]
		def, exists := registry[command]
		if exists {
			def.callback()
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry map[string]cliCommand

func init() {
	registry = map[string]cliCommand{
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for k, v := range registry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
