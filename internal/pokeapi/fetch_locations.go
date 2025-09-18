package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationResponse{}, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	var locations LocationResponse
	if err = json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return LocationResponse{}, err
	}

	return locations, nil
}
