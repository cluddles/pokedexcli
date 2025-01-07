package pokeapi

import (
	"encoding/json"
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

func GetLocationAreas(client Client, url *string) (*LocationAreasResponse, error) {
	data, err := client.DoGet(ApiUrlOrOverride("/location-area/", url))
	if err != nil {
		return nil, err
	}

	var result *LocationAreasResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
