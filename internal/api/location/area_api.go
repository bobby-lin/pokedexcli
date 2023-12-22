package location

import (
	"encoding/json"
	"github.com/bobby-lin/pokedexcli/internal/cache"
	"io"
	"net/http"
)

type Area struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationArea(cache *cache.Cache, url string) (Area, error) {
	a := Area{}
	cacheBody, isFound := cache.Get(url)
	if isFound {
		err := json.Unmarshal(cacheBody, &a)
		if err != nil {
			return a, err
		}
		return a, nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &a)

	if err != nil {
		return a, err
	}

	cache.Add(url, body)

	return a, nil

}
