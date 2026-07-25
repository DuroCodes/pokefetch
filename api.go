package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	PokemonAPI        = "https://pokeapi.co/api/v2/pokemon"
	PokemonSpeciesAPI = "https://pokeapi.co/api/v2/pokemon-species"
	PokemonTypeAPI    = "https://pokeapi.co/api/v2/type"
)

func fetchPokemonData(pokemon string) PokemonData {
	return fetchData[PokemonData](fmt.Sprintf("%s/%s", PokemonAPI, pokemon))
}

func fetchPokemonSpeciesData(pokemon string) PokemonSpeciesData {
	return fetchData[PokemonSpeciesData](fmt.Sprintf("%s/%s", PokemonSpeciesAPI, pokemon))
}

func fetchPokemonTypeData(typeName string) PokemonTypeData {
	return fetchData[PokemonTypeData](fmt.Sprintf("%s/%s", PokemonTypeAPI, typeName))
}

func isValidPokemonName(name string) bool {
	resp, err := http.Get(fmt.Sprintf("%s/%s", PokemonAPI, name))

	if err != nil {
		return false
	}

	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func fetchData[T any](url string) T {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data T
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}

func fetchPokemonImage(url string) (string, bool) {
	resp, err := http.Get(url)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}
	return string(body), true
}

// loads sprite by PokeAPI form slug, falling back to species name
// forms like deoxys-normal are keyed as deoxys in colorscripts
func fetchColorscript(shiny bool, pokemonName, speciesName string) string {
	for _, name := range []string{pokemonName, speciesName} {
		if name == "" {
			continue
		}
		url := fmt.Sprintf(
			"https://gitlab.com/phoneybadger/pokemon-colorscripts/-/raw/main/colorscripts/small/%s/%s",
			getShinyOrRegular(shiny),
			name,
		)
		if image, ok := fetchPokemonImage(url); ok {
			return image
		}
	}
	log.Fatalf("failed to fetch colorscript for %q (species %q)", pokemonName, speciesName)
	return ""
}
