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

func fetchPokemonData(dexID int) PokemonData {
	return fetchData[PokemonData](fmt.Sprintf("%s/%d", PokemonAPI, dexID))
}

func fetchPokemonSpeciesData(dexID int) PokemonSpeciesData {
	return fetchData[PokemonSpeciesData](fmt.Sprintf("%s/%d", PokemonSpeciesAPI, dexID))
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
