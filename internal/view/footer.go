package view

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss/v2"
)

func Apply_footer(e bridge.Event, p bridge.Panel, width int) string {

	style := lipgloss.NewStyle().MarginLeft(1).Foreground(lipgloss.Color("#2F74C0"))

	var (
		up, down, left, right                      = "up: ↑/k", "down: ↓/j", "←/h", "→/l"
		enter, esc                                 = "<enter>", "<esc>"
		keybinds                                   = "Keybinds: ?" // not implemented yet
		sep                                        = " | "
		bridgePan, groupsPan, lightsPan, scenesPan = "Bridge: [1]", "Groups: [2]", "Lights: [3]", "Scenes: [4]"
		bri                                        = "Change bright: b"
		decrease_bri, increase_bri                 = left + ": decrease bright", right + ": increase bright"
		off                                        = "on/off: " + enter
		cancel                                     = "cancel: " + esc + "/b"
		apply                                      = "apply: " + enter
		color                                      = "change color: c"
		nextPage, prevPage                         = "PgNext: " + right, "PgPrev: " + left
		filter                                     = "Filter: /"
	)

	var strs []string
	s := ""
	switch p {
	case bridge.BridgePanel:
		strs = append(strs, groupsPan, lightsPan, scenesPan, keybinds)
	case bridge.GroupPanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, up, down, bridgePan, lightsPan, scenesPan, off, bri, keybinds, decrease_bri, increase_bri)
		case bridge.DisplayingBrightness:
			strs = append(strs, cancel, apply)
		}
	case bridge.LightPanel:
		switch e {
		case bridge.DisplayingLights:
			strs = append(strs, up,
				down, bridgePan, groupsPan, scenesPan,
				off, bri, color, keybinds, decrease_bri, increase_bri)
		case bridge.DisplayingBrightness:
			strs = append(strs, cancel, apply)
		case bridge.DisplayingColors:
			//
			strs = append(strs, cancel, prevPage, nextPage, apply, filter)
		}
	case bridge.ScenePanel:
		// there is only one case
	}

	for i := 0; len(s) <= width && i < len(strs); i++ {
		s += strs[i]
		s += sep
	}
	if len(s) == 0 {
		return style.Render(s)
	}
	if len(s)-1 > width {
		return style.Render(s[:width-3] + "..")
	}
	return style.Render(s[:len(s)-2])
}
