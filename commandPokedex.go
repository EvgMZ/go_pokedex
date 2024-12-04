package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range cfg.caughtPokemon {
		fmt.Println("  -", val.Name)
	}
	return nil
}
