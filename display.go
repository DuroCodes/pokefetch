package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const ShinyChance = 0.5

func getEnglishName(names []PokemonSpeciesName) string {
	for _, name := range names {
		if name.Language.Name == "en" {
			return name.Name
		}
	}
	return ""
}

func getEnglishGenus(genera []PokemonSpeciesGenera) string {
	for _, genus := range genera {
		if genus.Language.Name == "en" {
			return genus.Genus
		}
	}
	return ""
}

func getEnglishFlavorText(entries []PokemonSpeciesFlavorText) string {
	for i := len(entries) - 1; i >= 0; i-- {
		if entries[i].Language.Name == "en" {
			return strings.ReplaceAll(entries[i].FlavorText, "\n", " ")
		}
	}
	return ""
}

func getTypeBadges(types []PokemonType) string {
	var badges []string
	for _, t := range types {
		color := pokemonTypeColor(t.TypeData.Name)
		badges = append(badges, createTextBadge(strings.ToUpper(t.TypeData.Name), color, false))
	}
	return strings.Join(badges, " ")
}

func createTextBadge(text string, color lipgloss.Color, bold bool) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(color).Padding(0, 1)
	if bold {
		style = style.Bold(true)
	}
	return style.Render(text)
}

func pokemonTypeColor(pokemonType string) lipgloss.Color {
	colors := map[string]lipgloss.Color{
		"normal":   lipgloss.Color("15"),
		"fire":     lipgloss.Color("9"),
		"water":    lipgloss.Color("12"),
		"electric": lipgloss.Color("11"),
		"grass":    lipgloss.Color("10"),
		"ice":      lipgloss.Color("14"),
		"fighting": lipgloss.Color("1"),
		"poison":   lipgloss.Color("5"),
		"ground":   lipgloss.Color("3"),
		"flying":   lipgloss.Color("13"),
		"psychic":  lipgloss.Color("13"),
		"bug":      lipgloss.Color("2"),
		"rock":     lipgloss.Color("3"),
		"ghost":    lipgloss.Color("5"),
		"dragon":   lipgloss.Color("13"),
		"dark":     lipgloss.Color("8"),
		"steel":    lipgloss.Color("15"),
		"fairy":    lipgloss.Color("13"),
	}

	if color, ok := colors[pokemonType]; ok {
		return color
	}
	return lipgloss.Color("0")
}

func getShinyOrRegular() string {
	if rollShiny() {
		return "shiny"
	}
	return "regular"
}

func rollShiny() bool {
	return rand.Float32() < ShinyChance
}

func formatPokemonInfo(dexBadge, name, genus, typeBadges, height, weight, flavorText string) string {
	title := formatTitle(dexBadge, name, genus)
	details := formatDetails(height, weight)
	flavorTextBox := formatFlavorText(flavorText)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		title,
		lipgloss.JoinVertical(
			lipgloss.Center,
			"",
			typeBadges,
			"",
			details,
			"",
			flavorTextBox,
		),
	)
}

func formatTitle(dexBadge, name, genus string) string {
	return dexBadge + lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")).Render(fmt.Sprintf(" %s - %s", name, genus))
}

func formatDetails(height, weight string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render("Height: "+lipgloss.NewStyle().Bold(true).Render(height)),
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render("	Weight: "+lipgloss.NewStyle().Bold(true).Render(weight)),
	)
}

func formatFlavorText(flavorText string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("15")).
		Foreground(lipgloss.Color("15")).
		Padding(0, 1).
		Width(40).
		Render(flavorText)
}
