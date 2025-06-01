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

// The lights are different for each bridge, we need to handle that, maybe update lights when we choose another bridge in the update method
// or we coudld move lights and groups to the bridge struct!
type model struct {
	win     window
	bridges bridge.Bridge
	groups  bridge.Groups
	lights  bridge.Lights
	event   bridge.Event
}

func (m model) Init() tea.Cmd {
	return bridge.Find_bridges
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return " "
}
