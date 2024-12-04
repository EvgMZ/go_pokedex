package main

import (
	"errors"
	"fmt"

	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	res, err := cfg.pokeApiClient.GetPokemonInfo(args[0])
	if err != nil {
		return err
	}
	random := rand.Intn(res.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)
	if random > 50 {
		fmt.Printf("%s escaped!\n", res.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", res.Name)
	cfg.caughtPokemon[res.Name] = res
	return nil
}
