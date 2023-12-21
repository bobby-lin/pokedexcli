package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bobby-lin/pokedexcli/internal/api/location"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	nextUrl     string
	previousUrl string
}

func commandHelp(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range getCommands() {
		fmt.Println(k, ": ", v.description)
	}

	return nil
}

func commandExit(c *config) error {
	return errors.New("exit")
}

func commandLocation(c *config) error {
	locations := location.GetLocation()

	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandPreviousLocation(c *config) error {
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

	for s.Scan() {
		commandName := s.Text()
		c := cmdMap[commandName]
		err := c.callback(&cf)

		if err != nil {
			break
		}

		fmt.Println()
		fmt.Print("Pokedex > ")
	}
}
