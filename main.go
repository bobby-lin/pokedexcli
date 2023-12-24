package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bobby-lin/pokedexcli/internal/api/location"
	"github.com/bobby-lin/pokedexcli/internal/api/pokemon"
	"github.com/bobby-lin/pokedexcli/internal/cache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, param string) error
}

type config struct {
	nextUrl     string
	previousUrl string
}

func commandHelp(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, param string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range getCommands() {
		fmt.Println(k, ": ", v.description)
	}

	return nil
}

func commandExit(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, param string) error {
	return errors.New("exit")
}

func commandLocation(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, param string) error {
	if c.nextUrl == "" {
		c.nextUrl = "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	}

	locationData := location.GetLocationData(c.nextUrl, cache)

	c.nextUrl = locationData.Next
	c.previousUrl = locationData.Previous

	for _, v := range locationData.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandPreviousLocation(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, param string) error {
	if c.previousUrl == "" {
		fmt.Println("Error: No previous 20 locations!")
		return nil
	}

	locationData := location.GetLocationData(c.previousUrl, cache)
	c.nextUrl = locationData.Next
	c.previousUrl = locationData.Previous

	for _, v := range locationData.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandExplore(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, area string) error {
	fmt.Println("Exploring", area+"...")
	a, err := location.GetLocationArea(cache, "https://pokeapi.co/api/v2/location-area/"+area)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range a.PokemonEncounters {
		fmt.Println("-", v.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config, cache *cache.Cache, pokedex map[string]pokemon.Pokemon, name string) error {
	p, err := pokemon.GetPokemonInfo(name)

	if err != nil {
		return err
	}

	isCaught := pokemon.CatchPokemon(name, p.BaseExperience)

	if isCaught {
		fmt.Println("Storing pokemon to Pokedex")
		pokedex[name] = p
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays name of 20 new locations",
			callback:    commandLocation,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 new locations",
			callback:    commandPreviousLocation,
		},
		"explore": {
			name:        "explore",
			description: "Displays list of pokemon in the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func main() {
	fmt.Print("Pokedex > ")
	s := bufio.NewScanner(os.Stdin)
	cmdMap := getCommands()
	cf := config{}
	currentCache := cache.NewCache(5 * time.Minute)
	pokedexData := make(map[string]pokemon.Pokemon)

	for s.Scan() {
		commandInputs := s.Text()
		commandArr := strings.Split(commandInputs, " ")
		c, ok := cmdMap[commandArr[0]]

		if c.name == "exit" {
			break
		}

		param := ""
		if len(commandArr) > 1 {
			param = commandArr[1]
		}

		if !ok {
			fmt.Println("Error: Invalid command...")
		} else {
			err := c.callback(&cf, currentCache, pokedexData, param)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		fmt.Println()
		fmt.Print("Pokedex > ")
	}
}
