package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	pokemonID := flag.Int("id", 0, "Pokémon ID to fetch. If not provided, a random ID will be used.")
	pokemonName := flag.String("name", "", "Pokémon name to fetch. If not provided, a random Pokémon will be used.")
	shinyFlag := flag.Float64("shiny", 0.5, "Odds of the Pokémon being shiny. Default is 0.5.")
	localeFlag := flag.String("locale", "en", "Locale to fetch the pokemon name and flavor text for. Default is 'en'.")

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

	name := getEnglishName(pokemonSpeciesData.Names)
	localeName := getLocalizedName(pokemonSpeciesData.Names, locale)
	weight := fmt.Sprintf("%.1fkg", weightKg)
	height := fmt.Sprintf("%.1fm", float32(pokemonData.Height)/10)
	genus := getLocalizedGenus(pokemonSpeciesData.Genera, locale)
	flavorText := getLocalizedFlavorText(pokemonSpeciesData.FlavorTextEntries, locale)
	typeBadges := getTypeBadges(pokemonData.Types)
	isShiny := rollShiny(shinyOdds)

	mainColor := getShinyOrRegularColor(isShiny)
	dexBadge := createTextBadge(fmt.Sprintf("No.%03d", dexId), mainColor, true)

	pokemonImageURL := fmt.Sprintf("https://gitlab.com/phoneybadger/pokemon-colorscripts/-/raw/main/colorscripts/small/%s/%s", getShinyOrRegular(isShiny), strings.ToLower(name))

	pokemonImage := fetchPokemonImage(pokemonImageURL)
	pokemonInfo := formatPokemonInfo(dexBadge, localeName, genus, typeBadges, height, weight, flavorText, mainColor)

	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().MarginRight(4).Render(pokemonImage),
		pokemonInfo,
	)

	println(output)
}
