package main

import (
	"fmt"
	"slices"
)

func commandHelp(args []string, cfg *config) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	// Sort the output for consistency
	usage := []string{}
	for k, v := range commandRegistry {
		usage = append(usage, fmt.Sprintf("%s: %s", k, v.description))
	}
	slices.Sort(usage)
	for _, s := range usage {
		fmt.Println(s)
	}
	fmt.Println("")
	return nil
}
