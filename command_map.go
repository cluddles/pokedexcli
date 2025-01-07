package main

import (
	"fmt"
	"pokedexcli/pokeapi"
)

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
