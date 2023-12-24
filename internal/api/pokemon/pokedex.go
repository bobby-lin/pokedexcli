package pokemon

import (
	"errors"
	"fmt"
)

func InspectPokedex(pokedex map[string]Pokemon, name string) (Pokemon, error) {
	p, ok := pokedex[name]

	if !ok {
		return Pokemon{}, errors.New("you have not caught that pokemon")
	}

	return p, nil
}

func DisplayPokemonInfo(p Pokemon) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)

	fmt.Println("Stats:")
	for _, v := range p.Stats {
		fmt.Println("  -"+v.Stat.Name+":", v.BaseStat)
	}

	fmt.Println("Types:")
	for _, v := range p.Types {
		fmt.Println("  -" + v.Type.Name)
	}
}

func DisplayAllCaughtPokemon(pokedex map[string]Pokemon) {
	fmt.Println("Your Pokedex:")
	for k, _ := range pokedex {
		fmt.Println(" -", k)
	}
}
