package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	location, err := cfg.pokeApiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s ...\n", location.Name)
	fmt.Println("Found Pokemon")
	for _, val := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", val.Pokemon.Name)
	}
	return nil
}
