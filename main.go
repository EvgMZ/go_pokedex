package main

import (
	"pokedex/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeApiClient: pokeClient,
	}
	// start()
	start(cfg)
}
