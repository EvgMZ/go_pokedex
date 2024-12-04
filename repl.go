package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/pokeapi"
	"strings"
)

type config struct {
	pokeApiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemon   map[string]pokeapi.Pokemon
}
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := inputWords(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Command does not exists")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func inputWords(input string) []string {
	tolower := strings.ToLower(input)
	words := strings.Fields(tolower)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Показывает список доступных команд",
			callback:    commandHelp, // Передаем список команд
		},
		"pokedex": {
			name:        "pokedex",
			description: "Показывает список доступных команд",
			callback:    commandPokedex, // Передаем список команд
		},
		"inspect": {
			name:        "inspect",
			description: "Показывает список доступных команд",
			callback:    commandInspect, // Передаем список команд
		},
		"catch": {
			name:        "catch",
			description: "Показывает список доступных команд",
			callback:    commandCatch, // Передаем список команд
		},
		"explore": {
			name:        "explore",
			description: "Показывает список доступных команд",
			callback:    commandExplore, // Передаем список команд
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the prev page of locations",
			callback:    commandMapf,
		},
		"exit": {
			name:        "exit",
			description: "Выход из программы",
			callback:    commandExit,
		},
	}
}
