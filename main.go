package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	dexId := rand.Intn(898) + 1
	pokemonData := fetchPokemonData(dexId)
	pokemonSpeciesData := fetchPokemonSpeciesData(dexId)

	pokemonName := getEnglishName(pokemonSpeciesData.Names)
	weight := fmt.Sprintf("%dkg", pokemonData.Weight)
	height := fmt.Sprintf("%.1fm", float32(pokemonData.Height)/10)
	genus := getEnglishGenus(pokemonSpeciesData.Genera)
	flavorText := getEnglishFlavorText(pokemonSpeciesData.FlavorTextEntries)
	typeBadges := getTypeBadges(pokemonData.Types)
	dexBadge := createTextBadge(fmt.Sprintf("No.%03d", dexId), lipgloss.Color("15"), true)

	pokemonImageURL := fmt.Sprintf("https://gitlab.com/phoneybadger/pokemon-colorscripts/-/raw/main/colorscripts/small/%s/%s",
		getShinyOrRegular(), strings.ToLower(pokemonName))

	pokemonImage := fetchPokemonImage(pokemonImageURL)
	pokemonInfo := formatPokemonInfo(dexBadge, pokemonName, genus, typeBadges, height, weight, flavorText)

	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().MarginRight(4).Render(pokemonImage),
		pokemonInfo,
	)

	println(output)
}
