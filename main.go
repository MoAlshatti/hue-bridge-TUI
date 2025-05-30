package main

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

}

type window struct {
	width  int
	height int
}

type model struct {
	win    window
	groups bridge.Groups
	lights bridge.Lights
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return " "
}
