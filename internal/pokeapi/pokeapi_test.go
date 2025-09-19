package pokeapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arthwr/pokedex/internal/pokeapi"
)

func TestFetchLocations_CacheWorks(t *testing.T) {
	expected := pokeapi.LocationResponse{
		Count: 1,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "pallet-town", URL: "/location-area/1/"},
		},
	}

	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		_ = json.NewEncoder(w).Encode(expected)
	}))
	defer ts.Close()

	client := pokeapi.NewClient(50*time.Millisecond, 1*time.Minute)

	pageURL := ts.URL

	// First server call
	resp1, err := client.FetchLocations(&pageURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp1.Count != expected.Count {
		t.Errorf("expected Count=%d, got %d", expected.Count, resp1.Count)
	}
	if callCount != 1 {
		t.Errorf("expected 1 server call, got %d", callCount)
	}

	// Second server call with expected cache hit
	resp2, err := client.FetchLocations(&pageURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp2.Count != expected.Count {
		t.Errorf("expected Count=%d, got %d", expected.Count, resp2.Count)
	}
	if callCount != 1 {
		t.Errorf("expected cache hit (still 1 server call), got %d", callCount)
	}
}
