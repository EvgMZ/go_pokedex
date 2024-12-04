package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	val, exist := cfg.caughtPokemon[name]
	if !exist {
		return errors.New("you have not caught that pokemin")
	}
	fmt.Println("Name:", val.Name)
	fmt.Println("Height:", val.Height)
	fmt.Println("Weight:", val.Weight)
	fmt.Println("Stats:")
	for _, vals := range val.Stats {
		fmt.Printf("  -%s: %v\n", vals.Stat.Name, vals.BaseStat)
	}
	fmt.Println("Types:")
	for _, vals := range val.Types {
		fmt.Printf("  -%s\n", vals.Type.Name)
	}
	return nil
}
