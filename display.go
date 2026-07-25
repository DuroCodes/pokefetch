package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func getLocalizedName(names []PokemonSpeciesName, locale string) string {
	for _, name := range names {
		if name.Language.Name == locale {
			return name.Name
		}
	}
	if locale != "en" {
		return getLocalizedName(names, "en")
	}
	return ""
}

func getLocalizedGenus(genera []PokemonSpeciesGenera, locale string) string {
	for _, genus := range genera {
		if genus.Language.Name == locale {
			return genus.Genus
		}
	}
	if locale != "en" {
		return getLocalizedGenus(genera, "en")
	}
	return ""
}

func getLocalizedFlavorText(entries []PokemonSpeciesFlavorText, locale string) string {
	for i := len(entries) - 1; i >= 0; i-- {
		if entries[i].Language.Name == locale {
			return strings.ReplaceAll(entries[i].FlavorText, "\n", " ")
		}
	}
	if locale != "en" {
		return getLocalizedFlavorText(entries, "en")
	}
	return ""
}

func getUILabels(locale string) uiLabels {
	labels := map[string]uiLabels{
		"en":      {Height: "Height", Weight: "Weight"},
		"fr":      {Height: "Taille", Weight: "Poids"},
		"de":      {Height: "Größe", Weight: "Gewicht"},
		"es":      {Height: "Altura", Weight: "Peso"},
		"es-419":  {Height: "Altura", Weight: "Peso"},
		"it":      {Height: "Altezza", Weight: "Peso"},
		"ja":      {Height: "高さ", Weight: "重さ"},
		"ja-hrkt": {Height: "たかさ", Weight: "おもさ"},
		"ko":      {Height: "키", Weight: "몸무게"},
		"zh-hans": {Height: "身高", Weight: "体重"},
		"zh-hant": {Height: "身高", Weight: "體重"},
	}

	if label, ok := labels[locale]; ok {
		return label
	}
	return labels["en"]
}

func getTypeBadges(types []PokemonType, locale string) string {
	var badges []string
	for _, t := range types {
		typeSlug := t.TypeData.Name
		color := pokemonTypeColor(typeSlug)
		label := getLocalizedTypeName(typeSlug, locale)
		badges = append(badges, createTextBadge(strings.ToUpper(label), color, false))
	}
	return strings.Join(badges, " ")
}

func getLocalizedTypeName(typeSlug, locale string) string {
	typeData := fetchPokemonTypeData(typeSlug)
	for _, name := range typeData.Names {
		if name.Language.Name == locale {
			return name.Name
		}
	}
	if locale != "en" {
		return getLocalizedTypeName(typeSlug, "en")
	}
	return typeSlug
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
		"steel":    lipgloss.Color("7"),
		"fairy":    lipgloss.Color("13"),
	}

	if color, ok := colors[pokemonType]; ok {
		return color
	}
	return lipgloss.Color("0")
}

func getShinyOrRegular(shiny bool) string {
	if shiny {
		return "shiny"
	}
	return "regular"
}

func getShinyOrRegularColor(shiny bool) lipgloss.Color {
	if shiny {
		return lipgloss.Color("11")
	}
	return lipgloss.Color("15")
}

func rollShiny(chance float64) bool {
	return rand.Float64() < chance
}

func formatPokemonInfo(dexBadge, name, genus, typeBadges, heightLabel, weightLabel, height, weight, flavorText string, mainColor lipgloss.Color) string {
	title := formatTitle(dexBadge, name, genus, mainColor)
	details := formatDetails(heightLabel, weightLabel, height, weight)
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

func formatTitle(dexBadge, name, genus string, mainColor lipgloss.Color) string {
	return dexBadge + lipgloss.NewStyle().Bold(true).Foreground(mainColor).Render(fmt.Sprintf(" %s - %s", name, genus))
}

func formatDetails(heightLabel, weightLabel, height, weight string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render(heightLabel+": "+lipgloss.NewStyle().Bold(true).Render(height)),
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render("	"+weightLabel+": "+lipgloss.NewStyle().Bold(true).Render(weight)),
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
