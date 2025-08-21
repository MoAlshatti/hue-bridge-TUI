package main

import (
	"log"

	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/MoAlshatti/hue-bridge-TUI/internal/view"
	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

const logFileName = ".debug.log"

var p *tea.Program

func main() {
	file, err := tea.LogToFile(logFileName, "")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	p = tea.NewProgram(initalModel(), tea.WithAltScreen())
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
	win        window
	log        *bridge.LogFile
	userpage   bridge.UserPage
	bridge     bridge.Bridge
	user       bridge.User
	groups     bridge.Groups
	scenes     bridge.Scenes
	brightness bridge.BrightnessModal
	color      bridge.ColorModal
	lights     bridge.Lights
	panel      bridge.Panel
	event      bridge.Event
}

func initalModel() model {
	bm := bridge.BrightnessModal{}
	bm.Init()
	cm := bridge.ColorModal{List: bridge.Initialize_list()}
	view.Apply_list_style(&cm.List)
	return model{
		userpage:   bridge.UserPage{Items: [2]string{"Quit", "Done!"}},
		lights:     bridge.Lights{Cursor: 0},
		groups:     bridge.Groups{Cursor: 0},
		log:        &bridge.LogFile{},
		brightness: bm,
		color:      cm,
		panel:      bridge.BridgePanel,
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
		view.Update_list_size(&m.color.List, msg.Width, msg.Height)
	case bridge.BridgeFoundMsg:
		m.bridge = bridge.Bridge(msg)
		m.event = bridge.FindingUser
		return m, bridge.Find_User(m.bridge)
	case bridge.NoBridgeFoundMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.NoUserFoundMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		m.event = bridge.RequestPressButton
	case bridge.UserFoundMsg:
		m.event = bridge.FetchingGroups
		m.user.Username = string(msg)
		return m, tea.Batch(bridge.Fetch_groups(m.bridge, m.user.Username, m.log),
			bridge.Initiate_sse(m.bridge, m.user.Username, p))
	case bridge.ClientCreatedMsg:
		return m, bridge.Find_bridges
	case bridge.NoClientCreatedMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.UserCreatedMsg:
		m.log.Log_Print("User Created Successfully!")
		m.user.Username = string(msg)
		m.event = bridge.FetchingGroups
		return m, tea.Batch(bridge.Save_Username(string(msg)),
			bridge.Fetch_groups(m.bridge, m.user.Username, m.log),
			bridge.Initiate_sse(m.bridge, m.user.Username, p))
	case bridge.UserCreationFailedMsg:
		m.log.Log_Print("Failed to create user, err: ", bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.ButtonNotPressed:
		m.log.Log_Print(string(msg))
	case bridge.FailedFetchingLightsMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.SseFailedMsg:
		m.log.Log_Print("sse failed: ", msg)
	case bridge.LightsMsg:
		m.log.Log_Print("Lights Fetched!")
		m.lights.AllItems = []bridge.Light(msg)
		// populate Items based on the chosen room
		bridge.Filter_lights(&m.lights, m.groups)
		m.event = bridge.DisplayingLights
		return m, bridge.Fetch_connectivity(m.bridge, m.user.Username)
	case bridge.FailedToFetchGroupsMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.GroupsMsg:
		m.log.Log_Print("Groups Fetched!")
		m.groups.Items = append(m.groups.Items, bridge.Group{ID: "None", Metadata: struct {
			Name      string
			Archetype string
		}{Name: "None"}})
		m.groups.Items = append(m.groups.Items, []bridge.Group(msg)...)
		m.event = bridge.FetchingLights
		return m, tea.Batch(bridge.Fetch_lights(m.bridge, m.user.Username, m.log),
			bridge.Fetch_Scenes(m.bridge, m.user.Username, m.log))
	case bridge.FailedToFetchScenesMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
		return m, tea.Quit
	case bridge.ScenesMsg:
		m.log.Log_Print("Scenes Fetched!")
		m.scenes.AllItems = []bridge.Scene(msg)
		bridge.Filter_scenes(&m.scenes, m.groups)
	case bridge.ResourceErrMsg:
		m.log.Log_Print(bridge.ErrMsg(msg))
	case bridge.ResourceSuccessMsg:
		m.log.Log_Print(string(msg))
	case bridge.ConnectivityMsg:
		bridge.Sort_Connectivity(&m.lights, msg)
	case bridge.StateUpdate:
		switch msg.Type {
		case "light":
			bridge.Update_light_status(m.lights.AllItems, msg)
		case "grouped_light":
			bridge.Update_group_status(m.groups.Items, msg)
		}
	case bridge.BriUpdate:
		switch msg.Type {
		case "light":
			bridge.Update_light_brightness(m.lights.AllItems, msg)
		case "grouped_light":
			bridge.Update_group_brightness(m.groups.Items, msg)
		}
	case bridge.ColorUpdate:
		bridge.Update_light_color(m.lights.AllItems, msg)
	case bridge.SceneStateUpdate:
		bridge.Update_scene_status(m.scenes.AllItems, msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			switch m.event {
			case bridge.DisplayingBrightness, bridge.DisplayingColors:
				// do nothing
			default:
				return m, tea.Quit
			}
		case "ctrl+c":
			return m, tea.Quit
		case "1":
			if m.event == bridge.DisplayingLights {
				m.panel = bridge.BridgePanel
			}
		case "2":
			if m.event == bridge.DisplayingLights {
				m.panel = bridge.GroupPanel
			}
		case "3":
			if m.event == bridge.DisplayingLights {
				m.panel = bridge.LightPanel
			}
		case "4":
			if m.event == bridge.DisplayingLights {
				m.panel = bridge.ScenePanel
			}
		case "j", "down":
			switch m.event {
			case bridge.DisplayingLights:
				if m.panel == bridge.GroupPanel {
					bridge.Increment_cursor(&m.groups)
					bridge.Filter_lights(&m.lights, m.groups)
					bridge.Filter_scenes(&m.scenes, m.groups)
				} else if m.panel == bridge.LightPanel {
					bridge.Increment_cursor(&m.lights)
				} else if m.panel == bridge.ScenePanel {
					bridge.Increment_cursor(&m.scenes)
				}
			case bridge.DisplayingBrightness:
				//
			case bridge.DisplayingColors:
				//
			}

		case "k", "up":
			switch m.event {
			case bridge.DisplayingLights:
				if m.panel == bridge.GroupPanel {
					bridge.Decrement_cusror(&m.groups)
					bridge.Filter_lights(&m.lights, m.groups)
					bridge.Filter_scenes(&m.scenes, m.groups)
				} else if m.panel == bridge.LightPanel {
					bridge.Decrement_cusror(&m.lights)
				} else if m.panel == bridge.ScenePanel {
					bridge.Decrement_cusror(&m.scenes)
				}
			case bridge.DisplayingBrightness:
				//
			case bridge.DisplayingColors:
				//
			}
		case "h", "left":
			switch m.event {
			case bridge.RequestPressButton:
				bridge.Increment_cursor(&m.userpage)
			case bridge.DisplayingLights:
				switch m.panel {
				case bridge.LightPanel:
					light := *m.lights.Items[m.lights.Cursor]
					if light.Dimming.Brightness > 0 && light.On && light.Connected {
						bri := max(light.Dimming.Brightness-15, 0.0)
						return m, bridge.Change_light_brightness(m.bridge, light, bri, m.user.Username)
					}
				case bridge.GroupPanel:
					group := m.groups.Items[m.groups.Cursor]
					if group.Brightness > 0 && group.On {
						bri := max(group.Brightness-15, 0.0)
						return m, bridge.Change_group_brightness(m.bridge, group, bri, m.user.Username)
					}
				}
			case bridge.DisplayingBrightness:
				//
			case bridge.DisplayingColors:
				//
			}
		case "c":
			if m.panel == bridge.LightPanel && m.event == bridge.DisplayingLights {
				light := m.lights.Items[m.lights.Cursor]
				if light.Connected && (light.Color.X != 0 && light.Color.Y != 0) {
					m.event = bridge.DisplayingColors
					m.color.List.FilterInput.Focus()
					return m, nil
				}
			}
		case "b":
			if m.event == bridge.DisplayingBrightness {
				m.brightness.Off()
				m.event = bridge.DisplayingLights
				m.brightness.Input.Err = nil
				return m, nil
			}
			switch m.panel {
			// could potentially be reformatted for conciseness
			case bridge.LightPanel:
				if m.event == bridge.DisplayingLights {
					light := m.lights.Items[m.lights.Cursor]
					if light.Connected {
						m.event = bridge.DisplayingBrightness
						m.brightness.On()
					}
				}
			case bridge.GroupPanel:
				if m.event == bridge.DisplayingLights && m.groups.Cursor > 0 {
					m.event = bridge.DisplayingBrightness
					m.brightness.On()
				}
			}
			return m, nil
		case "esc":
			switch m.event {
			case bridge.DisplayingColors:
				//
				if m.color.List.SettingFilter() {
					m.color.List.SetFilterState(list.Unfiltered)
					m.color.List.ResetFilter()
					m.color.List.FilterInput.Blur()
					m.color.List.ResetSelected()
					return m, nil
				}
				m.event = bridge.DisplayingLights
				return m, nil
			case bridge.DisplayingBrightness:
				m.brightness.Off()
				m.brightness.Input.Err = nil
				m.event = bridge.DisplayingLights
				return m, nil
			}
		case "l", "right":
			switch m.event {
			case bridge.RequestPressButton:
				bridge.Decrement_cusror(&m.userpage)
			case bridge.DisplayingLights:
				switch m.panel {
				case bridge.LightPanel:
					light := *m.lights.Items[m.lights.Cursor]
					if light.Dimming.Brightness < 100 && light.On && light.Connected {
						bri := min(light.Dimming.Brightness+15, 100.0)
						return m, bridge.Change_light_brightness(m.bridge, light, bri, m.user.Username)
					}
				case bridge.GroupPanel:
					group := m.groups.Items[m.groups.Cursor]
					if group.Brightness < 100 && group.On {
						bri := min(group.Brightness+15, 100.0)
						return m, bridge.Change_group_brightness(m.bridge, group, bri, m.user.Username)
					}
				}
			}
		case "enter":
			switch m.event {
			case bridge.RequestPressButton:
				switch m.userpage.Cursor {
				case bridge.Quit:
					return m, tea.Quit
				case bridge.PressTheButton:
					return m, bridge.Create_User(m.bridge)
				}
			case bridge.DisplayingLights:
				switch m.panel {
				case bridge.LightPanel:
					light := *m.lights.Items[m.lights.Cursor]
					if light.Connected {
						return m, bridge.Change_light_state(m.bridge, light, !light.On, m.user.Username)
					}
				case bridge.GroupPanel:
					if m.groups.Cursor > 0 {
						group := m.groups.Items[m.groups.Cursor]
						return m, bridge.Change_group_state(m.bridge, group, !group.On, m.user.Username)
					}
				case bridge.ScenePanel:
					scene := *m.scenes.Items[m.scenes.Cursor]
					if !scene.Active {
						return m, bridge.Pick_scene(m.bridge, scene, m.user.Username)
					}
				}
			case bridge.DisplayingBrightness:
				// send the command then reset the textinput value
				if m.brightness.Input.Err == nil {
					bri, err := m.brightness.Parse()
					if err == nil {
						m.brightness.Off()
						m.event = bridge.DisplayingLights
						switch m.panel {
						case bridge.LightPanel:
							light := m.lights.Items[m.lights.Cursor]
							return m, bridge.Change_light_brightness(m.bridge, *light, bri, m.user.Username)
						case bridge.GroupPanel:
							group := m.groups.Items[m.groups.Cursor]
							return m, bridge.Change_group_brightness(m.bridge, group, bri, m.user.Username)
						}
					}
				}
			case bridge.DisplayingColors:
				color, ok := m.color.List.SelectedItem().(bridge.Color)
				if ok && !m.color.List.SettingFilter() {
					light := *m.lights.Items[m.lights.Cursor]
					m.color.List.FilterInput.Blur()
					m.color.List.ResetFilter()
					m.color.List.ResetSelected()
					m.event = bridge.DisplayingLights
					return m, bridge.Change_light_color(m.bridge, light, color, m.user.Username)
				}
			}
		}
	}
	cmds := make([]tea.Cmd, 2)
	switch m.event {
	case bridge.DisplayingColors:
		m.color.List, cmds[1] = m.color.List.Update(msg)
	case bridge.DisplayingBrightness:
		*m.brightness.Input, cmds[0] = m.brightness.Input.Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch event := m.event; event {
	case bridge.RequestPressButton:
		userpage := view.Render_userpage(m.userpage)
		return lipgloss.Place(m.win.width, m.win.height, lipgloss.Center, lipgloss.Center, userpage)
	case bridge.DisplayingLights, bridge.DisplayingBrightness, bridge.DisplayingColors:
		bridgepanel := view.Render_bridge(m.bridge, m.panel, m.win.width, m.win.height)

		grouppanel := view.Render_group(m.groups, m.panel, m.win.width, m.win.height)

		lightpanel := view.Render_lights(m.lights, m.groups, m.panel, m.win.width, m.win.height)

		scenePanel := view.Render_scenes(m.scenes, m.panel, m.win.width, m.win.height)

		detailsPanel := view.Render_details(m.bridge, m.groups, m.lights, m.scenes, m.panel, m.win.width, m.win.height)

		logcontent := view.Render_log_title(m.log.Content, m.win.width, m.win.height)
		logPanel := view.Render_log_panel(logcontent, m.win.width, m.win.height)

		leftSide := lipgloss.JoinVertical(lipgloss.Left, bridgepanel, grouppanel, lightpanel, scenePanel)
		rightSide := lipgloss.JoinVertical(lipgloss.Bottom, detailsPanel, logPanel)

		output := lipgloss.JoinHorizontal(lipgloss.Right, leftSide, rightSide)

		if m.event == bridge.DisplayingColors {
			output = view.Render_color_modal(output, m.color.List.View(), m.win.width, m.win.height)
		} else if m.event == bridge.DisplayingBrightness {
			output = view.Render_bri_modal(output,
				m.brightness.Input.View(),
				m.brightness.Input.Err == nil,
				m.win.width,
				m.win.height)
		}
		return output
	}
	return lipgloss.Place(m.win.width, m.win.height, lipgloss.Center, lipgloss.Center, view.Render_loading_text(m.event))
}
