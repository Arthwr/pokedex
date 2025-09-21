package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.pokeCache.Get(url); ok {
		locations := LocationResponse{}
		if err := json.Unmarshal(val, &locations); err != nil {
			return LocationResponse{}, err
		}
		return locations, nil
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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var locations LocationResponse
	if err = json.Unmarshal(data, &locations); err != nil {
		return LocationResponse{}, err
	}

	c.pokeCache.Add(url, data)
	return locations, nil
}

func (c *Client) FetchEncountersFromLocation(locationID string) (EncountersResponse, error) {
	url := baseURL + "/location-area/" + locationID

	if val, ok := c.pokeCache.Get(url); ok {
		encounters := EncountersResponse{}
		if err := json.Unmarshal(val, &encounters); err != nil {
			return EncountersResponse{}, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EncountersResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return EncountersResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return EncountersResponse{}, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EncountersResponse{}, err
	}

	var encounters EncountersResponse
	if err = json.Unmarshal(data, &encounters); err != nil {
		return EncountersResponse{}, err
	}

	c.pokeCache.Add(url, data)
	return encounters, nil
}
