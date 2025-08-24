package view

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss/v2"
)

func Apply_footer(e bridge.Event, p bridge.Panel, width int) string {

	style := lipgloss.NewStyle().MarginLeft(1).Foreground(lipgloss.Color("#2F74C0"))

	var strs []string
	s := ""
	switch p {
	case bridge.BridgePanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, render_keybinds(bridge.GroupsPan, bridge.LightsPan, bridge.ScenesPan, bridge.Help)...)
		case bridge.DisplayingHelp:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Up, bridge.Down, bridge.PrevPage, bridge.NextPage)...)
		}
	case bridge.GroupPanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, render_keybinds(bridge.Up, bridge.Down,
				bridge.BridgePan, bridge.LightsPan, bridge.ScenesPan,
				bridge.Off, bridge.Bri, bridge.Help,
				bridge.Decrease_bri, bridge.Increase_bri)...)
		case bridge.DisplayingBrightness:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Apply)...)
		case bridge.DisplayingHelp:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Up, bridge.Down, bridge.PrevPage, bridge.NextPage)...)
		}
	case bridge.LightPanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, render_keybinds(bridge.Up, bridge.Down,
				bridge.BridgePan, bridge.GroupsPan, bridge.ScenesPan,
				bridge.Off, bridge.Bri, bridge.Col, bridge.Help,
				bridge.Decrease_bri, bridge.Increase_bri)...)
		case bridge.DisplayingBrightness:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Apply)...)
		case bridge.DisplayingColors:
			strs = append(strs, render_keybinds(bridge.Cancel,
				bridge.PrevPage, bridge.NextPage,
				bridge.Apply, bridge.Filter)...)
		case bridge.DisplayingHelp:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Up, bridge.Down, bridge.PrevPage, bridge.NextPage)...)
		}
	case bridge.ScenePanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, render_keybinds(bridge.Up, bridge.Down,
				bridge.BridgePan, bridge.GroupsPan, bridge.LightsPan,
				bridge.Off, bridge.Help)...)
		case bridge.DisplayingHelp:
			strs = append(strs, render_keybinds(bridge.Cancel, bridge.Up, bridge.Down, bridge.PrevPage, bridge.NextPage)...)
		}
	}

	for i := 0; len(s) <= width && i < len(strs); i++ {
		s += strs[i]
		s += bridge.Sep
	}
	if len(s) == 0 {
		return style.Render(s)
	}
	if len(s)-1 > width {
		return style.Render(s[:width-3] + "..")
	}
	return style.Render(s[:len(s)-2])
}
func render_keybind(k bridge.Keybind) string {
	return k.Name + ": " + k.Key
}
func render_keybinds(ks ...bridge.Keybind) []string {
	var strs []string
	for _, v := range ks {
		strs = append(strs, render_keybind(v))
	}
	return strs
}
