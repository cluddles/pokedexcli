package main

import (
	"fmt"
)

func commandPokedex(args []string, cfg *config) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.caught {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
