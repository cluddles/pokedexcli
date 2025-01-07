package main

import (
	"fmt"
)

func commandInspect(args []string, cfg *config) error {
	if len(args) < 1 {
		return fmt.Errorf("expected pokemon name")
	}
	name := args[0]
	pokemon, exists := cfg.caught[name]
	if !exists {
		fmt.Printf("You haven't caught %s\n", name)
		return nil
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
