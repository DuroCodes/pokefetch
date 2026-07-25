package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	pokemonID := flag.Int("id", 0, "Pokémon ID to fetch. If not provided, a random ID will be used.")
	pokemonName := flag.String("name", "", "Pokémon name to fetch. If not provided, a random Pokémon will be used.")
	shinyFlag := flag.Float64("shiny", 0.5, "Odds of the Pokémon being shiny. Default is 0.5.")
	localeFlag := flag.String("locale", "en", "Locale for name, genus, flavor text, types, and labels (default: en)")

	flag.Parse()

	dexId := *pokemonID
	pokeName := *pokemonName
	shinyOdds := *shinyFlag
	locale := *localeFlag

	if dexId == 0 {
		if pokeName == "" || !isValidPokemonName(pokeName) {
			dexId = rand.Intn(898) + 1
		} else {
			dexId = fetchPokemonData(pokeName).Id
		}
	}

	dexIdStr := fmt.Sprintf("%d", dexId)
	pokemonData := fetchPokemonData(dexIdStr)
	pokemonSpeciesData := fetchPokemonSpeciesData(dexIdStr)

	weightKg := float64(pokemonData.Weight) / 10.0

	localeName := getLocalizedName(pokemonSpeciesData.Names, locale)
	weight := fmt.Sprintf("%.1fkg", weightKg)
	height := fmt.Sprintf("%.1fm", float32(pokemonData.Height)/10)
	genus := getLocalizedGenus(pokemonSpeciesData.Genera, locale)
	flavorText := getLocalizedFlavorText(pokemonSpeciesData.FlavorTextEntries, locale)
	typeBadges := getTypeBadges(pokemonData.Types, locale)
	labels := getUILabels(locale)
	isShiny := rollShiny(shinyOdds)

	mainColor := getShinyOrRegularColor(isShiny)
	dexBadge := createTextBadge(fmt.Sprintf("No.%03d", dexId), mainColor, true)

	pokemonImage := fetchColorscript(isShiny, pokemonData.Name, pokemonSpeciesData.Name)
	pokemonInfo := formatPokemonInfo(dexBadge, localeName, genus, typeBadges, labels.Height, labels.Weight, height, weight, flavorText, mainColor)

	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().MarginRight(4).Render(pokemonImage),
		pokemonInfo,
	)

	println(output)
}
