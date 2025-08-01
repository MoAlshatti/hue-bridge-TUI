package main

import (
	"log"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/MoAlshatti/hue-bridge-TUI/internal/view"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const logFileName = ".debug.log"

func main() {
	file, err := tea.LogToFile(logFileName, "")
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

type model struct {
	win      window
	log      *bridge.LogFile
	userpage bridge.UserPage
	bridge   bridge.Bridge
	user     bridge.User
	groups   bridge.Groups
	scenes   bridge.Scenes
	lights   bridge.Lights
	event    bridge.Event
}

func initalModel() model {
	return model{
		userpage: bridge.UserPage{Items: [2]string{"Quit", "Done!"}},
		lights:   bridge.Lights{Cursor: 0},
		groups:   bridge.Groups{Cursor: 0},
		bridge:   bridge.Bridge{Selected: true},
		log:      &bridge.LogFile{},
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(bridge.Init_client)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.win.width = msg.Width
		m.win.height = msg.Height
	case bridge.BridgeFoundMsg:
		m.log.Log_Print("Bridge found!")
		m.bridge = bridge.Bridge(msg)
		m.bridge.Selected = true
		m.event = bridge.FindingUser
		return m, bridge.Find_User(m.bridge)
	case bridge.NoBridgeFoundMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.NoUserFoundMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		m.event = bridge.RequestPressButton
	case bridge.UserFoundMsg:
		m.log.Log_Print("User found!")
		m.event = bridge.FetchingLights
		m.user.Username = string(msg)

		return m, tea.Batch(bridge.Fetch_lights(m.bridge, m.user.Username, m.log),
			bridge.Fetch_groups(m.bridge, m.user.Username, m.log),
			bridge.Fetch_Scenes(m.bridge, m.user.Username, m.log))
	case bridge.ClientCreatedMsg:
		return m, bridge.Find_bridges
	case bridge.NoClientCreatedMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.UserCreatedMsg:
		m.user.Username = string(msg)
		m.event = bridge.FetchingLights
		return m, tea.Batch(bridge.Save_Username(string(msg)),
			bridge.Fetch_lights(m.bridge, m.user.Username, m.log),
			bridge.Fetch_groups(m.bridge, m.user.Username, m.log),
			bridge.Fetch_Scenes(m.bridge, m.user.Username, m.log))
	case bridge.UserCreationFailedMsg:
		m.log.Log_Print("Failed to create user, err: ", bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.ButtonNotPressed:
		m.log.Log_Print(string(msg))
	case bridge.FailedFetchingLightsMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.LightsMsg:

		m.log.Log_Print("Lights found")
		m.lights.Items = []bridge.Light(msg)
		m.event = bridge.DisplayingLights
	case bridge.FailedToFetchGroupsMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.GroupsMsg:
		m.groups.Items = []bridge.Group(msg)
	case bridge.FailedToFetchScenesMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.ScenesMsg:
		m.scenes.Items = []bridge.Scene(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			log.Println("bridge panel")
			m.groups.Selected, m.lights.Selected, m.scenes.Selected = false, false, false
			m.bridge.Selected = true
		case "2":
			m.bridge.Selected, m.lights.Selected, m.scenes.Selected = false, false, false
			m.groups.Selected = true
		case "3":
			m.bridge.Selected, m.groups.Selected, m.scenes.Selected = false, false, false
			m.lights.Selected = true
		case "4":
			m.bridge.Selected, m.groups.Selected, m.lights.Selected = false, false, false
			m.scenes.Selected = true
		case "j", "down":
			if m.groups.Selected {
				if m.groups.Cursor < len(m.groups.Items)-1 {
					m.groups.Cursor++
				}
			} else if m.lights.Selected {
				if m.lights.Cursor < len(m.lights.Items)-1 {
					m.lights.Cursor++
				}
			} else if m.scenes.Selected {
				if m.scenes.Cursor < len(m.scenes.Items)-1 {
					m.scenes.Cursor++
				}
			}
		case "k", "up":
			if m.groups.Selected {
				if m.groups.Cursor > 0 {
					m.groups.Cursor--
				}
			} else if m.lights.Selected {
				if m.lights.Cursor > 0 {
					m.lights.Cursor--
				}
			} else if m.scenes.Selected {
				if m.scenes.Cursor > 0 {
					m.scenes.Cursor--
				}
			}

		case "h", "left":
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
		quitOpt = view.Render_userpage_options(m.userpage.Items[bridge.Quit], m.userpage.Cursor == bridge.Quit)
		pressOpt = view.Render_userpage_options(m.userpage.Items[bridge.PressTheButton], m.userpage.Cursor != bridge.Quit)

		userpage := view.Render_userpage(title, quitOpt, pressOpt)
		return lipgloss.Place(m.win.width, m.win.height, lipgloss.Center, lipgloss.Center, userpage)
	case bridge.DisplayingLights:
		title := view.Render_bridge_title("Hue Bridge", m.win.width, m.win.height)

		bridgepanel := view.Render_bridge_panel(title, m.bridge.Selected, m.win.width, m.win.height)

		var groups []string
		for i, v := range m.groups.Items {
			groups = append(groups, view.Render_group_title(v.Metadata.Name, i == m.groups.Cursor, m.win.width, m.win.height))
		}
		grouppanel := view.Render_group_panel(groups, m.groups.Selected, m.groups.Cursor, m.win.width, m.win.height)

		var lights []string
		for i, v := range m.lights.Items {
			lights = append(lights, view.Render_light_title(v.Metadata.Name,
				v.Dimming.Brightness,
				v.On, i == m.lights.Cursor && m.lights.Selected, m.win.width, m.win.height))
		}
		lightpanel := view.Render_light_panel(lights, m.lights.Selected, m.lights.Cursor, m.win.width, m.win.height)

		var scenes []string
		for i, v := range m.scenes.Items {
			scenes = append(scenes, view.Render_scene_title(v.Name,
				v.Active,
				i == m.scenes.Cursor && m.scenes.Selected, m.win.width, m.win.height))
		}
		scenePanel := view.Render_scene_panel(scenes, m.scenes.Selected, m.scenes.Cursor, m.win.width, m.win.height)

		var details string
		if m.bridge.Selected {
			details = view.Render_bridge_details(m.bridge, m.win.width, m.win.height)
		} else if m.groups.Selected {
			details = view.Render_group_details(m.groups.Items[m.groups.Cursor], m.win.width, m.win.height)
		} else if m.lights.Selected {
			details = view.Render_light_details(m.lights.Items[m.lights.Cursor], m.win.width, m.win.height)
		} else if m.scenes.Selected {
			details = view.Render_scene_details(m.scenes.Items[m.scenes.Cursor], m.win.width, m.win.height)
		}

		detailsPanel := view.Render_details_panel(details, m.win.width, m.win.height)

		logcontent := view.Render_log_title(m.log.Content, m.win.width, m.win.height)
		logPanel := view.Render_log_panel(logcontent, m.win.width, m.win.height)

		leftSide := lipgloss.JoinVertical(lipgloss.Left, bridgepanel, grouppanel, lightpanel, scenePanel)
		rightSide := lipgloss.JoinVertical(lipgloss.Bottom, detailsPanel, logPanel)

		output := lipgloss.JoinHorizontal(lipgloss.Top, leftSide, rightSide)
		return output
	}
	return " "
}
