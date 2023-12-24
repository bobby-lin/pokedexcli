package pokemon

import (
	"fmt"
	"math"
	"math/rand"
)

func CatchPokemon(name string, baseExperience int) bool {
	fmt.Println("Throwing a Pokeball at " + name + "...")

	n := rand.Float64()
	catchThreshold := getCatchThreshold(float64(baseExperience))

	if n < catchThreshold {
		fmt.Println(name, "was caught!")
		return true
	}

	fmt.Println(name, "escaped!")
	return false
}

func getCatchThreshold(difficulty float64) float64 {
	if difficulty <= 200 {
		return 1.0 - (difficulty / 1000)
	} else if difficulty < 300 {
		return 1.0 - (difficulty/1000)*2
	} else {
		return 1.0 - math.Min(0.999999, difficulty/1000*3.333)
	}
}
