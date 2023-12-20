package location

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation() location {
	fmt.Println("Getting location...")

	res, err := http.Get("https://pokeapi.co/api/v2/location")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(err)
	}

	l := location{}

	err = json.Unmarshal(body, &l)
	return l
}

func CommandLocation() error {
	locations := GetLocation()

	for _, v := range locations.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func CommandPreviousLocation() error {
	return nil
}
