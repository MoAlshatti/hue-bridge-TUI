package main

import (
	"log"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	file, err := tea.LogToFile(".debug.log", "")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	p := tea.NewProgram(initalModel(), tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

type window struct {
	width  int
	height int
}

// The lights are different for each bridge, we need to handle that, maybe update lights when we choose another bridge in the update method
// or we coudld move lights and groups to the bridge struct!
type model struct {
	win    window
	bridge bridge.Bridge
	groups bridge.Groups
	lights bridge.Lights
	event  bridge.Event
}

func initalModel() model {
	return model{} // TODO
}

func (m model) Init() tea.Cmd {
	return bridge.Find_bridges
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.win.width = msg.Width
		m.win.height = msg.Height
	case bridge.BridgeFoundMsg:
		log.Println("Bridge Found, ", msg.Ip_addr)
		m.bridge = bridge.Bridge(msg)
		m.event = bridge.FindingUser
		return m, bridge.Find_User(m.bridge)
	case bridge.NoBridgeFoundMsg:
		log.Println(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.NoUserFoundMsg:
		log.Println(bridge.ErrMsg(msg))
		m.event = bridge.RequestPressButton
		// here you should trigger the user to press the bridge button
	case bridge.UserFoundMsg:
		log.Println(string(msg))
		//start displaying

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return " "
}
