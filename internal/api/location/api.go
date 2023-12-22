package location

import (
	"encoding/json"
	"fmt"
	"github.com/bobby-lin/pokedexcli/internal/cache"
	"io"
	"log"
	"net/http"
)

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationData(url string, cache *cache.Cache) location {
	l := location{}

	cacheBody, isFound := cache.Get(url)

	if isFound {
		fmt.Println("Getting data from cache: " + url)
		err := json.Unmarshal(cacheBody, &l)

		if err != nil {
			log.Fatal(err)
		}

		return l
	}

	fmt.Println("Fetching data from url: " + url)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &l)

	cache.Add(url, body)

	return l
}
