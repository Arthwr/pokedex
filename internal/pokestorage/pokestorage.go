package pokestorage

import "fmt"

type PokemonData struct {
	Name           string
	BaseExperience int
}

type Storage struct {
	pokemon map[string]PokemonData
}

func NewStorage() *Storage {
	return &Storage{
		pokemon: make(map[string]PokemonData),
	}
}

func (s *Storage) Add(pokemon PokemonData) {
	s.pokemon[pokemon.Name] = pokemon
}

func (s *Storage) Get(name string) (PokemonData, bool) {
	p, ok := s.pokemon[name]
	return p, ok
}

func (s *Storage) List() []string {
	if len(s.pokemon) == 0 {
		return nil
	}

	names := make([]string, 0, len(s.pokemon))
	for name := range s.pokemon {
		names = append(names, name)
	}

	return names
}

func (s *Storage) PrintList() {
	names := s.List()
	if names == nil {
		fmt.Println("No Pokemon has been caught.")
		return
	}

	fmt.Println("List of stored Pokemon:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
}
