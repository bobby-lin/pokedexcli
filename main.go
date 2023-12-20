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
	callback    func() error
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range getCommands() {
		fmt.Println(k, ": ", v.description)
	}

	return nil
}

func commandExit() error {
	return errors.New("exit")
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
			callback:    location.CommandLocation,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 new locations",
			callback:    location.CommandPreviousLocation,
		},
	}
}

func main() {
	fmt.Print("Pokedex > ")
	s := bufio.NewScanner(os.Stdin)
	cmdMap := getCommands()

	for s.Scan() {
		commandName := s.Text()
		c := cmdMap[commandName]
		err := c.callback()

		if err != nil {
			break
		}

		fmt.Println()
		fmt.Print("Pokedex > ")
	}
}
