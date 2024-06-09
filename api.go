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
)

func fetchPokemonData(pokemon string) PokemonData {
	return fetchData[PokemonData](fmt.Sprintf("%s/%s", PokemonAPI, pokemon))
}

func fetchPokemonSpeciesData(pokemon string) PokemonSpeciesData {
	return fetchData[PokemonSpeciesData](fmt.Sprintf("%s/%s", PokemonSpeciesAPI, pokemon))
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

func fetchPokemonImage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
