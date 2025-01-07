package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationAreasResponse struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const baseUrl = "https://pokeapi.co/api/v2"

func GetLocationAreas(url *string) (LocationAreasResponse, error) {
	var requestUrl string
	if url != nil {
		requestUrl = *url
	} else {
		requestUrl = baseUrl + "/location-area/"
	}

	res, err := http.Get(requestUrl)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var result LocationAreasResponse
	if err := decoder.Decode(&result); err != nil {
		return LocationAreasResponse{}, err
	}

	return result, nil
}
