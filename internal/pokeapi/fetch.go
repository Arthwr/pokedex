package pokeapi

func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {
	var locations LocationResponse
	err := c.doRequest(baseURL+"/location-area?offset=0&limit=20", &locations)
	return locations, err
}

func (c *Client) FetchEncountersFromLocation(locationID string) (EncountersResponse, error) {
	var encounters EncountersResponse
	err := c.doRequest(baseURL+"/location-area/"+locationID, &encounters)
	return encounters, err
}

func (c *Client) FetchPokemon(pokemonID string) (PokemonResponse, error) {
	var pokemon PokemonResponse
	err := c.doRequest(baseURL+"/pokemon/"+pokemonID, &pokemon)
	return pokemon, err
}
