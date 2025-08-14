package view

import (
	"math/rand/v2"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
)

func Render_loading_text(e bridge.Event) string {
	style := lipgloss.NewStyle().Bold(true).Foreground(green)

	switch e {
	case bridge.FindingBridge:
		return style.Render("Looking for the bridge...")
	case bridge.FindingUser:
		return style.Render("Searching for the app key...")
	case bridge.CreateUser:
		return style.Render("Key not found, Creating user...")
	case bridge.FetchingGroups:
		return style.Render("User found! Fetching Groups...")
	case bridge.FetchingLights:
		return style.Render("Groups Fetched! Fetching Lights and scenes...")
	}

	facts := []string{
		"Cats can't taste sweetness,",
		"Avocados are toxic to birds",
		"Peanuts aren't nuts",
		"sharks have eyelids",
		"space smells like burnt steak",
		"Bananas are berries, but strawberries aren't",
		"Pineapples take  2 years to grow",
		"Lobsters pee from their faces",
		"Male seahorses give birth",
	}
	return style.Render("Did you know that", facts[rand.IntN(len(facts)-1)])
}
