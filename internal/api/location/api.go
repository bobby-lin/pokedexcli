package location

import (
	"encoding/json"
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

func GetLocationData(url string) location {
	res, err := http.Get(url)

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
