package main

type PokemonGeneralData struct {
	Name string `json:"name"`
}

type PokemonType struct {
	TypeData PokemonGeneralData `json:"type"`
}

type PokemonData struct {
	Types  []PokemonType `json:"types"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Id     int           `json:"id"`
}

type PokemonSpeciesGenera struct {
	Genus    string             `json:"genus"`
	Language PokemonGeneralData `json:"language"`
}

type PokemonSpeciesName struct {
	Name     string             `json:"name"`
	Language PokemonGeneralData `json:"language"`
}

type PokemonSpeciesFlavorText struct {
	FlavorText string             `json:"flavor_text"`
	Language   PokemonGeneralData `json:"language"`
}

type PokemonSpeciesData struct {
	Names             []PokemonSpeciesName       `json:"names"`
	Genera            []PokemonSpeciesGenera     `json:"genera"`
	FlavorTextEntries []PokemonSpeciesFlavorText `json:"flavor_text_entries"`
}
