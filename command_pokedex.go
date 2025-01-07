package main

import (
	"fmt"
)

func commandPokedex(args []string, cfg *config) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.caught {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
