package pokeapi

// Data on multiple locations.
// Result of GetLocationAreas call
type LocationAreas struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Get paginated list of locations
func GetLocationAreas(client Client, url *string) (*LocationAreas, error) {
	var result *LocationAreas
	if err := client.GetAndUnmarshal(ApiUrlOrOverride("/location-area/", url), &result); err != nil {
		return nil, err
	}
	return result, nil
}
