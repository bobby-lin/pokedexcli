package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bobby-lin/pokedexcli/internal/api/location"
	"github.com/bobby-lin/pokedexcli/internal/cache"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, cache *cache.Cache) error
}

type config struct {
	nextUrl     string
	previousUrl string
}

func commandHelp(c *config, cache *cache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range getCommands() {
		fmt.Println(k, ": ", v.description)
	}

	return nil
}

func commandExit(c *config, cache *cache.Cache) error {
	return errors.New("exit")
}

func commandLocation(c *config, cache *cache.Cache) error {
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

func commandPreviousLocation(c *config, cache *cache.Cache) error {
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
	}
}

func main() {
	fmt.Print("Pokedex > ")
	s := bufio.NewScanner(os.Stdin)
	cmdMap := getCommands()
	cf := config{}
	currentCache := cache.NewCache(5 * time.Minute)

	for s.Scan() {
		commandName := s.Text()
		c, ok := cmdMap[commandName]

		if !ok {
			fmt.Println("Error: Invalid command...")
		} else {
			err := c.callback(&cf, currentCache)
			if err != nil {
				break
			}
		}

		fmt.Println()
		fmt.Print("Pokedex > ")
	}
}
