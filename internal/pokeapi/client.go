package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/arthwr/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

func (c *Client) doRequest(url string, target any) error {
	if val, ok := c.pokeCache.Get(url); ok {
		return json.Unmarshal(val, target)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.pokeCache.Add(url, data)
	return json.Unmarshal(data, target)
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: *pokecache.NewCache(cacheInterval),
	}
}
