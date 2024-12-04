package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	res, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = res.Next
	cfg.prevLocationUrl = res.Previous
	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	// cfg.pokeApiClient.ListLocations(c)
	return nil

}

func commandMapf(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you are on the first time ")
	}
	res, err := cfg.pokeApiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationUrl = res.Next
	cfg.prevLocationUrl = res.Previous
	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}
