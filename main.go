package main

import (
	"log"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/MoAlshatti/hue-bridge-TUI/internal/view"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	win      window
	userpage bridge.UserPage
	bridge   bridge.Bridge
	user     bridge.User
	groups   bridge.Groups
	lights   bridge.Lights
	event    bridge.Event
}

func initalModel() model {
	return model{
		userpage: bridge.UserPage{Items: [2]string{"Quit", "Done!"}},
	} // TODO
}

func (m model) Init() tea.Cmd {
	return bridge.Init_client
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
		m.event = bridge.FetchingLights
		log.Println("user found!")
		m.user.Username = string(msg)

		// retreive lights here

		//start displaying

	case bridge.ClientCreatedMsg:
		return m, bridge.Find_bridges
	case bridge.NoClientCreatedMsg:
		log.Println(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.UserCreatedMsg:
		m.event = bridge.FetchingLights
		//saved user to a file

		// send retrieve resources cmd

		// batch cmds
		return m, tea.Batch(bridge.Save_Username(string(msg))) //add retreive resoruscres later
	case bridge.UserCreationFailedMsg:
		log.Println("Failed to create user, quitting...", bridge.ErrMsg(msg))
		return m, tea.Quit

	//case bridge.ButtonNotPressed:
	//recall create user function after displaying a display button message

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "h", "left":
			//
			switch m.event {
			case bridge.RequestPressButton:
				if m.userpage.Cursor <= 1 {
					m.userpage.Cursor++
				}
			}
		case "l", "right":
			switch m.event {
			case bridge.RequestPressButton:
				if m.userpage.Cursor > 0 {
					m.userpage.Cursor--
				}
			}
		case "enter":
			switch m.event {
			case bridge.RequestPressButton:
				//
				if m.userpage.Cursor == bridge.Quit {
					return m, tea.Quit
				} else if m.userpage.Cursor == bridge.PressTheButton {
					return m, bridge.Create_User(m.bridge)
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	switch event := m.event; event {
	case bridge.RequestPressButton:
		//
		title := view.Render_userpage_title("Press the hue bridge button!")
		var (
			quitOpt, pressOpt string
		)
		if m.userpage.Cursor == bridge.Quit {
			quitOpt = view.Render_userpage_options(m.userpage.Items[bridge.Quit], true)
			pressOpt = view.Render_userpage_options(m.userpage.Items[bridge.PressTheButton], false)
		} else {
			quitOpt = view.Render_userpage_options(m.userpage.Items[bridge.Quit], false)
			pressOpt = view.Render_userpage_options(m.userpage.Items[bridge.PressTheButton], true)
		}
		userpage := view.Render_userpage(title, quitOpt, pressOpt)
		return lipgloss.Place(m.win.width, m.win.height, lipgloss.Center, lipgloss.Center, userpage)

	}
	return " "
}
