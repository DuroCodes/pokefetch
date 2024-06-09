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
	shinyFlag := flag.Bool("shiny", false, "Should the Pokémon be shiny?")

	flag.Parse()

	dexId := *pokemonID
	pokeName := *pokemonName
	isShiny := *shinyFlag

	if dexId == 0 {
		if pokeName == "" {
			dexId = rand.Intn(898) + 1
		} else if isValidPokemonName(pokeName) {
			dexId = fetchPokemonData(pokeName).Id
		} else {
			dexId = rand.Intn(898) + 1
		}
	}

	if !isShiny {
		isShiny = rollShiny()
	}

	dexIdStr := fmt.Sprintf("%d", dexId)
	pokemonData := fetchPokemonData(dexIdStr)
	pokemonSpeciesData := fetchPokemonSpeciesData(dexIdStr)

	name := getEnglishName(pokemonSpeciesData.Names)
	weight := fmt.Sprintf("%dkg", pokemonData.Weight)
	height := fmt.Sprintf("%.1fm", float32(pokemonData.Height)/10)
	genus := getEnglishGenus(pokemonSpeciesData.Genera)
	flavorText := getEnglishFlavorText(pokemonSpeciesData.FlavorTextEntries)
	typeBadges := getTypeBadges(pokemonData.Types)

	mainColor := getShinyOrRegularColor(isShiny)
	dexBadge := createTextBadge(fmt.Sprintf("No.%03d", dexId), mainColor, true)

	pokemonImageURL := fmt.Sprintf("https://gitlab.com/phoneybadger/pokemon-colorscripts/-/raw/main/colorscripts/small/%s/%s", getShinyOrRegular(isShiny), strings.ToLower(name))

	pokemonImage := fetchPokemonImage(pokemonImageURL)
	pokemonInfo := formatPokemonInfo(dexBadge, name, genus, typeBadges, height, weight, flavorText, mainColor)

	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().MarginRight(4).Render(pokemonImage),
		pokemonInfo,
	)

	println(output)
}
