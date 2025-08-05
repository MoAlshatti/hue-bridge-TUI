package view

import (
	"github.com/MoAlshatti/hue-bridge-TUI/internal/bridge"
	"github.com/charmbracelet/lipgloss"
)

func Render_details_panel(elems string, width, height int) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		PaddingLeft(1).
		MarginTop(1).
		MarginRight(3).
		Height(get_detailspanel_height(height))

	return style.Render(elems)
}

func Render_details(b bridge.Bridge, g bridge.Groups, l bridge.Lights, s bridge.Scenes, p bridge.Panel, width, height int) string {
	var details string
	if p == bridge.BridgePanel {
		details = Render_bridge_details(b, width, height)
	} else if p == bridge.GroupPanel {
		if len(g.Items) > 0 {
			details = Render_group_details(g.Items[g.Cursor], width, height)
		}
	} else if p == bridge.LightPanel {
		if len(l.Items) > 0 {
			details = Render_light_details(*l.Items[l.Cursor], width, height)
		}
	} else if p == bridge.ScenePanel {
		if len(s.Items) > 0 {
			details = Render_scene_details(*s.Items[s.Cursor], width, height)
		}
	}
	return Render_details_panel(details, width, height)
}
